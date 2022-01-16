package chatting

import (
	"context"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
)

type Repository interface {
	GetUserInfoByUserId(ctx context.Context, userId string) (websocket.UserDetail, error)
	CreateMessage(ctx context.Context, message *Message) error
	GetChannelMessages(ctx context.Context, channelId string, limit, offset uint64) ([]Message, error)
	GetAllChannelMessages(ctx context.Context, channelId string) ([]Message, error)
	GetAllChannelMessagesByWorkspaceId(ctx context.Context, workspaceId string) ([]Message, error)
	UpdateChannelAccessTime(ctx context.Context, channelId string, userId string) error
}
