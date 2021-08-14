package editing

import "context"

type Repository interface {
	UpdateWorkspace(ctx context.Context, workspace Workspace) error
	UpdateChannel(ctx context.Context, channel Channel) error
}
