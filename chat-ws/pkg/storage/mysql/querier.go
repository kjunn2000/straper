package mysql

import (
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/listing"
)

type Querier interface {
	SaveUser(u *account.User) error
	CheckUsernameExist(username string) (bool, error)
	FindUserByUsername(username string) (*auth.User, error)
	CreateWorkspace(w adding.Workspace) error
	DeleteWorkspace(id string) error
	RemoveUserFromWorkspace(workspaceId, userId string) error
	GetWorkspaceByWorkspaceId(workspaceId string) (listing.Workspace, error)
	GetWorkspacesByUserId(userId string) ([]listing.Workspace, error)
	AddUserToWorkspace(workspaceId string, userIdList []string) error
	CreateChannel(channel *adding.Channel) error
	AddUserToChannel(channelId string, userIdList []string) error
	GetAllChannelByUserAndWorkspaceId(userId, workspaceId string) ([]listing.Channel, error)
	GetClientListByChannelId(channelId string) ([]chatting.Client, error)
	GetAllChannelByWorkspaceId(workspaceId string) ([]listing.Channel, error)
	GetDefaultChannelByWorkspaceId(workspaceId string) (adding.Channel, error)
	DeleteChannel(channelId string) error
	RemoveUserFromChannelList(channelIdList []string, userId string) error
}
