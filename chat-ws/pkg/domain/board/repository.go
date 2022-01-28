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

	CreateTaskList(ctx context.Context, taskList TaskList) error
	UpdateTaskList(ctx context.Context, taskList UpdateListParams) error
	UpdateTaskListOrder(ctx context.Context, listId string, orderIndex int) error
	DeleteTaskList(ctx context.Context, listId string) error

	CreateCard(ctx context.Context, card Card) error
	UpdateCard(ctx context.Context, params UpdateCardParams) error
	UpdateCardDueDate(ctx context.Context, params UpdateCardDueDateParams) error
	UpdateCardOrder(ctx context.Context, cardId string, orderIndex int, listId string, updateListId bool) error
	DeleteCard(ctx context.Context, cardId string) error

	GetUserFromCard(ctx context.Context, cardId string) ([]string, error)
	AddUserListToCard(ctx context.Context, cardId string, userIdList []string) error
	DeleteUserFromCard(ctx context.Context, cardId, userId string) error

	GetChecklistItemsByCardId(ctx context.Context, cardId string) ([]string, error)
	CreateChecklistItem(ctx context.Context, checklistItem CardChecklistItem) error
	UpdateChecklistItem(ctx context.Context, checklistItem CardChecklistItem) error
	DeleteChecklistItem(ctx context.Context, itemId string) error
}
