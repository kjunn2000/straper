package adding

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	CreateWorkspace(ctx context.Context, newWorkspaceName, userId string) (Workspace, error)
	AddUserToWorkspace(ctx context.Context, workspaceId string, userIdList []string) error
	CreateChannel(ctx context.Context, workspaceId, channelName, userId string) (Channel, error)
	AddUserToChannel(ctx context.Context, channelId string, userIdList []string) error
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

func (s *service) SeedWorkspaces() {
	for i := 1; i <= 25; i++ {
		s.CreateWorkspace(context.Background(), "workspace"+strconv.Itoa(i), "U000011")
	}
}

func (s *service) CreateWorkspace(ctx context.Context, newWorkspaceName, userId string) (Workspace, error) {
	id, _ := uuid.NewRandom()
	w := Workspace{
		Id:          id.String(),
		Name:        newWorkspaceName,
		CreatorId:   userId,
		CreatedDate: time.Now(),
	}
	c := NewChannel(uuid.New().String(), "General", w.Id, userId, true, time.Now())
	w.ChannelList = []Channel{c}
	w, err := s.r.CreateNewWorkspace(ctx, w, c, userId)
	if err != nil {
		return Workspace{}, err
	}
	return w, nil
}

func (s *service) AddUserToWorkspace(ctx context.Context, workspaceId string, userIdList []string) error {
	err := s.r.AddNewUserToWorkspace(ctx, workspaceId, userIdList)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateChannel(ctx context.Context, workspaceId, channelName, userId string) (Channel, error) {
	c := NewChannel(uuid.New().String(), channelName, workspaceId, userId, false, time.Now())
	channel, err := s.r.CreateNewChannel(ctx, c, userId)
	if err != nil {
		return channel, err
	}
	return channel, nil
}

func (s *service) AddUserToChannel(ctx context.Context, channelId string, userIdList []string) error {
	err := s.r.AddUserToChannel(ctx, channelId, userIdList)
	if err != nil {
		if strings.HasPrefix(err.Error(), "Error 1062") {
			return errors.New("channel.user.record.exist")
		}
		return errors.New("invalid.channel.id")
	}
	return nil
}
