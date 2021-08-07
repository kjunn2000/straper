package listing

import (
	"go.uber.org/zap"
)

type Service interface {
	GetWorkspaceData(userId string) ([]Workspace, error)
	GetWorkspaceByWorkspaceId(workspaceId string) (Workspace, error)
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

type Repository interface {
	GetWorkspacesByUserId(userId string) ([]Workspace, error)
	GetWorkspaceByWorkspaceId(workspaceId string) (Workspace, error)
	GetAllChannelByWorkspaceId(workspaceId string) ([]Channel, error)
	GetAllChannelByUserAndWorkspaceId(userId, workspaceId string) ([]Channel, error)
}

func (s *service) GetWorkspaceData(userId string) ([]Workspace, error) {
	workspaces, err := s.r.GetWorkspacesByUserId(userId)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(workspaces); i++ {
		channels, err := s.r.GetAllChannelByUserAndWorkspaceId(userId, workspaces[i].Id)
		if err != nil {
			return nil, err
		}
		workspaces[i].ChannelList = channels
	}
	return workspaces, nil
}

func (s *service) GetWorkspaceByWorkspaceId(workspaceId string) (Workspace, error) {
	workspace, err := s.r.GetWorkspaceByWorkspaceId(workspaceId)
	if err != nil {
		return Workspace{}, err
	}

	channelList, err := s.r.GetAllChannelByWorkspaceId(workspaceId)
	workspace.ChannelList = channelList

	if err != nil {
		return Workspace{}, err
	}
	return workspace, nil
}
