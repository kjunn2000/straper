package deleting

import "go.uber.org/zap"

type Service interface {
	DeleteWorkspace(workspaceId string) error
	LeaveWorkspace(workspaceId, userId string) error
	DeleteChannel(channelId string) error
	LeaveChannel(channelId, userId string) error
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
	DeleteWorkspace(workspaceId string) error
	RemoveUserFromWorkspace(workspaceId, userId string) error
	DeleteChannel(channelId string) error
	RemoveUserFromChannelList(channelId []string, userId string) error
}

type Channel struct {
	ChannelId string `json:"channel_id" db:"channel_id"`
}

func (s *service) DeleteWorkspace(workspaceId string) error {
	err := s.r.DeleteWorkspace(workspaceId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) LeaveWorkspace(workspaceId, userId string) error {
	err := s.r.RemoveUserFromWorkspace(workspaceId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteChannel(channelId string) error {
	err := s.r.DeleteChannel(channelId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) LeaveChannel(channelId, userId string) error {
	err := s.r.RemoveUserFromChannelList([]string{channelId}, userId)
	if err != nil {
		return err
	}
	return nil
}
