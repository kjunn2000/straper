package board

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	ws "github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
	"go.uber.org/zap"
)

const (
	BoardAddList    = "BOARD_ADD_LIST"
	BoardUpdateList = "BOARD_UPDATE_LIST"
	BoardDeleteList = "BOARD_DELETE_LIST"
	BoardOrderList  = "BOARD_ORDER_LIST"

	BoardAddCard    = "BOARD_ADD_CARD"
	BoardUpdateCard = "BOARD_UPDATE_CARD"
	BoardDeleteCard = "BOARD_DELETE_CARD"
	BoardOrderCard  = "BOARD_ORDER_CARD"
)

type Service interface {
	GetTaskBoardData(ctx context.Context, workspaceId string) (TaskBoardDataResponse, error)
	HandleBroadcast(ctx context.Context, msg *ws.Message, publishPubSub func(context.Context, *ws.Message) error) error
	GetBoarcastUserListByMessageType(ctx context.Context, msg *ws.Message) ([]ws.UserData, error)
}

type service struct {
	log   *zap.Logger
	store Repository
}

func NewService(log *zap.Logger, store Repository) *service {
	return &service{
		log:   log,
		store: store,
	}
}

func (service *service) GetTaskBoardData(ctx context.Context, workspaceId string) (TaskBoardDataResponse, error) {
	taskBoard, err := service.store.GetTaskBoardByWorkspaceId(ctx, workspaceId)
	if err != nil {
		return TaskBoardDataResponse{}, err
	}
	var taskBoardData TaskBoardDataResponse
	taskBoardData.TaskBoard = taskBoard
	taskLists, err := service.store.GetTaskListsByBoardId(ctx, taskBoard.BoardId)
	if err != nil && err != sql.ErrNoRows {
		return TaskBoardDataResponse{}, err
	}
	for i, taskList := range taskLists {
		cardList, err := service.store.GetCardListByListId(ctx, taskList.ListId)
		if err != nil && err != sql.ErrNoRows {
			return TaskBoardDataResponse{}, err
		}
		taskLists[i].CardList = cardList
	}
	taskBoardData.TaskLists = taskLists
	return taskBoardData, err
}

func (service *service) HandleBroadcast(ctx context.Context, msg *ws.Message, publishPubSub func(context.Context, *ws.Message) error) error {
	bytePayload, err := msg.Payload.MarshalJSON()
	if err != nil {
		return err
	}
	switch msg.MessageType {
	case BoardAddList:
		if newPayload, err := service.handleAddList(ctx, bytePayload); err != nil {
			return err
		} else {
			if err := msg.Payload.UnmarshalJSON(newPayload); err != nil {
				return err
			}
		}
	case BoardUpdateList:
		if err := service.handleUpdateList(ctx, bytePayload); err != nil {
			return err
		}
	case BoardDeleteList:
		if err := service.handleDeleteList(ctx, bytePayload); err != nil {
			return err
		}
	case BoardOrderList:
		if err := service.handleOrderList(ctx, bytePayload); err != nil {
			return err
		}
	case BoardAddCard:
		if newPayload, err := service.handleAddCard(ctx, bytePayload); err != nil {
			return err
		} else {
			if err := msg.Payload.UnmarshalJSON(newPayload); err != nil {
				return err
			}
		}
	case BoardUpdateCard:
		if err := service.handleUpdateCard(ctx, bytePayload); err != nil {
			return err
		}
	case BoardDeleteCard:
		if err := service.handleDeleteCard(ctx, bytePayload); err != nil {
			return err
		}
	case BoardOrderCard:
		if err := service.handleOrderCard(ctx, bytePayload); err != nil {
			return err
		}
	}
	if err := publishPubSub(ctx, msg); err != nil {
		return err
	}
	return nil
}

func (s *service) GetBoarcastUserListByMessageType(ctx context.Context, msg *ws.Message) ([]ws.UserData, error) {
	return s.store.GetUserListByWorkspaceId(ctx, msg.WorkspaceId)
}

func (service *service) handleAddList(ctx context.Context, bytePayload []byte) ([]byte, error) {
	var taskList TaskList
	if err := json.Unmarshal(bytePayload, &taskList); err != nil {
		return []byte{}, err
	}
	listId, _ := uuid.NewRandom()
	taskList.ListId = listId.String()
	if err := service.store.CreateTaskList(ctx, taskList); err != nil {
		return []byte{}, err
	}
	newPayload, _ := json.Marshal(taskList)
	return newPayload, nil
}

func (service *service) handleUpdateList(ctx context.Context, bytePayload []byte) error {
	var taskList TaskList
	if err := json.Unmarshal(bytePayload, &taskList); err != nil {
		return err
	}
	return service.store.UpdateTaskList(ctx, taskList)
}

func (service *service) handleDeleteList(ctx context.Context, bytePayload []byte) error {
	var listId string
	if err := json.Unmarshal(bytePayload, &listId); err != nil {
		return err
	}
	return service.store.DeleteTaskList(ctx, listId)
}

func (service *service) handleOrderList(ctx context.Context, bytePayload []byte) error {
	var taskLists []TaskList
	if err := json.Unmarshal(bytePayload, &taskLists); err != nil {
		return err
	}
	for _, taskList := range taskLists {
		if err := service.store.UpdateTaskList(ctx, taskList); err != nil {
			return err
		}
	}
	return nil
}

func (service *service) handleAddCard(ctx context.Context, bytePayload []byte) ([]byte, error) {
	var card Card
	if err := json.Unmarshal(bytePayload, &card); err != nil {
		return []byte{}, err
	}
	cardId, _ := uuid.NewRandom()
	card.CardId = cardId.String()
	card.Status = NoAssign
	card.Priority = NoPriority
	card.CreatedDate = time.Now()
	card.DueDate = time.Now().Add(time.Hour * 24 * 7)
	if err := service.store.CreateCard(ctx, card); err != nil {
		return []byte{}, err
	}
	newPayload, _ := json.Marshal(card)
	return newPayload, nil
}

func (service *service) handleUpdateCard(ctx context.Context, bytePayload []byte) error {
	var updateCardParams UpdateCardParams
	if err := json.Unmarshal(bytePayload, &updateCardParams); err != nil {
		return err
	}
	return service.store.UpdateCard(ctx, updateCardParams)
}

func (service *service) handleDeleteCard(ctx context.Context, bytePayload []byte) error {
	var cardId string
	if err := json.Unmarshal(bytePayload, &cardId); err != nil {
		return err
	}
	return service.store.DeleteTaskList(ctx, cardId)
}

func (service *service) handleOrderCard(ctx context.Context, bytePayload []byte) error {
	var updateCardOrderParams []UpdateCardOrderParams
	if err := json.Unmarshal(bytePayload, &updateCardOrderParams); err != nil {
		return err
	}
	for _, updateCardOrderParam := range updateCardOrderParams {
		if err := service.store.UpdateCardOrder(ctx, updateCardOrderParam); err != nil {
			return err
		}
	}
	return nil
}
