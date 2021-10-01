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
	CreateUserDetail(ctx context.Context, params CreateUserDetailParam) error
	CreateUserCredential(ctx context.Context, params CreateUserCredentialParam) error
	CreateUserAccessInfo(ctx context.Context, params CreateUserAccessInfo) error
	GetUserDetailByUsername(ctx context.Context, username string) (account.UserDetail, error)
	GetUserDetailByUserId(ctx context.Context, userId string) (account.UserDetail, error)
	GetUserCredentialByUsername(ctx context.Context, username string) (auth.User, error)
	GetUserCredentialByUserId(ctx context.Context, userId string) (auth.User, error)
	UpdateUser(ctx context.Context, params account.UpdateUserParam) error
	UpdateAccountStatus(ctx context.Context, userId, status string) error
	UpdateAccountPassword(ctx context.Context, userId, password string) error
	DeleteUser(ctx context.Context, userId string) error

	// verify_email
	CreateVerifyEmailToken(ctx context.Context, token account.VerifyEmailToken) error
	GetVerifyEmailToken(ctx context.Context, userId string) (account.VerifyEmailToken, error)
	DeleteVerifyEmailToken(ctx context.Context, tokenId string) error

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
