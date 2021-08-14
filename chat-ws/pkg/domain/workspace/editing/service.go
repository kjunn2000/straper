package editing

import (
	"context"

	"go.uber.org/zap"
)

type Service interface {
	UpdateWorkspace(ctx context.Context, workspace Workspace) error
	UpdateChannel(ctx context.Context, channel Channel) error
}

type service struct {
	log *zap.Logger
	wr  Repository
}

func NewService(log *zap.Logger, wr Repository) *service {
	return &service{
		log: log,
		wr:  wr,
	}
}

func (s *service) UpdateWorkspace(ctx context.Context, workspace Workspace) error {
	return s.wr.UpdateWorkspace(ctx, workspace)
}

func (s *service) UpdateChannel(ctx context.Context, channel Channel) error {
	return s.wr.UpdateChannel(ctx, channel)
}
