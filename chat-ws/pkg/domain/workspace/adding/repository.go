package adding

import "context"

type Repository interface {
	CreateNewWorkspace(ctx context.Context, w Workspace, c Channel, userId string) (Workspace, error)
	AddNewUserToWorkspace(ctx context.Context, workspaceId string, userIdList []string) error
	CreateNewChannel(ctx context.Context, channel Channel, userId string) (Channel, error)
	AddUserToChannel(ctx context.Context, channelId string, userId []string) error
}
