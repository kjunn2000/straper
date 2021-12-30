package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
	SubscribeToChannel(ctx context.Context, channelName string) <-chan *redis.Message
	PublishToChannel(ctx context.Context, channelName string, payload []byte) error
}

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(rdb *redis.Client) redisClient {
	return redisClient{
		client: rdb,
	}
}

func (c *redisClient) SubscribeToChannel(ctx context.Context, channelName string) <-chan *redis.Message {
	return c.client.Subscribe(ctx, channelName).Channel()
}

func (c *redisClient) PublishToChannel(ctx context.Context, channelName string, payload []byte) error {
	return c.client.Publish(ctx, channelName, payload).Err()
}
