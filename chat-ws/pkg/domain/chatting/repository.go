package chatting

import (
	"context"

	ws "github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
)

type Repository interface {
	GetUserListByChannelId(ctx context.Context, channelId string) ([]ws.UserData, error)
	GetUserInfoByUserId(ctx context.Context, userId string) (UserDetail, error)
	CreateMessage(ctx context.Context, message *Message) error
	GetChannelMessages(ctx context.Context, channelId string, limit, offset uint64) ([]Message, error)
	GetAllChannelMessages(ctx context.Context, channelId string) ([]Message, error)
	GetAllChannelMessagesByWorkspaceId(ctx context.Context, workspaceId string) ([]Message, error)
	UpdateChannelAccessTime(ctx context.Context, channelId string, userId string) error

	CreateCardComment(ctx context.Context, comment *CardComment) error
	GetCardComments(ctx context.Context, cardId string) ([]CardComment, error)
}
