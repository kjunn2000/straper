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

type UserData struct {
	UserId string `db:"user_id"`
}

func (user *User) readMsg(ctx context.Context, log *zap.Logger) {
	defer func() {
		user.conn.Close()
	}()
	for {
		var msg Message
		err := user.conn.ReadJSON(&msg)
		log.Info("Received message from user id :", zap.String("user_id", user.UserId))
		if websocket.IsCloseError(err, 1000, 1001, 1005) {
			log.Info("Websocket Conn Closed.", zap.String("user_id", user.UserId))
			user.wsServer.unregister <- user
			return
		} else if err != nil {
			log.Warn("Receive error.", zap.Error(err))
			return
		}
		switch msg.Type {
		case UserLeave:
			user.wsServer.unregister <- user
			return
		case Messaging:
			user.wsServer.broadcast <- &msg
		}
	}
}
