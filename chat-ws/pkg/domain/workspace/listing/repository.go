package listing

import "context"

type Repository interface {
	GetWorkspacesByUserId(ctx context.Context, userId string) ([]Workspace, error)
	GetWorkspaceByWorkspaceId(ctx context.Context, workspaceId string) (Workspace, error)
	GetChannelsByUserId(ctx context.Context, userId string) ([]Channel, error)
	GetDefaultChannel(ctx context.Context, workspaceId string) (Channel, error)
}
