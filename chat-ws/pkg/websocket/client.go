package websocket

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Pool *Pool
}

func NewClient(conn *websocket.Conn, pool *Pool) *Client {
	return &Client{
		Conn: conn,
		Pool: pool,
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
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()
	for {
		var msg Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Unable to read message from connection")
			return
		}
		c.Pool.Broadcast <- msg
	}
}
