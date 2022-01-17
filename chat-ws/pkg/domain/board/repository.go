package board

import (
	"context"

	ws "github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
)

type Repository interface {
	GetUserListByWorkspaceId(ctx context.Context, workspaceId string) ([]ws.UserData, error)
}
