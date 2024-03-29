package board

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"time"

	ws "github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
	"go.uber.org/zap"
)

type Service interface {
	HandleBroadcast(ctx context.Context, msg *ws.Message, publishPubSub func(context.Context, *ws.Message) error) error
	GetTaskBoardData(ctx context.Context, workspaceId string) (TaskBoardDataResponse, error)
	GetBroadcastUserListByMessageType(ctx context.Context, msg *ws.Message) ([]ws.UserData, error)
	GetCardComments(ctx context.Context, cardId string, param PaginationCommentParam) ([]CardComment, error)
}

type SeaweedfsClient interface {
	SaveFile(ctx context.Context, reader io.Reader) (string, error)
	GetFile(ctx context.Context, fid string) ([]byte, error)
	DeleteFile(ctx context.Context, fid string) error
}

type PaginationService interface {
	DecodeCursor(encodedCursor string) (res time.Time, uuid string, err error)
	EncodeCursor(t time.Time, uuid string) string
}

type service struct {
	log   *zap.Logger
	store Repository
	sc    SeaweedfsClient
	ps    PaginationService
}

func NewService(log *zap.Logger, store Repository, sc SeaweedfsClient, ps PaginationService) *service {
	return &service{
		log:   log,
		store: store,
		sc:    sc,
		ps:    ps,
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
		for i, card := range cardList {
			memberList, err := service.getCardMember(ctx, card.CardId)
			if err != nil {
				return TaskBoardDataResponse{}, err
			}
			cardList[i].MemberList = memberList
			checklist, err := service.store.GetChecklistItemsByCardId(ctx, card.CardId)
			if err != nil {
				return TaskBoardDataResponse{}, err
			}
			cardList[i].Checklist = checklist
		}
		taskLists[i].CardList = cardList
	}
	taskBoardData.TaskLists = taskLists
	return taskBoardData, err
}

func (service *service) getCardMember(ctx context.Context, cardId string) ([]string, error) {
	userIdList, err := service.store.GetUserFromCard(ctx, cardId)
	if err != nil {
		return []string{}, err
	}
	return userIdList, err
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
	case BoardUpdateCardDueDate:
		if err := service.handleUpdateCardDueDate(ctx, bytePayload); err != nil {
			return err
		}
	case BoardUpdateCardIssueLink:
		if err := service.handleUpdateIssueLink(ctx, bytePayload); err != nil {
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
	case BoardCardAddMembers:
		if err := service.handleCardAddMembers(ctx, bytePayload); err != nil {
			return err
		}
	case BoardCardRemoveMember:
		if err := service.handleCardRemoveMember(ctx, bytePayload); err != nil {
			return err
		}
	case BoardCardAddChecklistItem:
		if newPayload, err := service.handleCardAddChecklistItem(ctx, bytePayload); err != nil {
			return err
		} else {
			if err := msg.Payload.UnmarshalJSON(newPayload); err != nil {
				return err
			}
		}
	case BoardCardUpdateChecklistItem:
		if err := service.handleCardUpdateChecklistItem(ctx, bytePayload); err != nil {
			return err
		}
	case BoardCardDeleteChecklistItem:
		if err := service.handleCardDeleteChecklistItem(ctx, bytePayload); err != nil {
			return err
		}
	case BoardCardAddComment:
		if newPayload, err := service.handleBoardAddComment(ctx, bytePayload); err != nil {
			return err
		} else {
			if err := msg.Payload.UnmarshalJSON(newPayload); err != nil {
				return err
			}
		}
	case BoardCardEditComment:
		if err := service.handleCardEditComment(ctx, bytePayload); err != nil {
			return err
		}
	case BoardCardDeleteComment:
		if err := service.handleCardDeleteComment(ctx, bytePayload); err != nil {
			return err
		}
	}
	if err := publishPubSub(ctx, msg); err != nil {
		return err
	}
	return nil
}

func (s *service) GetBroadcastUserListByMessageType(ctx context.Context, msg *ws.Message) ([]ws.UserData, error) {
	return s.store.GetUserListByWorkspaceId(ctx, msg.WorkspaceId)
}

func (s *service) GetCardComments(ctx context.Context, cardId string, param PaginationCommentParam) ([]CardComment, error) {
	if param.Cursor != "" {
		time, uuid, err := s.ps.DecodeCursor(param.Cursor)
		if err != nil {
			return []CardComment{}, err
		}
		param.CreatedTime = time
		param.Id = uuid
	}
	msgs, err := s.store.GetCardComments(ctx, cardId, param)
	if err == sql.ErrNoRows {
		return []CardComment{}, errors.New("no.card.comment.available")
	} else if err != nil {
		return []CardComment{}, err
	}
	for i, msg := range msgs {
		if msg.Type == TypeFile {
			bytesData, err := s.sc.GetFile(ctx, msg.Content)
			if err != nil {
				return []CardComment{}, err
			}
			msg.FileBytes = bytesData
			msgs[i] = msg
		}
		userDetail, err := s.store.GetBoardUserInfoByUserId(ctx, msg.CreatorId)
		if err != nil {
			return []CardComment{}, err
		} else {
			msgs[i].UserDetail = userDetail
		}
		msgs[i].Cursor = s.ps.EncodeCursor(msg.CreatedDate, msg.CommentId)
	}
	return msgs, nil
}

func (s *service) deleteSeaweedfsFiles(ctx context.Context, fileComment []CardComment) error {
	for _, comment := range fileComment {
		if err := s.sc.DeleteFile(ctx, comment.Content); err != nil {
			return err
		}
	}
	return nil
}
