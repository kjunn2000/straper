package chatting

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	UserId string `json:"user_id" db:"user_id"`
	Conn   *websocket.Conn
}

func NewClient(userId string, conn *websocket.Conn) *Client {
	return &Client{
		UserId: userId,
		Conn:   conn,
	}
}

type Message struct {
	ChannelId string `json:"channel_id"`
	Type      int    `json:"type"`
	Content   string `json:"content"`
}
