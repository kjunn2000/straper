package mysql

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
// 	"github.com/kjunn2000/straper/chat-ws/pkg/storage"
// 	"github.com/stretchr/testify/require"
// )

// func createRandomWorkspace(t *testing.T) adding.Workspace {
// 	randUser := createRandomUser(t)
// 	user, err := store.GetUserByUsername(context.Background(), randUser.Username)
// 	require.NoError(t, err)
// 	newId, err := uuid.NewRandom()
// 	require.NoError(t, err)
// 	workspaceId := newId.String()
// 	workspace := adding.Workspace{
// 		Id:          workspaceId,
// 		Name:        storage.RandomString(6),
// 		CreatorId:   user.UserId,
// 		CreatedDate: time.Now(),
// 	}
// 	channel := adding.Channel{
// 		ChannelId:   uuid.New().String(),
// 		ChannelName: "General",
// 		WorkspaceId: workspaceId,
// 		CreatorId:   user.UserId,
// 		CreatedDate: time.Now(),
// 	}
// 	newWorkspace, err := store.CreateNewWorkspace(context.Background(), workspace, channel, user.UserId)
// 	require.NoError(t, err)

// 	require.Equal(t, workspace.Id, newWorkspace.Id)
// 	require.Equal(t, workspace.Name, newWorkspace.Name)
// 	require.Equal(t, workspace.CreatorId, newWorkspace.CreatorId)

// 	return newWorkspace
// }

// func createRandomChannel(t *testing.T) adding.Channel {
// 	randUser := createRandomUser(t)
// 	user, err := store.GetUserByUsername(context.Background(), randUser.Username)
// 	require.NoError(t, err)
// 	workspace := createRandomWorkspace(t)
// 	c := adding.Channel{
// 		ChannelId:   uuid.New().String(),
// 		ChannelName: storage.RandomString(6),
// 		WorkspaceId: workspace.Id,
// 		CreatorId:   user.UserId,
// 		CreatedDate: time.Now(),
// 	}
// 	err = store.AddUserToWorkspace(context.Background(), workspace.Id, []string{user.UserId})
// 	require.NoError(t, err)
// 	newChannel, err := store.CreateNewChannel(context.Background(), c, user.UserId)
// 	require.NoError(t, err)
// 	require.Equal(t, c.ChannelId, newChannel.ChannelId)
// 	require.Equal(t, c.ChannelName, newChannel.ChannelName)
// 	require.Equal(t, c.WorkspaceId, newChannel.WorkspaceId)
// 	return c
// }

// func TestCreateNewWorkspace(t *testing.T) {
// 	createRandomWorkspace(t)
// }

// func TestAddNewUserToWorksapce(t *testing.T) {
// 	randUser := createRandomUser(t)
// 	user, err := store.GetUserByUsername(context.Background(), randUser.Username)
// 	require.NoError(t, err)
// 	workspace := createRandomWorkspace(t)
// 	err = store.AddNewUserToWorkspace(context.Background(), workspace.Id, []string{user.UserId})
// 	require.NoError(t, err)
// }
