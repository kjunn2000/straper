package deleting

import "context"

type Repository interface {
	DeleteWorkspace(ctx context.Context, workspaceId string) error
	RemoveUserFromWorkspace(ctx context.Context, workspaceId, userId string) error
	DeleteChannel(ctx context.Context, channelId string) error
	RemoveUserFromChannel(ctx context.Context, channelId string, userId string) error
}
