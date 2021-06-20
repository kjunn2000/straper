package websocket

import (
	rdb "github.com/kjunn2000/straper/chat-ws/pkg/redis"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
	Clients    map[*Client]bool
	Rdb        rdb.RedisClient
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
		Clients:    make(map[*Client]bool),
		Rdb:        rdb.NewRedisClient(),
	}
}
