package domain

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn    *websocket.Conn
	Channel *Channel
}

func NewClient(conn *websocket.Conn, channel *Channel) *Client {
	return &Client{
		Conn:    conn,
		Channel: channel,
	}
}

type Message struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type UpdateHistoryMessage struct {
	Type    int      `json:"type"`
	Content []string `json:"content"`
}

func (c *Client) ReadMsg() {
	defer func() {
		c.Channel.Unregister <- c
		c.Conn.Close()
	}()
	for {
		var msg Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Unable to read message from connection")
			return
		}
		c.Channel.Broadcast <- msg
	}
}
