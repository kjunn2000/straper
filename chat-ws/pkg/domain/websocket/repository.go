package websocket

import "context"

type Repository interface {
	GetUserListByChannelId(ctx context.Context, channelId string) ([]UserData, error)
	GetUserListByWorkspaceId(ctx context.Context, workspaceId string) ([]UserData, error)
}
