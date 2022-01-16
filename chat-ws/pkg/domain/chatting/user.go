package chatting

import (
	"github.com/gorilla/websocket"
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
