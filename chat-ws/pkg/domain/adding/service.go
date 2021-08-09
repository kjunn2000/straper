package adding

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	CreateWorkspace(w Workspace, userId string) (Workspace, error)
	AddUserToWorkspace(workspaceId string, userIdList []string) error
	CreateChannel(workspaceId, channelName, userId string) (Channel, error)
	AddUserToChannel(channelId string, userIdList []string) error
}

type Repository interface {
	CreateNewWorkspace(w Workspace, c Channel, userId string) (Workspace, error)
	AddUserToWorkspace(workspaceId string, userIdList []string) error
	CreateNewChannel(channel Channel, userId string) (Channel, error)
	AddUserToChannel(channelId string, userId []string) error
	GetDefaultChannelByWorkspaceId(workspaceId string) (Channel, error)
}

type service struct {
	r   Repository
	log *zap.Logger
}

func NewService(log *zap.Logger, r Repository) *service {
	return &service{
		log: log,
		r:   r,
	}
}

func (s *service) CreateWorkspace(w Workspace, userId string) (Workspace, error) {
	w.Id = uuid.New().String()
	w.CreatorId = userId
	c := Channel{
		ChannelId:   uuid.New().String(),
		ChannelName: "General",
		WorkspaceId: w.Id,
	}
	w, err := s.r.CreateNewWorkspace(w, c, userId)
	if err != nil {
		return Workspace{}, err
	}
	return w, nil
}

func (s *service) AddUserToWorkspace(workspaceId string, userIdList []string) error {
	err := s.r.AddUserToWorkspace(workspaceId, userIdList)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateChannel(workspaceId, channelName, userId string) (Channel, error) {
	c := NewChannel(uuid.New().String(), channelName, workspaceId)
	channel, err := s.r.CreateNewChannel(*c, userId)
	if err != nil {
		return channel, err
	}
	return channel, nil
}

func (s *service) AddUserToChannel(channelId string, userIdList []string) error {
	return s.r.AddUserToChannel(channelId, userIdList)
}
