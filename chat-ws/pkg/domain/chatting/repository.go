package chatting

import "context"

type Repository interface {
	GetUserListByChannelId(ctx context.Context, channelId string) ([]UserData, error)
	CreateMessage(ctx context.Context, message *Message) error
}
