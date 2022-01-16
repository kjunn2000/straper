package board

import (
	"context"

	ws "github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
	"go.uber.org/zap"
)

type Service interface {
	GetBoardData(ctx context.Context, workspaceId string)
	HandleBroadcast(ctx context.Context, msg *ws.Message, publishPubSub func(context.Context, *ws.Message) error) error
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

func (service *service) GetBoardData(ctx context.Context, workspaceId string) {

}

func (service *service) HandleBroadcast(ctx context.Context, msg *ws.Message, publishPubSub func(context.Context, *ws.Message) error) error {
	return nil
}
