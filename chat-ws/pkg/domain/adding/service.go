package adding

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	CreateWorkspace(w Workspace, userId string) (Workspace, error)
	AddUserToWorkspace(workspaceId string, userIdList []string) error
	CreateChannel(workspaceId, channelName, userId string) error
	AddUserToChannel(channelId string, userIdList []string) error
}

type WorkspaceRepository interface {
	CreateWorkspace(w Workspace) error
	AddUserToWorkspace(workspaceId string, userIdList []string) error
}

type ChannelRepository interface {
	CreateChannel(channel *Channel) error
	AddUserToChannel(channelId string, userId []string) error
	GetDefaultChannelByWorkspaceId(workspaceId string) (Channel, error)
}

type service struct {
	wr  WorkspaceRepository
	cr  ChannelRepository
	log *zap.Logger
}

func NewService(log *zap.Logger, wr WorkspaceRepository, cr ChannelRepository) *service {
	return &service{
		wr:  wr,
		cr:  cr,
		log: log,
	}
}

func (s *service) CreateWorkspace(w Workspace, userId string) (Workspace, error) {
	w.Id = uuid.New().String()
	w.CreatorId = userId
	err := s.wr.CreateWorkspace(w)
	if err != nil {
		return Workspace{}, err
	}
	err = s.AddUserToWorkspace(w.Id, []string{userId})
	if err != nil {
		return Workspace{}, err
	}
	err = s.CreateChannel(w.Id, "General", userId)
	if err != nil {
		return Workspace{}, err
	}
	return w, nil
}

func (s *service) AddUserToWorkspace(workspaceId string, userIdList []string) error {
	err := s.wr.AddUserToWorkspace(workspaceId, userIdList)
	if err != nil {
		return err
	}
	c, err := s.cr.GetDefaultChannelByWorkspaceId(workspaceId)
	if err != nil {
		return err
	}
	err = s.AddUserToChannel(c.ChannelId, userIdList)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateChannel(workspaceId, channelName, userId string) error {
	c := NewChannel(uuid.New().String(), channelName, workspaceId)
	err := s.cr.CreateChannel(c)
	if err != nil {
		return err
	}
	err = s.AddUserToChannel(c.ChannelId, []string{userId})
	if err != nil {
		return err
	}
	return nil
}

func (s *service) AddUserToChannel(channelId string, userIdList []string) error {
	return s.cr.AddUserToChannel(channelId, userIdList)
}
