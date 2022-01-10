package listing

import (
	"context"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

type Service interface {
	GetWorkspaceData(ctx context.Context, userId string) ([]Workspace, error)
	GetWorkspaceByWorkspaceId(ctx context.Context, workspaceId string) (Workspace, error)
	GetChannelByChannelId(ctx context.Context, channelId string) (Channel, error)
	VerifyAndGetChannel(ctx context.Context, workspaceId string, channelId string) (Channel, error)
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
	if err == sql.ErrNoRows {
		return Workspace{}, errors.New("workspace.id.not.found")
	} else if err != nil {
		return Workspace{}, err
	}
	c, err := s.r.GetDefaultChannel(ctx, workspaceId)
	if err != nil && err != sql.ErrNoRows {
		return Workspace{}, err
	}
	if err != sql.ErrNoRows {
		w.ChannelList = []Channel{c}
	}
	return w, nil
}

func (s *service) GetChannelByChannelId(ctx context.Context, channelId string) (Channel, error) {
	c, err := s.r.GetChannelByChannelId(ctx, channelId)
	if err == sql.ErrNoRows {
		return Channel{}, errors.New("channel.id.not.found")
	} else if err != nil {
		return Channel{}, err
	}
	return c, nil
}

func (s *service) VerifyAndGetChannel(ctx context.Context, workspaceId string, channelId string) (Channel, error) {
	_, err := s.GetWorkspaceByWorkspaceId(ctx, workspaceId)
	if err != nil {
		return Channel{}, err
	}
	channel, err := s.GetChannelByChannelId(ctx, channelId)
	if err != nil {
		return Channel{}, err
	}
	channelList, err := s.r.GetChannelListByWorkspaceId(ctx, workspaceId)
	if err != nil {
		return Channel{}, err
	}
	exist := false
	for _, c := range channelList {
		if channelId == c.ChannelId {
			exist = true
			break
		}
	}
	if !exist {
		return Channel{}, errors.New("channel.id.not.found")
	}
	return channel, nil
}
