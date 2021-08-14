package listing

import "context"

type Repository interface {
	GetWorkspacesByUserId(ctx context.Context, userId string) ([]Workspace, error)
	GetChannelsByUserId(ctx context.Context, userId string) ([]Channel, error)
}
