package mysql

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"go.uber.org/zap"
)

type Store interface {
	CreateNewWorkspace(ctx context.Context, w adding.Workspace, c adding.Channel, userId string) (adding.Workspace, error)
	CreateNewChannel(ctx context.Context, channel adding.Channel, userId string) (adding.Channel, error)
	AddNewUserToWorkspace(ctx context.Context, workspaceId string, userIdList []string) error
	Querier
}

type SQLStore struct {
	db *sqlx.DB
	*Queries
}

func NewStore(db *sqlx.DB, log *zap.Logger) Store {
	return &SQLStore{
		db:      db,
		Queries: NewQueries(db, log),
	}
}

func (s *SQLStore) execTx(fn func(*Queries) error) error {
	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}
	err = fn(NewQueries(tx, s.log))
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

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

func (s *SQLStore) CreateNewChannel(ctx context.Context, channel adding.Channel, userId string) (adding.Channel, error) {
	err := s.execTx(func(q *Queries) error {
		err := q.CreateChannel(ctx, channel)
		if err != nil {
			return err
		}
		err = q.AddUserToChannel(ctx, channel.ChannelId, []string{userId})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		s.log.Info("Failed to create a new channel.", zap.Error(err))
		return adding.Channel{}, err
	}
	return channel, nil
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
