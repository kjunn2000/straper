	package domain

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type ChannelRepository interface {
	CreateChannel(channel *Channel) error
}

type ChannelService interface {
	CreateChannel(workspaceId, channelName string) error
	RegisterClient(conn *websocket.Conn, channel *Channel) error
}

type channelService struct {
	log *zap.Logger
	cr  ChannelRepository
}

func NewChannelService(log *zap.Logger, cr ChannelRepository) *channelService {
	return &channelService{
		log: log,
		cr:  cr,
	}
}

func (ch *channelService) CreateChannel(workspaceId, channelName string) error {
	c := ch.openChannelServer(channelName)
	err := ch.cr.CreateChannel(c)
	if err != nil {
		return err
	}
	return nil
}

func (ch *channelService) openChannelServer(name string) *Channel {
	channel := NewChannel(name)
	go channel.StartWSServer()
	return channel
}

func (ch *channelService) RegisterClient(conn *websocket.Conn, channel *Channel) {
	c := NewClient(conn, channel)
	channel.Register <- c
	c.ReadMsg()
}
