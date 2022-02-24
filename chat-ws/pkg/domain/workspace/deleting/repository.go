package deleting

import (
	"context"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/listing"
)

type Repository interface {
	GetWorkspaceByWorkspaceId(ctx context.Context, workspaceId string) (listing.Workspace, error)
	GetChannelByChannelId(ctx context.Context, channelId string) (listing.Channel, error)
	DeleteWorkspace(ctx context.Context, workspaceId string) error
	RemoveUserFromWorkspace(ctx context.Context, workspaceId, userId string) error
	DeleteChannel(ctx context.Context, channelId string) error
	RemoveUserFromChannel(ctx context.Context, channelId string, userId string) error

	GetFidsByWorkspaceId(ctx context.Context, workspaceId string) ([]string, error)
	GetAttachmentFidsByWorkspaceId(ctx context.Context, workspaceId string) ([]string, error)
}
