package board

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	ws "github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
	"go.uber.org/zap"
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
	var orderListParams OrderListParams
	if err := json.Unmarshal(bytePayload, &orderListParams); err != nil {
		return err
	}
	taskLists, err := service.store.GetTaskListsByBoardId(ctx, orderListParams.BoardId)
	if err != nil {
		return err
	}
	target := taskLists[orderListParams.OldListIndex]
	taskLists = append(taskLists[:orderListParams.OldListIndex], taskLists[orderListParams.OldListIndex+1:]...)
	taskLists = append(taskLists[:orderListParams.NewListIndex],
		append([]TaskList{target}, taskLists[orderListParams.NewListIndex:]...)...)
	for i, taskList := range taskLists {
		if err := service.store.UpdateTaskListOrder(ctx, taskList.ListId, i+1); err != nil {
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
	fmt.Println(card)
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
	var orderCardParams OrderCardParams
	if err := json.Unmarshal(bytePayload, &orderCardParams); err != nil {
		return err
	}
	cardList, err := service.store.GetCardListByListId(ctx, orderCardParams.SourceListId)
	if err != nil {
		return err
	}

	target := cardList[orderCardParams.OldCardIndex]

	cardList = append(cardList[:orderCardParams.OldCardIndex], cardList[orderCardParams.OldCardIndex+1:]...)
	if orderCardParams.SourceListId == orderCardParams.DestListId {
		cardList = append(cardList[:orderCardParams.NewCardIndex],
			append([]Card{target}, cardList[orderCardParams.NewCardIndex:]...)...)
	}
	if err := service.updateCardListOrderIndex(ctx, cardList, orderCardParams.SourceListId); err != nil {
		return err
	}

	if orderCardParams.SourceListId != orderCardParams.DestListId {
		destCardList, err := service.store.GetCardListByListId(ctx, orderCardParams.DestListId)
		if err != nil {
			return err
		}
		destCardList = append(destCardList[:orderCardParams.NewCardIndex],
			append([]Card{target}, destCardList[orderCardParams.NewCardIndex:]...)...)
		if err := service.updateCardListOrderIndex(ctx, destCardList, orderCardParams.DestListId); err != nil {
			return err
		}
	}
	return nil
}

func (service *service) updateCardListOrderIndex(ctx context.Context, cardList []Card, listId string) error {
	for i, card := range cardList {
		if err := service.store.UpdateCardOrder(ctx, card.CardId, i+1, listId, card.ListId != listId); err != nil {
			return err
		}
	}
	return nil
}
