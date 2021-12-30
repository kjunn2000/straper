package chatting

import (
	"context"
	"encoding/json"

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

type PubSub interface {
	SubscribeToChannel(ctx context.Context, channelName string) <-chan *redis.Message
	PublishToChannel(ctx context.Context, channelName string, payload []byte) error
}

type service struct {
	log      *zap.Logger
	store    Repository
	pubsub   PubSub
	wsServer *WSServer
}

func NewService(log *zap.Logger, store Repository, pubsub PubSub) *service {
	return &service{
		log:      log,
		store:    store,
		pubsub:   pubsub,
		wsServer: NewWSServer(),
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
				err := s.publishPubSub(ctx, msg)
				if err != nil {
					return err
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

func (s *service) handleRegister(ctx context.Context, user *User) {
	s.wsServer.activeUser[user.UserId] = user
}

func (s *service) handleUnregister(ctx context.Context, userId string) {
	delete(s.wsServer.activeUser, userId)
}

func (s *service) handleBroadcast(ctx context.Context, msg *Message) error {
	userList, err := s.store.GetUserListByChannelId(ctx, msg.ChannelId)
	if err != nil {
		return err
	}
	for _, user := range userList {
		u, ok := s.wsServer.activeUser[user.UserId]
		if !ok {
			continue
		}
		encodeMsg, err := msg.Encode()
		if err != nil {
			return err
		}
		err = s.broadcastMessage(ctx, u.UserId, u.conn, encodeMsg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) broadcastMessage(ctx context.Context, userId string, conn *websocket.Conn, msg []byte) error {
	err := conn.WriteJSON(msg)
	if err != nil {
		s.log.Info("Unable to write json message.")
		return err
	}
	s.log.Info("Sent message to user id : ", zap.String("user_id", userId))
	return nil
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
