package deleting

import (
	"context"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

type Service interface {
	DeleteWorkspace(ctx context.Context, workspaceId, userId string) error
	LeaveWorkspace(ctx context.Context, workspaceId, userId string) error
	DeleteChannel(ctx context.Context, channelId, userId string) error
	LeaveChannel(ctx context.Context, channelId, userId string) error
}

type service struct {
	log *zap.Logger
	r   Repository
}

func NewService(log *zap.Logger, r Repository) *service {
	return &service{
		log: log,
		r:   r,
	}
}

type Channel struct {
	ChannelId string `json:"channel_id" db:"channel_id"`
}

func (s *service) DeleteWorkspace(ctx context.Context, workspaceId, userId string) error {
	w, err := s.r.GetWorkspaceByWorkspaceId(ctx, workspaceId)
	if err != nil {
		return err
	}
	if w.CreatorId != userId {
		return errors.New("invalid.delete.workspace.authority")
	}
	err = s.r.DeleteWorkspace(ctx, workspaceId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) LeaveWorkspace(ctx context.Context, workspaceId, userId string) error {
	err := s.r.RemoveUserFromWorkspace(ctx, workspaceId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteChannel(ctx context.Context, channelId, userId string) error {
	channel, err := s.r.GetChannelByChannelId(ctx, channelId)
	if err == sql.ErrNoRows {
		return errors.New("invalid.channel.id")
	}
	if err != nil {
		return err
	}
	if channel.CreatorId != userId {
		return errors.New("invalid.delete.workspace.authority")
	}
	err = s.r.DeleteChannel(ctx, channelId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) LeaveChannel(ctx context.Context, channelId, userId string) error {
	err := s.r.RemoveUserFromChannel(ctx, channelId, userId)
	if err != nil {
		return err
	}
	return nil
}
