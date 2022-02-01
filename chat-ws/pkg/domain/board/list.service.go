package board

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

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
	var updateListParams UpdateListParams
	if err := json.Unmarshal(bytePayload, &updateListParams); err != nil {
		return err
	}
	return service.store.UpdateTaskList(ctx, updateListParams)
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
