package mysql

import (
	"database/sql"

	"go.uber.org/zap"
)

type DBTX interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

type WorkspaceQuery struct {
	db  DBTX
	Log *zap.Logger
}

func NewWorkspaceQuery(db DBTX, log *zap.Logger) *WorkspaceQuery {
	return &WorkspaceQuery{
		db:  db,
		Log: log,
	}
}
