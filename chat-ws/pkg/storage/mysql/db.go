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

type Queries struct {
	db  DBTX
	log *zap.Logger
}

func NewQueries(db DBTX, log *zap.Logger) *Queries {
	return &Queries{
		db:  db,
		log: log,
	}
}
