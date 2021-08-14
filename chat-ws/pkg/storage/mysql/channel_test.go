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

func TestCreateChannel(t *testing.T) {
	randUser := createRandomUser(t)
	user, err := store.GetUserByUsername(context.Background(), randUser.Username)
	require.NoError(t, err)
	workspace := createRandomWorkspace(t)
	c := adding.Channel{
		ChannelId:   uuid.New().String(),
		ChannelName: storage.RandomString(6),
		WorkspaceId: workspace.Id,
		CreatorId:   user.UserId,
		CreatedDate: time.Now(),
	}
	err = store.CreateChannel(context.Background(), c)
	require.NoError(t, err)
}

func TestAddUserToChannel(t *testing.T) {
	randUser := createRandomUser(t)
	user, err := store.GetUserByUsername(context.Background(), randUser.Username)
	require.NoError(t, err)
	channel := createRandomChannel(t)
	err = store.AddUserToChannel(context.Background(), channel.ChannelId, []string{user.UserId})
	require.NoError(t, err)
}

func TestGetChannelsByUserId(t *testing.T) {
	newUser := createRandomUser(t)
	user, err := store.GetUserByUsername(context.Background(), newUser.Username)
	require.NoError(t, err)
	workspace := createRandomWorkspace(t)
	err = store.AddNewUserToWorkspace(context.Background(), workspace.Id, []string{user.UserId})
	require.NoError(t, err)
	randomChannels := make([]adding.Channel, 5)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		randomChannels[i] = adding.Channel{
			ChannelId:   uuid.New().String(),
			ChannelName: storage.RandomString(6),
			WorkspaceId: workspace.Id,
			CreatorId:   user.UserId,
			CreatedDate: time.Now(),
		}
		_, err := store.CreateNewChannel(context.Background(), randomChannels[i], user.UserId)
		require.NoError(t, err)
	}
	channels, err := store.GetChannelsByUserId(context.Background(), user.UserId)
	require.NoError(t, err)
	for i := 0; i < 5; i++ {
		require.Equal(t, randomChannels[i].ChannelId, channels[i+1].ChannelId)
		require.Equal(t, randomChannels[i].ChannelName, channels[i+1].ChannelName)
		require.Equal(t, randomChannels[i].CreatorId, channels[i+1].CreatorId)
		require.WithinDuration(t, randomChannels[i].CreatedDate, channels[i+1].CreatedDate, time.Second)
	}
	require.NoError(t, err)
}

func TestGetClientListByChannelId(t *testing.T) {
	channel := createRandomChannel(t)
	userIdList := make([]string, 5)
	for i := 0; i < 5; i++ {
		newUser := createRandomUser(t)
		user, err := store.GetUserByUsername(context.Background(), newUser.Username)
		require.NoError(t, err)
		userIdList[i] = user.UserId
	}
	err := store.AddUserToChannel(context.Background(), channel.ChannelId, userIdList)
	require.NoError(t, err)
	userList, err := store.GetUserListByChannelId(context.Background(), channel.ChannelId)
	require.NoError(t, err)
	require.Equal(t, len(userIdList), len(userList)-1)
}

func TestUpdateChannel(t *testing.T) {
	channel := createRandomChannel(t)
	newChannel := editing.Channel{
		ChannelId:   channel.ChannelId,
		ChannelName: storage.RandomString(6),
	}
	err := store.UpdateChannel(context.Background(), newChannel)
	require.NoError(t, err)
}

func TestDeleteChannel(t *testing.T) {
	channel := createRandomChannel(t)
	err := store.DeleteChannel(context.Background(), channel.ChannelId)
	require.NoError(t, err)
}

func TestRemoveUserFromChannel(t *testing.T) {
	channel := createRandomChannel(t)
	userList, err := store.GetUserListByChannelId(context.Background(), channel.ChannelId)
	require.NoError(t, err)
	require.NotEmpty(t, userList)
	err = store.RemoveUserFromChannel(context.Background(), channel.ChannelId, userList[0].UserId)
	require.NoError(t, err)
}
