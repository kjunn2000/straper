package chatting

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
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
	GetChannelMessages(ctx context.Context, channelId string, param PaginationMessagesParam) ([]Message, error)
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
	SaveFile(ctx context.Context, reader io.Reader) (string, error)
	GetFile(ctx context.Context, fid string) ([]byte, error)
	DeleteFile(ctx context.Context, fid string) error
}

type PaginationService interface {
	DecodeCursor(encodedCursor string) (res time.Time, uuid string, err error)
	EncodeCursor(t time.Time, uuid string) string
}

type service struct {
	log   *zap.Logger
	store Repository
	sc    SeaweedfsClient
	ps    PaginationService
}

func NewService(log *zap.Logger, store Repository, sc SeaweedfsClient, ps PaginationService) *service {
	return &service{
		log:   log,
		store: store,
		sc:    sc,
		ps:    ps,
	}
}

func (s *service) HandleBroadcast(ctx context.Context, msg *ws.Message, publishPubSub func(context.Context, *ws.Message) error) error {
	bytePayload, err := msg.Payload.MarshalJSON()
	if err != nil {
		return err
	}
	switch msg.MessageType {
	case ChatAddMessage:
		if newPayload, err := s.handleAddMessage(ctx, bytePayload); err != nil {
			return err
		} else {
			if err := msg.Payload.UnmarshalJSON(newPayload); err != nil {
				return err
			}
		}
	case ChatEditMessage:
		if err := s.handleEditMessage(ctx, bytePayload); err != nil {
			return err
		}
	case ChatDeleteMessage:
		if err := s.handleDeleteMessage(ctx, bytePayload); err != nil {
			return err
		}
	}
	if err := publishPubSub(ctx, msg); err != nil {
		return err
	}
	return nil
}

func (s *service) handleAddMessage(ctx context.Context, bytePayload []byte) ([]byte, error) {
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
		fid, err := s.sc.SaveFile(ctx, bytes.NewReader(message.FileBytes))
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

func (service *service) handleEditMessage(ctx context.Context, bytePayload []byte) error {
	var editChatMessageParams EditChatMessageParams
	if err := json.Unmarshal(bytePayload, &editChatMessageParams); err != nil {
		return err
	}
	return service.store.EditMessage(ctx, editChatMessageParams)
}

func (service *service) handleDeleteMessage(ctx context.Context, bytePayload []byte) error {
	var deleteChatMessageParams DeleteChatMessageParams
	if err := json.Unmarshal(bytePayload, &deleteChatMessageParams); err != nil {
		return err
	}
	if deleteChatMessageParams.Type == TypeFile {
		if err := service.sc.DeleteFile(ctx, deleteChatMessageParams.Fid); err != nil {
			return err
		}
	}
	return service.store.DeleteMessage(ctx, deleteChatMessageParams.MessageId)
}

func (s *service) GetBroadcastUserListByMessageType(ctx context.Context, msg *ws.Message) ([]ws.UserData, error) {
	return s.store.GetUserListByChannelId(ctx, msg.ChannelId)
}

func (s *service) GetChannelMessages(ctx context.Context, channelId string, param PaginationMessagesParam) ([]Message, error) {
	if param.Cursor != "" {
		time, uuid, err := s.ps.DecodeCursor(param.Cursor)
		if err != nil {
			return []Message{}, err
		}
		param.CreatedTime = time
		param.Id = uuid
	}
	msgs, err := s.store.GetChannelMessages(ctx, channelId, param)
	if err == sql.ErrNoRows {
		return []Message{}, errors.New("invalid.channel.id")
	} else if err != nil {
		return []Message{}, err
	}
	for i, msg := range msgs {
		if msg.Type == "FILE" {
			bytesData, err := s.sc.GetFile(ctx, msg.Content)
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
		msgs[i].Cursor = s.ps.EncodeCursor(msg.CreatedDate, msg.MessageId)
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
			if err := s.sc.DeleteFile(ctx, msg.Content); err != nil {
				return err
			}
		}
	}
	return nil
}
