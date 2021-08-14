package mysql

import (
	"context"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/editing"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/listing"
)

type Querier interface {
	// user
	CreateUser(ctx context.Context, params account.CreateUserParam) error
	GetUserByUserId(ctx context.Context, userId string) (account.User, error)
	GetUserByUsername(ctx context.Context, username string) (auth.User, error)
	UpdateUser(ctx context.Context, params account.UpdateUserParam) error
	DeleteUser(ctx context.Context, userId string) error

	// workspace
	CreateWorkspace(ctx context.Context, w adding.Workspace) error
	AddUserToWorkspace(ctx context.Context, workspaceId string, userIdList []string) error
	GetWorkspaceByWorkspaceId(ctx context.Context, workspaceId string) (listing.Workspace, error)
	GetWorkspacesByUserId(ctx context.Context, userId string) ([]listing.Workspace, error)
	UpdateWorkspace(ctx context.Context, workspace editing.Workspace) error
	DeleteWorkspace(ctx context.Context, id string) error
	RemoveUserFromWorkspace(ctx context.Context, workspaceId, userId string) error

	// channel
	CreateChannel(ctx context.Context, channel adding.Channel) error
	AddUserToChannel(ctx context.Context, channelId string, userIdList []string) error
	GetChannelsByUserId(ctx context.Context, userId string) ([]listing.Channel, error)
	GetUserListByChannelId(ctx context.Context, channelId string) ([]chatting.User, error)
	UpdateChannel(ctx context.Context, channel editing.Channel) error
	DeleteChannel(ctx context.Context, channelId string) error
	RemoveUserFromChannel(ctx context.Context, channelId string, userId string) error
}
