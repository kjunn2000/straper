package chatting

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Service interface {
	SaveConnectionToCache(userId string, conn *websocket.Conn) error
	BroadcastMessage(msg Message) error
	StopConnection(userId string) error
}

type Repository interface {
	GetClientListByChannelId(channelId string) ([]Client, error)
}

type RedisClient interface {
	SaveConnnection(userId string, conn *websocket.Conn) error
	GetConnection(userId string) (*websocket.Conn, error)
	StopConnection(userId string) error
}

type service struct {
	Log   *zap.Logger
	Store Repository
	Rdb   RedisClient
}

func NewService(log *zap.Logger, store Repository, rdb RedisClient) *service {
	return &service{
		Log:   log,
		Store: store,
		Rdb:   rdb,
	}
}

func (s *service) ReadMsg(c Client) error {
	defer func() {
		c.Conn.Close()
	}()
	for {
		var msg Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			s.Log.Info("Unable to read json message.")
			continue
		}
		switch msg.Type {
		case 1:
			if err := s.SaveConnectionToCache(c.UserId, c.Conn); err != nil {
				s.Log.Info("Unable save user connection.")
			}
		case 2:
			if err := s.BroadcastMessage(msg); err != nil {
				s.Log.Info("Unable to broadcast message.")
			}
		case 3:
			if err := s.StopConnection(c.UserId); err != nil {
				s.Log.Info("Unable to stop user connection.")
			}
		}
	}
}

func (s *service) SaveConnectionToCache(userId string, conn *websocket.Conn) error {
	return s.Rdb.SaveConnnection(userId, conn)
}

func (s *service) BroadcastMessage(msg Message) error {

	clientList, err := s.Store.GetClientListByChannelId(msg.ChannelId)
	if err != nil {
		s.Log.Info(err.Error())
		return err
	}
	for _, client := range clientList {
		conn, err := s.Rdb.GetConnection(client.UserId)
		if err != nil {
			s.Log.Info(err.Error())
			return err
		}
		err = conn.WriteJSON(msg)
		if err != nil {
			s.Log.Info(err.Error())
			return err
		}
	}
	s.Log.Info("Message sent sucessfully.")
	return nil
}

func (s *service) StopConnection(userId string) error {
	return s.Rdb.StopConnection(userId)
}
