package mysql

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
// 	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/editing"
// 	"github.com/kjunn2000/straper/chat-ws/pkg/storage"
// 	"github.com/stretchr/testify/require"
// )

// func TestCreateWorkspace(t *testing.T) {
// 	user := createRandomUser(t)
// 	newUser, err := store.GetUserByUsername(context.Background(), user.Username)
// 	require.NoError(t, err)
// 	id, err := uuid.NewRandom()
// 	require.NoError(t, err)
// 	workspace := adding.Workspace{
// 		Id:          id.String(),
// 		Name:        storage.RandomString(6),
// 		CreatorId:   newUser.UserId,
// 		CreatedDate: time.Now(),
// 	}
// 	err = store.CreateWorkspace(context.Background(), workspace)
// 	require.NoError(t, err)
// }

// func TestAddUserToWorksapce(t *testing.T) {
// 	randUser := createRandomUser(t)
// 	user, err := store.GetUserByUsername(context.Background(), randUser.Username)
// 	require.NoError(t, err)
// 	workspace := createRandomWorkspace(t)
// 	err = store.AddUserToWorkspace(context.Background(), workspace.Id, []string{user.UserId})
// 	require.NoError(t, err)
// }

// func TestGetWorkspaceByWorkspaceId(t *testing.T) {
// 	newWorkspace := createRandomWorkspace(t)
// 	workspace, err := store.GetWorkspaceByWorkspaceId(context.Background(), newWorkspace.Id)
// 	require.NoError(t, err)

// 	require.Equal(t, workspace.Id, newWorkspace.Id)
// 	require.Equal(t, workspace.Name, newWorkspace.Name)
// 	require.Equal(t, workspace.CreatorId, newWorkspace.CreatorId)
// }

// func TestGetWorkspaceByUserId(t *testing.T) {
// 	randUser := createRandomUser(t)
// 	user, err := store.GetUserByUsername(context.Background(), randUser.Username)
// 	require.NoError(t, err)
// 	randomWorkspaces := make([]adding.Workspace, 5)
// 	for i := 0; i < 5; i++ {
// 		randomWorkspaces[i] = createRandomWorkspace(t)
// 		err = store.AddUserToWorkspace(context.Background(), randomWorkspaces[i].Id, []string{user.UserId})
// 		require.NoError(t, err)
// 		time.Sleep(time.Second)
// 	}

// 	workspaces, err := store.GetWorkspacesByUserId(context.Background(), user.UserId)
// 	require.NoError(t, err)

// 	for i := 0; i < 5; i++ {
// 		require.Equal(t, randomWorkspaces[i].Id, workspaces[i].Id)
// 		require.Equal(t, randomWorkspaces[i].Name, workspaces[i].Name)
// 		require.Equal(t, randomWorkspaces[i].CreatorId, workspaces[i].CreatorId)
// 	}
// }

// func TestUpdateWorkspace(t *testing.T) {
// 	workspace := createRandomWorkspace(t)
// 	toUpdateWorkspace := editing.Workspace{
// 		Id:   workspace.Id,
// 		Name: storage.RandomString(6),
// 	}
// 	err := store.UpdateWorkspace(context.Background(), toUpdateWorkspace)
// 	require.NoError(t, err)
// }

// func TestDeleteWorkspace(t *testing.T) {
// 	workspace := createRandomWorkspace(t)
// 	err := store.DeleteWorkspace(context.Background(), workspace.Id)
// 	require.NoError(t, err)
// }

// func TestRemoveUserFormWorkspace(t *testing.T) {
// 	user := createRandomUser(t)
// 	newUser, err := store.GetUserByUsername(context.Background(), user.Username)
// 	require.NoError(t, err)
// 	workspace := createRandomWorkspace(t)
// 	err = store.AddUserToWorkspace(context.Background(), workspace.Id, []string{newUser.UserId})
// 	require.NoError(t, err)
// 	err = store.RemoveUserFromWorkspace(context.Background(), workspace.Id, newUser.UserId)
// 	require.NoError(t, err)
// }
