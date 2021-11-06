package mysql

import (
	"context"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"go.uber.org/zap"
)

func (s *SQLStore) CreateNewWorkspace(ctx context.Context, w adding.Workspace, c adding.Channel, userId string) (adding.Workspace, error) {
	err := s.execTx(func(q *Queries) error {
		err := q.CreateWorkspace(ctx, w)
		if err != nil {
			return err
		}
		userIdList := []string{userId}
		err = q.AddUserToWorkspace(ctx, w.Id, userIdList)
		if err != nil {
			return err
		}
		err = q.CreateChannel(ctx, c)
		if err != nil {
			return err
		}
		err = q.AddUserToChannel(ctx, c.ChannelId, userIdList)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		s.log.Info("Failed to create a new workspace.", zap.Error(err))
		return adding.Workspace{}, err
	}
	return w, nil
}

func (s *SQLStore) AddNewUserToWorkspace(ctx context.Context, workspaceId string, userIdList []string) error {
	err := s.execTx(func(q *Queries) error {
		err := q.AddUserToWorkspace(ctx, workspaceId, userIdList)
		if err != nil {
			return err
		}
		c, err := q.GetDefaultChannelByWorkspaceId(ctx, workspaceId)
		if err != nil {
			return err
		}
		err = s.AddUserToChannel(ctx, c.ChannelId, userIdList)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		s.log.Info("Failed to add user to workspace.", zap.Error(err))
		return err
	}
	return nil
}