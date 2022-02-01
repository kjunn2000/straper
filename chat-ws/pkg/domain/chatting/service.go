package chatting

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	ws "github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
	"go.uber.org/zap"
)

var (
	ChannelGeneral = "General"
	ChannelMessage = "channel-message"
)

type Service interface {
	GetChannelMessages(ctx context.Context, channelId string, userId string, limit, offset uint64) ([]Message, error)
	HandleBroadcast(ctx context.Context, msg *ws.Message, publishPubSub func(context.Context, *ws.Message) error) error
	GetBroadcastUserListByMessageType(ctx context.Context, msg *ws.Message) ([]ws.UserData, error)
	DeleteSeaweedfsMessagesByChannelId(ctx context.Context, channelId string) error
	DeleteSeaweedfsMessagesByWorkspaceId(ctx context.Context, workspaceId string) error
}

type PubSub interface {
	SubscribeToChannel(ctx context.Context, channelName string) <-chan *redis.Message
	PublishToChannel(ctx context.Context, channelName string, payload []byte) error
}

type SeaweedfsClient interface {
	SaveSeaweedfsFile(ctx context.Context, fileBytes []byte) (string, error)
	GetSeaweedfsFile(ctx context.Context, fid string) ([]byte, error)
	DeleteSeaweedfsFile(ctx context.Context, fid string) error
}

type service struct {
	log   *zap.Logger
	store Repository
	sc    SeaweedfsClient
}

func NewService(log *zap.Logger, store Repository, sc SeaweedfsClient) *service {
	return &service{
		log:   log,
		store: store,
		sc:    sc,
	}
}

func (s *service) HandleBroadcast(ctx context.Context, msg *ws.Message, publishPubSub func(context.Context, *ws.Message) error) error {
	bytePayload, err := msg.Payload.MarshalJSON()
	if err != nil {
		return err
	}
	switch msg.MessageType {
	case ChatAddMessage:
		if newPayload, err := s.handleAddChatMessage(ctx, bytePayload); err != nil {
			return err
		} else {
			if err := msg.Payload.UnmarshalJSON(newPayload); err != nil {
				return err
			}
		}
	}
	if err := publishPubSub(ctx, msg); err != nil {
		return err
	}
	return nil
}

func (s *service) handleAddChatMessage(ctx context.Context, bytePayload []byte) ([]byte, error) {
	var message Message
	if err := json.Unmarshal(bytePayload, &message); err != nil {
		return []byte{}, err
	}
	newId, err := uuid.NewRandom()
	if err != nil {
		return []byte{}, err
	}
	message.MessageId = newId.String()
	message.CreatedDate = time.Now()
	if message.Type == TypeFile {
		fid, err := s.sc.SaveSeaweedfsFile(ctx, message.FileBytes)
		if err != nil {
			return []byte{}, err
		}
		message.Content = fid
	}
	if err := s.store.CreateMessage(ctx, &message); err != nil {
		return []byte{}, err
	}
	userDetail, err := s.store.GetChatUserInfoByUserId(ctx, message.CreatorId)
	if err != nil {
		s.log.Warn("Fail to fetch user data.", zap.Error(err))
		return []byte{}, err
	}
	message.UserDetail = userDetail
	newMsg, err := json.Marshal(message)
	if err != nil {
		return []byte{}, err
	}
	return newMsg, nil
}

func (s *service) GetBroadcastUserListByMessageType(ctx context.Context, msg *ws.Message) ([]ws.UserData, error) {
	var message Message
	bytePayload, err := msg.Payload.MarshalJSON()
	if err := json.Unmarshal(bytePayload, &message); err != nil {
		return []ws.UserData{}, err
	}
	if err != nil {
		return []ws.UserData{}, err
	}
	return s.store.GetUserListByChannelId(ctx, message.ChannelId)
}

func (s *service) GetChannelMessages(ctx context.Context, channelId string, userId string, limit, offset uint64) ([]Message, error) {
	msgs, err := s.store.GetChannelMessages(ctx, channelId, limit, offset)
	if err == sql.ErrNoRows {
		return []Message{}, errors.New("invalid.channel.id")
	} else if err != nil {
		return []Message{}, err
	}
	for i, msg := range msgs {
		if msg.Type == "FILE" {
			bytesData, err := s.sc.GetSeaweedfsFile(ctx, msg.Content)
			if err != nil {
				return []Message{}, err
			}
			msg.FileBytes = bytesData
			msgs[i] = msg
		}
		userDetail, err := s.store.GetChatUserInfoByUserId(ctx, msg.CreatorId)
		if err != nil {
			return []Message{}, err
		} else {
			msgs[i].UserDetail = userDetail
		}
	}
	if err = s.store.UpdateChannelAccessTime(ctx, channelId, userId); err != nil {
		return []Message{}, err
	}
	return msgs, nil
}

func (s *service) DeleteSeaweedfsMessagesByChannelId(ctx context.Context, channelId string) error {
	msgs, err := s.store.GetAllChannelMessages(ctx, channelId)
	if err != nil {
		return err
	}
	return s.deleteSeaweedfsMessages(ctx, msgs)
}

func (s *service) DeleteSeaweedfsMessagesByWorkspaceId(ctx context.Context, workspaceId string) error {
	msgs, err := s.store.GetAllChannelMessagesByWorkspaceId(ctx, workspaceId)
	if err != nil {
		return err
	}
	return s.deleteSeaweedfsMessages(ctx, msgs)
}

func (s *service) deleteSeaweedfsMessages(ctx context.Context, msgs []Message) error {
	for _, msg := range msgs {
		if msg.Type == TypeFile {
			if err := s.sc.DeleteSeaweedfsFile(ctx, msg.Content); err != nil {
				return err
			}
		}
	}
	return nil
}
