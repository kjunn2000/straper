package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/editing"
	"github.com/kjunn2000/straper/chat-ws/pkg/storage"
	"github.com/stretchr/testify/require"
)

func createRandomWorkspace(t *testing.T) adding.Workspace {
	user := createRandomUser(t)
	newUser, err := store.GetUserDetailByUsername(context.Background(), user.Username)
	require.NoError(t, err)
	id, err := uuid.NewRandom()
	require.NoError(t, err)
	workspace := adding.Workspace{
		Id:          id.String(),
		Name:        storage.RandomString(6),
		CreatorId:   newUser.UserId,
		CreatedDate: time.Now(),
	}
	err = store.CreateWorkspace(context.Background(), workspace)
	require.NoError(t, err)
	return workspace
}

func TestCreateWorkspace(t *testing.T) {
	createRandomWorkspace(t)
}

func TestAddUserToWorksapce(t *testing.T) {
	randUser := createRandomUser(t)
	user, err := store.GetUserDetailByUsername(context.Background(), randUser.Username)
	require.NoError(t, err)
	workspace := createRandomWorkspace(t)
	err = store.AddUserToWorkspace(context.Background(), workspace.Id, []string{user.UserId})
	require.NoError(t, err)
}

func TestGetWorkspaceByWorkspaceId(t *testing.T) {
	newWorkspace := createRandomWorkspace(t)
	workspace, err := store.GetWorkspaceByWorkspaceId(context.Background(), newWorkspace.Id)
	require.NoError(t, err)

	require.Equal(t, workspace.Id, newWorkspace.Id)
	require.Equal(t, workspace.Name, newWorkspace.Name)
	require.Equal(t, workspace.CreatorId, newWorkspace.CreatorId)
}

func TestGetWorkspaceByUserId(t *testing.T) {
	randUser := createRandomUser(t)
	user, err := store.GetUserDetailByUsername(context.Background(), randUser.Username)
	require.NoError(t, err)
	randomWorkspaces := make([]adding.Workspace, 5)
	for i := 0; i < 5; i++ {
		workspace := createRandomWorkspace(t)
		randomWorkspaces[i] = workspace
		err = store.AddUserToWorkspace(context.Background(), randomWorkspaces[i].Id, []string{user.UserId})
		require.NoError(t, err)
	}

	workspaces, err := store.GetWorkspacesByUserId(context.Background(), user.UserId)
	require.NoError(t, err)

	require.Equal(t, len(workspaces), 5)
}

func TestUpdateWorkspace(t *testing.T) {
	workspace := createRandomWorkspace(t)
	toUpdateWorkspace := editing.Workspace{
		Id:   workspace.Id,
		Name: storage.RandomString(6),
	}
	err := store.UpdateWorkspace(context.Background(), toUpdateWorkspace)
	require.NoError(t, err)
}

func TestDeleteWorkspace(t *testing.T) {
	workspace := createRandomWorkspace(t)
	err := store.DeleteWorkspace(context.Background(), workspace.Id)
	require.NoError(t, err)
}

func TestRemoveUserFormWorkspace(t *testing.T) {
	user := createRandomUser(t)
	newUser, err := store.GetUserDetailByUsername(context.Background(), user.Username)
	require.NoError(t, err)
	workspace := createRandomWorkspace(t)
	err = store.AddUserToWorkspace(context.Background(), workspace.Id, []string{newUser.UserId})
	require.NoError(t, err)
	err = store.RemoveUserFromWorkspace(context.Background(), workspace.Id, newUser.UserId)
	require.NoError(t, err)
}

// Test Workspace Store

func createNewRandomWorkspace(t *testing.T) adding.Workspace {
	u := createRandomUser(t)
	user, err := store.GetUserDetailByUsername(context.Background(), u.Username)
	require.NoError(t, err)
	randomWorkspace := adding.Workspace{
		Id:          uuid.New().String(),
		Name:        storage.RandomString(6),
		CreatorId:   user.UserId,
		CreatedDate: time.Now(),
	}
	randomChannel := adding.Channel{
		ChannelId:   uuid.New().String(),
		ChannelName: "General",
		WorkspaceId: randomWorkspace.Id,
		CreatorId:   user.UserId,
		CreatedDate: time.Now(),
	}
	w, err := store.CreateNewWorkspace(context.Background(), randomWorkspace, randomChannel, user.UserId)
	require.NoError(t, err)
	require.Equal(t, randomWorkspace.Id, w.Id)
	require.Equal(t, randomWorkspace.Name, w.Name)
	require.Equal(t, randomWorkspace.CreatorId, w.CreatorId)
	require.Equal(t, randomWorkspace.CreatedDate, w.CreatedDate)

	return w
}

func TestCreateNewWorkspace(t *testing.T) {
	createNewRandomWorkspace(t)
}

func TestAddNewUserToWorkspace(t *testing.T) {
	w := createNewRandomWorkspace(t)
	u := createRandomUser(t)
	user, err := store.GetUserDetailByUsername(context.Background(), u.Username)
	require.NoError(t, err)
	err = store.AddNewUserToWorkspace(context.Background(), w.Id, []string{user.UserId})
	require.NoError(t, err)
}
