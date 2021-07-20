package mysql

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
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

func (ws *WorkspaceStore) CreateWorkspace(w adding.Workspace) error {
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

// func (ws *WorkspaceStore) EditWorkspace(w adding.Workspace) error {
// 	sql, args, err := sq.Update("workspace").Set("workspace_name", w.Name).Where(sq.Eq{"workspace_id": w.Id}).ToSql()
// 	if err != nil {
// 		ws.Log.Warn("Unable to create edit workspace query.")
// 		return err
// 	}
// 	res, err := ws.Db.Exec(sql, args...)
// 	if err != nil {
// 		ws.Log.Warn("Unable to edit workspace.", zap.Error(err))
// 		return err
// 	}
// 	rowAffected, err := res.RowsAffected()
// 	if rowAffected == 0 || err != nil {
// 		ws.Log.Info("Workspace Id not found.", zap.Error(err))
// 		return errors.New("Workspace Id not found.")
// 	}
// 	ws.Log.Info("Successfully edit workspace", zap.String("id", w.Id))
// 	return nil
// }

// func (ws *WorkspaceStore) DeleteWorkspace(id string) error {
// 	sql, args, err := sq.Delete("workspace").Where(sq.Eq{"workspace_id": id}).ToSql()
// 	if err != nil {
// 		ws.Log.Warn("Unable to create delete workspace query.")
// 		return err
// 	}
// 	res, err := ws.Db.Exec(sql, args...)
// 	if err != nil {
// 		ws.Log.Warn("Unable to delete workspace.", zap.Error(err))
// 		return err
// 	}
// 	rowAffected, err := res.RowsAffected()
// 	if rowAffected == 0 || err != nil {
// 		ws.Log.Warn("Workspace Id not found.", zap.Error(err))
// 		return err
// 	}
// 	ws.Log.Info("Successfully delete workspace.", zap.String("id", id))
// 	return nil
// }

// func (ws *WorkspaceStore) GetWorkspacesByUserId(userId string) ([]adding.Workspace, error) {
// 	sql, args, err := sq.Select("workspace_user.workspace_id", "workspace_name").From("workspace_user").
// 		Where(sq.Eq{"user_id": userId}).
// 		InnerJoin("workspace as w on workspace_user.workspace_id = w.workspace_id").
// 		ToSql()
// 	if err != nil {
// 		ws.Log.Warn("Unable to create select workspace list query")
// 	}
// 	var Workspaces []adding.Workspace
// 	err = ws.Db.Select(&Workspaces, sql, args...)
// 	if err != nil {
// 		ws.Log.Warn("Unable to select workspace list from db")
// 		return nil, err
// 	}
// 	return Workspaces, nil
// }

func (ws *WorkspaceStore) AddUserToWorkspace(workspaceId string, userIdList []string) error {
	sqlBuilder := sq.Insert("workspace_user").Columns("workspace_id", "user_id")
	for _, id := range userIdList {
		sqlBuilder = sqlBuilder.Values(workspaceId, id)
	}
	sql, args, err := sqlBuilder.ToSql()
	if err != nil {
		ws.Log.Warn("Fail to create add user to workspace query.", zap.Error(err))
		return err
	}
	res, err := ws.Db.Exec(sql, args...)
	if err != nil {
		ws.Log.Info("Unable to execute add user to workspace query.", zap.Error(err))
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		ws.Log.Info("Unabel to extract affected rows.", zap.Error(err))
		return err
	}
	ws.Log.Info("Sucessful added new user to workspace.",
		zap.String("WorkspaceId", workspaceId),
		zap.Int64("RowsAffected", rows))
	return nil
}

func (ws *WorkspaceStore) GetClientIdListByWorkspaceId(workspaceId string) 
