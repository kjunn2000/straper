package redis

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return RedisClient{
		Client: rdb,
	}
}

func (c *RedisClient) UpdateChatHistory(message string) error {
	ch, err := c.GetChatHistory()
	if err != nil && err != redis.Nil {
		return err
	}
	if ch == nil {
		ch = make([]string, 0)
	}
	ch = append(ch, message)
	fmt.Println(ch)

	var v bytes.Buffer
	if err := gob.NewEncoder(&v).Encode(ch); err != nil {
		return err
	}
	c.Client.Set(context.Background(), "ChatHistory", v.Bytes(), time.Minute*1)
	return nil
}

func (c *RedisClient) GetChatHistory() ([]string, error) {
	get := c.Client.Get(context.Background(), "ChatHistory")
	historyBytes, err := get.Bytes()
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(historyBytes)
	chatHistory := make([]string, 0)
	if err = gob.NewDecoder(r).Decode(&chatHistory); err != nil {
		return nil, err
	}
	return chatHistory, nil
}
