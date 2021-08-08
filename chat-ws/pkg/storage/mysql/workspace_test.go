package mysql

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/storage"
	"github.com/stretchr/testify/require"
)

func TestCreateWorkspace(t *testing.T) {
	randUser := generateRandomoUser()
	err := store.SaveUser(randUser)
	require.NoError(t, err)
	user, err := store.FindUserByUsername(randUser.Username)
	require.NoError(t, err)
	workspaceId := uuid.New().String()
	workspace := adding.Workspace{
		Id:        workspaceId,
		Name:      storage.RandomString(6),
		CreatorId: user.UserId,
	}
	channel := adding.Channel{
		ChannelId:   uuid.New().String(),
		ChannelName: "General",
		WorkspaceId: workspaceId,
	}
	_, err = store.CreateWorkspace(workspace, channel, user.UserId)
	require.NoError(t, err)
}

func TestAddUserToWorksapce(t *testing.T) {
	randUser := generateRandomoUser()
	err := store.SaveUser(randUser)
	require.NoError(t, err)
	user, err := store.FindUserByUsername(randUser.Username)
	require.NoError(t, err)
	randUser2 := generateRandomoUser()
	err = store.SaveUser(randUser2)
	require.NoError(t, err)
	user2, err := store.FindUserByUsername(randUser2.Username)
	require.NoError(t, err)
	workspaceId := uuid.New().String()
	workspace := adding.Workspace{
		Id:        workspaceId,
		Name:      storage.RandomString(6),
		CreatorId: user.UserId,
	}
	channel := adding.Channel{
		ChannelId:   uuid.New().String(),
		ChannelName: "General",
		WorkspaceId: workspaceId,
	}
	_, err = store.CreateWorkspace(workspace, channel, user.UserId)
	require.NoError(t, err)
	err = store.AddUserToWorkspace(workspaceId, []string{user2.UserId})
	require.NoError(t, err)
}
