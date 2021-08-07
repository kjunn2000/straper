package mysql

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
	"github.com/stretchr/testify/require"
)

func TestCreateChannel(t *testing.T) {
	channel := adding.Channel{
		ChannelId:   uuid.New().String(),
		ChannelName: "Happy Channel 2",
		WorkspaceId: "4ce6120b-dfd4-4460-b41c-af4f7d8e0efb",
	}
	newChannel, err := store.CreateChannel(channel, "U000001")
	require.NoError(t, err)
	require.Equal(t, channel.ChannelId, newChannel.ChannelId)
	require.Equal(t, channel.ChannelName, newChannel.ChannelName)
	require.Equal(t, channel.WorkspaceId, newChannel.WorkspaceId)
}
