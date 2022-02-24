package deleting

import (
	"context"

	"go.uber.org/zap"
)

type Service interface {
	DeleteWorkspace(ctx context.Context, workspaceId string) error
	LeaveWorkspace(ctx context.Context, workspaceId, userId string) error
	DeleteChannel(ctx context.Context, channelId string) error
	LeaveChannel(ctx context.Context, channelId, userId string) error
}

type SeaweedfsClient interface {
	DeleteFile(ctx context.Context, fid string) error
}

type service struct {
	log *zap.Logger
	r   Repository
	sc  SeaweedfsClient
}

func NewService(log *zap.Logger, r Repository, sc SeaweedfsClient) *service {
	return &service{
		log: log,
		r:   r,
		sc:  sc,
	}
}

type Channel struct {
	ChannelId string `json:"channel_id" db:"channel_id"`
}

func (s *service) DeleteWorkspace(ctx context.Context, workspaceId string) error {
	fids, err := s.r.GetFidsByWorkspaceId(ctx, workspaceId)
	if err != nil {
		return err
	}
	aFids, err := s.r.GetAttachmentFidsByWorkspaceId(ctx, workspaceId)
	if err != nil {
		return err
	}
	for _, fid := range append(fids, aFids...) {
		if err := s.sc.DeleteFile(ctx, fid); err != nil {
			return err
		}
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

func (s *service) DeleteChannel(ctx context.Context, channelId string) error {
	err := s.r.DeleteChannel(ctx, channelId)
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
