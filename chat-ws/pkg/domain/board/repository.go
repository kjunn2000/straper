package board

import (
	"context"

	ws "github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
)

type Repository interface {
	GetUserListByWorkspaceId(ctx context.Context, workspaceId string) ([]ws.UserData, error)
	GetTaskBoardByWorkspaceId(ctx context.Context, workspaceId string) (TaskBoard, error)
	GetTaskListsByBoardId(ctx context.Context, boardId string) ([]TaskList, error)
	GetCardListByListId(ctx context.Context, listId string) ([]Card, error)
}
