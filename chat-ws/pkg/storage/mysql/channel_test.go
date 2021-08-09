package mysql

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/storage"
	"github.com/stretchr/testify/require"
)

func TestCreateChannel(t *testing.T) {
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
	_, err = store.CreateNewWorkspace(workspace, channel, user.UserId)
	require.NoError(t, err)
	c := adding.Channel{
		ChannelId:   uuid.New().String(),
		ChannelName: storage.RandomString(6),
		WorkspaceId: workspaceId,
	}
	newChannel, err := store.CreateNewChannel(c, user.UserId)
	require.NoError(t, err)
	require.Equal(t, c.ChannelId, newChannel.ChannelId)
	require.Equal(t, c.ChannelName, newChannel.ChannelName)
	require.Equal(t, c.WorkspaceId, newChannel.WorkspaceId)
}
