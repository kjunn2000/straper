package mysql

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type WorkspaceStore struct {
	db  *sqlx.DB
	Log *zap.Logger
	*WorkspaceQuery
}

func NewWorkspaceStore(db *sqlx.DB, log *zap.Logger) *WorkspaceStore {
	return &WorkspaceStore{
		db:             db,
		Log:            log,
		WorkspaceQuery: NewWorkspaceQuery(db, log),
	}
}

func (ws *WorkspaceStore) execTx(fn func(*WorkspaceQuery) error) error {
	tx, err := ws.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}
	err = fn(NewWorkspaceQuery(tx, ws.Log))
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
