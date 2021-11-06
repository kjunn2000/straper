package listing

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

type Service interface {
	GetWorkspaceData(ctx context.Context, userId string) ([]Workspace, error)
	GetWorkspaceByWorkspaceId(ctx context.Context, workspaceId string) (Workspace, error)
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

func (s *service) GetWorkspaceData(ctx context.Context, userId string) ([]Workspace, error) {
	workspaceList, err := s.r.GetWorkspacesByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	channelList, err := s.r.GetChannelsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	workspaceData, err := s.generateWorkspaceWithChannelData(workspaceList, channelList)
	if err != nil {
		return nil, err
	}
	return workspaceData, nil
}

func (s *service) generateWorkspaceWithChannelData(workspaceList []Workspace, channelList []Channel) ([]Workspace, error) {
	workspaceMap := make(map[string]Workspace)
	for _, workspace := range workspaceList {
		workspaceMap[workspace.Id] = workspace
	}
	for _, channel := range channelList {
		c, ok := workspaceMap[channel.WorkspaceId]
		if !ok {
			return nil, errors.New("workspace.and.channel.data.not.match")
		}
		c.ChannelList = append(c.ChannelList, channel)
		workspaceMap[channel.WorkspaceId] = c
	}
	workspaceData := make([]Workspace, 0)
	for _, workspace := range workspaceMap {
		workspaceData = append(workspaceData, workspace)
	}
	return workspaceData, nil
}

func (s *service) GetWorkspaceByWorkspaceId(ctx context.Context, workspaceId string) (Workspace, error) {
	w, err := s.r.GetWorkspaceByWorkspaceId(ctx, workspaceId)
	if err != nil {
		return Workspace{}, err
	}
	c, err := s.r.GetDefaultChannel(ctx, workspaceId)
	if err != nil {
		return Workspace{}, err
	}
	w.ChannelList = []Channel{c}
	return w, nil
}
