package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var (
	ChannelGeneral = "General"
	ChannelMessage = "channel-message"
)

type Service interface {
	SetUpWSServer(ctx context.Context) error
	SetUpUserConnection(ctx context.Context, userId string, conn *websocket.Conn)
}

type HandlingService interface {
	HandleBroadcast(ctx context.Context, msg *Message, publishPubSub func(context.Context, *Message) error) error
	GetBroadcastUserListByMessageType(ctx context.Context, msg *Message) ([]UserData, error)
}

type PubSub interface {
	SubscribeToChannel(ctx context.Context, channelName string) <-chan *redis.Message
	PublishToChannel(ctx context.Context, channelName string, payload []byte) error
}

type service struct {
	log             *zap.Logger
	wsServer        *WSServer
	pubsub          PubSub
	chattingService HandlingService
	boardService    HandlingService
}

func NewService(log *zap.Logger, pubsub PubSub, cs HandlingService, bs HandlingService) *service {
	return &service{
		log:             log,
		wsServer:        NewWSServer(),
		pubsub:          pubsub,
		chattingService: cs,
		boardService:    bs,
	}
}

func (s *service) SetUpWSServer(ctx context.Context) error {
	go s.subscribePubSub(ctx)
	go func(ctx context.Context) error {
		for {
			select {
			case user := <-s.wsServer.register:
				s.handleRegister(ctx, user)
			case user := <-s.wsServer.unregister:
				s.handleUnregister(ctx, user.UserId)
			case msg := <-s.wsServer.broadcast:
				if strings.HasPrefix(msg.MessageType, "CHAT") {
					if err := s.chattingService.HandleBroadcast(ctx, msg, s.publishPubSub); err != nil {
						return err
					}
				} else {
					if err := s.boardService.HandleBroadcast(ctx, msg, s.publishPubSub); err != nil {
						return err
					}
				}
			}
		}
	}(ctx)
	return nil
}

func (s *service) SetUpUserConnection(ctx context.Context, userId string, conn *websocket.Conn) {
	user := NewUser(userId, conn, s.wsServer)
	s.wsServer.register <- user
	go user.readMsg(ctx, s.log)
}

func (s *service) subscribePubSub(ctx context.Context) error {
	pubSubChannel := s.pubsub.SubscribeToChannel(ctx, ChannelMessage)

	for data := range pubSubChannel {

		var msg Message

		if err := json.Unmarshal([]byte(data.Payload), &msg); err != nil {
			panic(err)
		}

		if err := s.handleBroadcast(ctx, &msg); err != nil {
			return err
		}
	}
	return nil
}

func (s *service) publishPubSub(ctx context.Context, msg *Message) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		s.log.Warn("Fail to marshal message", zap.Error(err))
		return err
	}
	if err := s.pubsub.PublishToChannel(ctx, ChannelMessage, payload); err != nil {
		s.log.Warn("Fail to publish message to pub sub server", zap.Error(err))
		return err
	}
	return nil
}

func (s *service) handleRegister(ctx context.Context, user *User) {
	s.wsServer.activeUser[user.UserId] = user
}

func (s *service) handleUnregister(ctx context.Context, userId string) {
	delete(s.wsServer.activeUser, userId)
}

func (s *service) handleBroadcast(ctx context.Context, msg *Message) error {
	userList, err := s.getBroadcastUserList(ctx, msg)
	if err != nil {
		return err
	}
	for _, user := range userList {
		if strings.HasPrefix(msg.MessageType, "BOARD_ORDER") &&
			user.UserId == msg.SenderId {
			continue
		}
		u, ok := s.wsServer.activeUser[user.UserId]
		if !ok {
			continue
		}
		err := s.broadcastMessage(ctx, u.UserId, u.conn, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) getBroadcastUserList(ctx context.Context, msg *Message) ([]UserData, error) {
	if strings.HasPrefix(msg.MessageType, "CHAT") {
		return s.chattingService.GetBroadcastUserListByMessageType(ctx, msg)
	} else if strings.HasPrefix(msg.MessageType, "BOARD") {
		return s.boardService.GetBroadcastUserListByMessageType(ctx, msg)
	} else {
		return []UserData{}, errors.New("invalid.message.type")
	}
}

func (s *service) broadcastMessage(ctx context.Context, userId string, conn *websocket.Conn, msg *Message) error {
	err := conn.WriteJSON(msg)
	if err != nil {
		s.log.Info("Unable to write json message.")
		return err
	}
	s.log.Info("Sent message to user id : ", zap.String("user_id", userId))
	return nil
}
