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
	wr  WorkspaceRepository
	cr  ChannelRepository
}

func NewService(log *zap.Logger, wr WorkspaceRepository, cr ChannelRepository) *service {
	return &service{
		log: log,
		wr:  wr,
		cr:  cr,
	}
}

type WorkspaceRepository interface {
	GetWorkspacesByUserId(userId string) ([]Workspace, error)
	GetWorkspaceByWorkspaceId(workspaceId string) (Workspace, error)
}

type ChannelRepository interface {
	GetAllChannelByWorkspaceId(workspaceId string)([]Channel,error)
	GetAllChannelByUserAndWorkspaceId(userId, workspaceId string) ([]Channel, error)
}

func (s *service) GetWorkspaceData(userId string) ([]Workspace, error) {
	workspaces, err := s.wr.GetWorkspacesByUserId(userId)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(workspaces); i++ {
		channels, err := s.cr.GetAllChannelByUserAndWorkspaceId(userId, workspaces[i].Id)
		if err != nil {
			return nil, err
		}
		workspaces[i].ChannelList = channels
	}
	return workspaces, nil
}

func (s *service) GetWorkspaceByWorkspaceId(workspaceId string) (Workspace, error) {
	workspace, err := s.wr.GetWorkspaceByWorkspaceId(workspaceId)
	if err != nil {
		return Workspace{}, err
	}

	channelList, err := s.cr.GetAllChannelByWorkspaceId(workspaceId)
	workspace.ChannelList = channelList
	
	if err != nil {
		return Workspace{}, err
	}
	return workspace, nil
}
