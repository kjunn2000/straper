package redis

import (
	"bytes"
	"context"
	"encoding/gob"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type redisClient struct {
	Client *redis.Client
	Log    *zap.Logger
}

func NewRedisClient(rdb *redis.Client, log *zap.Logger) redisClient {
	return redisClient{
		Client: rdb,
		Log:    log,
	}
}

func (c *redisClient) SaveConnnection(userId string, conn *websocket.Conn) error {
	var v bytes.Buffer
	if err := gob.NewEncoder(&v).Encode(conn); err != nil {
		c.Log.Info("Unable to encode user connection.", zap.Error(err))
		return err
	}
	err := c.Client.Set(context.Background(), userId, v.Bytes(), 0).Err()
	if err != nil {
		c.Log.Info("Failed to set value to redis.", zap.Error(err))
		return err
	}
	return nil
}

func (c *redisClient) GetConnection(userId string) (*websocket.Conn, error) {
	connBytes, err := c.Client.Get(context.Background(), userId).Bytes()
	if err != nil {
		c.Log.Info("Failed to get user connection from redis.", zap.Error(err))
		return nil, err
	}
	r := bytes.NewReader(connBytes)
	var conn *websocket.Conn
	if err = gob.NewDecoder(r).Decode(conn); err != nil {
		c.Log.Info("Failed to decode connection.", zap.Error(err))
		return nil, err
	}
	return conn, nil
}

func (c *redisClient) StopConnection(userId string) error {
	err := c.Client.Del(context.Background(), userId).Err()
	if err != nil {
		c.Log.Info("Failed to delete connection at redis.", zap.Error(err))
		return err
	}
	return nil
}