package chatting

import (
	"context"

	ws "github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
)

type Repository interface {
	GetUserListByChannelId(ctx context.Context, channelId string) ([]ws.UserData, error)
	GetChatUserInfoByUserId(ctx context.Context, userId string) (UserDetail, error)
	CreateMessage(ctx context.Context, message *Message) error
	GetChannelMessages(ctx context.Context, channelId string, param PaginationMessagesParam) ([]Message, error)
	GetAllChannelMessages(ctx context.Context, channelId string) ([]Message, error)
	GetAllChannelMessagesByWorkspaceId(ctx context.Context, workspaceId string) ([]Message, error)
	EditMessage(ctx context.Context, params EditChatMessageParams) error
	DeleteMessage(ctx context.Context, messageId string) error
}
