package chatting

import "context"

type Repository interface {
	GetUserListByChannelId(ctx context.Context, channelId string) ([]UserData, error)
	CreateMessage(ctx context.Context, message *Message) error
	GetChannelMessages(ctx context.Context, channelId string, limit, offset uint64) ([]Message, error)
	GetAllChannelMessages(ctx context.Context, channelId string) ([]Message, error)
	GetAllChannelMessagesByWorkspaceId(ctx context.Context, workspaceId string) ([]Message, error)
	UpdateChannelAccessTime(ctx context.Context, channelId string, userId string) error
}
