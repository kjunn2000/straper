package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type RedisClient interface {
	SubscribeChannel(ctx context.Context, channelName string) <-chan *redis.Message
}

type redisClient struct {
	client *redis.Client
	log    *zap.Logger
}

func NewRedisClient(rdb *redis.Client, log *zap.Logger) redisClient {
	return redisClient{
		client: rdb,
		log:    log,
	}
}

func (c *redisClient) SubscribeChannel(ctx context.Context, channelName string) <-chan *redis.Message {
	return c.client.Subscribe(ctx, channelName).Channel()
}
