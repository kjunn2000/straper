package chatting

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var (
	ChannelGeneral = "General"
)

type Service interface {
	SetUpWSServer(ctx context.Context) error
	SetUpUserConnection(ctx context.Context, userId string, conn *websocket.Conn) error
}

type PubSub interface {
	SubscribeChannel(ctx context.Context, channelName string) <-chan *redis.Message
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
	go func(ctx context.Context) error {
		for {
			select {
			case user := <-s.wsServer.register:
				s.handleRegister(ctx, user)
			case user := <-s.wsServer.unregister:
				s.handleUnregister(ctx, user)
			case msg := <-s.wsServer.broadcast:
				err := s.handleBroadcast(ctx, msg)
				if err != nil {
					return err
				}
			}
		}
	}(ctx)
	return nil
}

func (s *service) SetUpUserConnection(ctx context.Context, userId string, conn *websocket.Conn) error {
	user := NewUser(userId, conn, s.wsServer)
	s.wsServer.register <- user
	go user.setUp(ctx, s.log)
	return nil
}

func (s *service) handleRegister(ctx context.Context, user *User) {
	s.wsServer.activeUser[user.UserId] = user
}

func (s *service) handleUnregister(ctx context.Context, user *User) {
	delete(s.wsServer.activeUser, user.UserId)
}

func (s *service) handleBroadcast(ctx context.Context, msg *Message) error {
	userList, err := s.store.GetUserListByChannelId(ctx, msg.ChannelId)
	if err != nil {
		return err
	}
	for _, user := range userList {
		u, ok := s.wsServer.activeUser[user.UserId]
		if !ok {
			return errors.New("failed.get.user.list")
		}
		value, err := msg.Encode()
		if err != nil {
			return err
		}
		err = s.broadcastMessage(ctx, u.UserId, u.conn, value)
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
