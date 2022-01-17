package board

import (
	"context"
	"database/sql"

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
	return nil
}

func (s *service) GetBoarcastUserListByMessageType(ctx context.Context, msg *ws.Message) ([]ws.UserData, error) {
	// message, err := s.convertByteArrayToMessage(ctx, msg)
	// if err != nil {
	// 	return []ws.UserData{}, err
	// }
	// return s.store.GetUserListByWorkspaceId(ctx, message.ChannelId)
	return []ws.UserData{}, nil
}

// func (s *service) convertByteArrayToMessage(ctx context.Context, msg *ws.Message) (Message, error) {
// bytePayload, err := msg.Payload.MarshalJSON()
// if err != nil {
// 	return Message{}, err
// }
// var message Message
// if err := json.Unmarshal(bytePayload, &message); err != nil {
// 	return Message{}, err
// }
// return Message{}, nil
// }
