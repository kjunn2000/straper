package domain

import "github.com/kjunn2000/straper/chat-ws/pkg/storage/redis"

type Channel struct {
	Name       string
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
	Clients    map[*Client]bool
	Rdb        redis.RedisClient
}

func NewChannel(name string) *Channel {
	return &Channel{
		Name:       name,
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
		Clients:    make(map[*Client]bool),
		Rdb:        redis.NewRedisClient(),
	}
}
