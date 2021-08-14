package chatting

import (
	"context"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type User struct {
	UserId   string `db:"user_id"`
	conn     *websocket.Conn
	wsServer *WSServer
}

func NewUser(userId string, conn *websocket.Conn, wsServer *WSServer) *User {
	return &User{
		UserId:   userId,
		conn:     conn,
		wsServer: wsServer,
	}
}

func (user *User) setUp(ctx context.Context, log *zap.Logger) error {
	go user.readMsg(ctx, log)
	return nil
}

func (user *User) readMsg(ctx context.Context, log *zap.Logger) error {
	defer func() {
		user.conn.Close()
	}()
	for {
		var msg Message
		err := user.conn.ReadJSON(&msg)
		log.Info("Received message from user id :", zap.String("user_id", user.UserId))
		if err != nil {
			log.Info("Unable to read json message.")
			continue
		}
		switch msg.Action {
		case UserLeaveAction:
			user.wsServer.unregister <- user
		case MessageAction:
			user.wsServer.broadcast <- &msg
		}
	}
}
