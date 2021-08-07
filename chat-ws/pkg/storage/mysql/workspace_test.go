package mysql

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
	"github.com/stretchr/testify/require"
)

func TestCreateWorkspace(t *testing.T) {
	workspaceId := uuid.New().String()
	workspace := adding.Workspace{
		Id:        workspaceId,
		Name:      "Dreaming",
		CreatorId: "U000001",
	}
	channel := adding.Channel{
		ChannelId:   uuid.New().String(),
		ChannelName: "Dreaming Channel",
		WorkspaceId: workspaceId,
	}
	workspace, err := store.CreateWorkspace(workspace, channel, "U000001")
	require.NoError(t, err)
}

func TestAddUserToWorksapce(t *testing.T) {
	err := store.AddUserToWorkspace("0f4d1de2-5eb4-40c9-a1dc-867dfc61440c", []string{"U000002"})
	require.NoError(t, err)
}
