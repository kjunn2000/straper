package mysql

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/admin"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/bug"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"go.uber.org/zap"
)

type Store interface {
	CreateUser(ctx context.Context, params account.CreateUserParam) error
	CreateNewWorkspace(ctx context.Context, w adding.Workspace, c adding.Channel, userId string) (adding.Workspace, error)
	CreateNewChannel(ctx context.Context, channel adding.Channel, userId string) (adding.Channel, error)
	AddNewUserToWorkspace(ctx context.Context, workspaceId string, userIdList []string) error
	ValidateAccountEmail(ctx context.Context, userId string, tokenId string) error
	DeleteIssueAndAttachments(ctx context.Context, issueId string, attachments []bug.Attachment) error
	UpdateUserByAdmin(ctx context.Context, param admin.UpdateUserParam) error
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
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
