package adding

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	CreateWorkspace(w Workspace, userId string) error
	AddUserToWorkspace(workspaceId string, userIdList []string) error
	CreateChannel(workspaceId, channelName string) (Channel, error)
	AddUserToChannel(channelId string, userIdList []string) error
}

type WorkspaceRepository interface {
	CreateWorkspace(w Workspace) error
	AddUserToWorkspace(workspaceId string, userIdList []string) error
}

type ChannelRepository interface {
	CreateChannel(channel *Channel) error
	AddUserToChannel(channelId string, userId []string) error
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

func (s *service) CreateWorkspace(w Workspace, userId string) error {
	w.Id = uuid.New().String()
	err := s.wr.CreateWorkspace(w)
	if err != nil {
		return err
	}
	err = s.AddUserToWorkspace(w.Id, []string{userId})
	if err != nil {
		return err
	}
	channel, err := s.CreateChannel(w.Id, "General")
	if err != nil {
		return err
	}
	err = s.AddUserToChannel(channel.ChannelId, []string{userId})
	if err != nil {
		return err
	}
	return nil
}

func (s *service) AddUserToWorkspace(workspaceId string, userIdList []string) error {
	return s.wr.AddUserToWorkspace(workspaceId, userIdList)
}

func (s *service) CreateChannel(workspaceId, channelName string) (Channel, error) {
	c := NewChannel(uuid.New().String(), channelName, workspaceId)
	err := s.cr.CreateChannel(c)
	if err != nil {
		return Channel{}, err
	}
	return *c, nil
}

func (s *service) AddUserToChannel(channelId string, userIdList []string) error {
	return s.cr.AddUserToChannel(channelId, userIdList)
}
