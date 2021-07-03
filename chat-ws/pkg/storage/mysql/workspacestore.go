package storage

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain"
	"github.com/kjunn2000/straper/chat-ws/pkg/storage"
	"go.uber.org/zap"
)

type WorkspaceStore struct {
	Db  *sqlx.DB
	Log *zap.Logger
}

func NewWorkspaceStore(db *sqlx.DB, log *zap.Logger) *WorkspaceStore {
	return &WorkspaceStore{
		Db:  db,
		Log: log,
	}
}

func (ws *WorkspaceStore) CreateWorkspace(w domain.Workspace) error {
	id, err := storage.GetId(w.Name)
	if err != nil {
		ws.Log.Warn("Unable to generate id")
		return err
	}
	w.Id = id
	sql, args, err := sq.Insert("workspace").Columns("workspace_id", "workspace_name").Values(w.Id, w.Name).ToSql()
	if err != nil {
		ws.Log.Warn("Unable to create insert workspace query.")
		return err
	}
	_, err = ws.Db.Exec(sql, args...)
	if err != nil {
		ws.Log.Warn("Unable to execute insert workspace query.", zap.Error(err))
		return err
	}
	ws.Log.Info("Successfully create new workspace")
	return nil
}

func (ws *WorkspaceStore) EditWorkspace(w domain.Workspace) error {
	_, err := ws.GetWorkspace(w.Id)
	if err != nil {
		return err
	}
	sql, args, err := sq.Update("workspace").Set("workspace_name", w.Name).Where(sq.Eq{"workspace_id": w.Id}).ToSql()
	if err != nil {
		ws.Log.Warn("Unable to create edit workspace query.")
		return err
	}
	_, err = ws.Db.Exec(sql, args...)
	if err != nil {
		ws.Log.Warn("Unable to edit workspace.", zap.Error(err))
		return err
	}
	ws.Log.Info("Successfully edit workspace", zap.String("id", w.Id))
	return nil
}

func (ws *WorkspaceStore) DeleteWorkspace(id string) error {
	_, err := ws.GetWorkspace(id)
	if err == sql.ErrNoRows {
		ws.Log.Warn("Workspace id not found.")
		return err
	}
	sql, args, err := sq.Delete("workspace").Where(sq.Eq{"workspace_id": id}).ToSql()
	if err != nil {
		ws.Log.Warn("Unable to create delete workspace query.")
		return err
	}
	_, err = ws.Db.Exec(sql, args...)
	if err != nil {
		ws.Log.Warn("Unable to delete workspace", zap.Error(err))
		return err
	}
	ws.Log.Info("Successfully delete workspace", zap.String("id", id))
	return nil
}

func (ws *WorkspaceStore) GetWorkspaces() ([]domain.Workspace, error) {
	sql, _, err := sq.Select("*").From("workspace").ToSql()
	if err != nil {
		ws.Log.Warn("Unable to create select workspace list query")
	}
	var Workspaces []domain.Workspace
	err = ws.Db.Select(&Workspaces, sql, nil)
	if err != nil {
		ws.Log.Warn("Unable to select workspace list from db")
		return nil, err
	}
	return Workspaces, nil
}

func (ws *WorkspaceStore) GetWorkspace(id string) (domain.Workspace, error) {
	s, args, err := sq.Select("*").From("workspace").Where(sq.Eq{"workspace_id": id}).Limit(1).ToSql()
	if err != nil {
		ws.Log.Warn("Unable to create select workspace list query")
	}
	var workspace domain.Workspace
	err = ws.Db.Get(&workspace, s, args...)
	if err == sql.ErrNoRows {
		ws.Log.Info("Workspace id not found", zap.String("id", id))
		return domain.Workspace{}, err
	} else if err != nil {
		ws.Log.Warn("Unable to select workspace from db", zap.Error(err))
		return domain.Workspace{}, err
	}
	return workspace, nil
}
