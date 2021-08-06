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
	DeleteWorkspace(workspaceId string) error
	RemoveUserFromWorkspace(workspaceId, userId string) error
}

type ChannelRepository interface {
	DeleteChannel(channelId string) error
	RemoveUserFromChannelList(channelId []string, userId string) error
}

type Channel struct {
	ChannelId string `json:"channel_id" db:"channel_id"`
}

func (s *service) DeleteWorkspace(workspaceId string) error {
	err := s.wr.DeleteWorkspace(workspaceId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) LeaveWorkspace(workspaceId, userId string) error {
	err := s.wr.RemoveUserFromWorkspace(workspaceId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteChannel(channelId string) error {
	err := s.cr.DeleteChannel(channelId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) LeaveChannel(channelId, userId string) error {
	err := s.cr.RemoveUserFromChannelList([]string{channelId}, userId)
	if err != nil {
		return err
	}
	return nil
}
