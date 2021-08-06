package mysql

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/listing"
	"go.uber.org/zap"
)

func (ws *WorkspaceQuery) CreateWorkspace(w adding.Workspace) error {
	sql, args, err := sq.Insert("workspace").Columns("workspace_id", "workspace_name", "creator_id").
		Values(w.Id, w.Name, w.CreatorId).ToSql()
	if err != nil {
		ws.Log.Warn("Unable to create insert workspace query.")
		return err
	}
	_, err = ws.db.Exec(sql, args...)
	if err != nil {
		ws.Log.Warn("Unable to execute insert workspace query.", zap.Error(err))
		return err
	}
	ws.Log.Info("Successfully create new workspace")
	return nil
}

func (ws *WorkspaceQuery) DeleteWorkspace(id string) error {
	sql, args, err := sq.Delete("workspace").Where(sq.Eq{"workspace_id": id}).ToSql()
	if err != nil {
		ws.Log.Warn("Unable to create delete workspace query.")
		return err
	}
	res, err := ws.db.Exec(sql, args...)
	if err != nil {
		ws.Log.Warn("Unable to delete workspace.", zap.Error(err))
		return err
	}
	rowAffected, err := res.RowsAffected()
	if rowAffected == 0 || err != nil {
		ws.Log.Warn("Workspace Id not found.", zap.Error(err))
		return err
	}
	ws.Log.Info("Successfully delete workspace.", zap.String("id", id))
	return nil
}

func (ws *WorkspaceQuery) RemoveUserFromWorkspace(workspaceId, userId string) error {
	sql := "DELETE workspace_user, channel_user FROM workspace_user " +
		"INNER JOIN channel ON workspace_user.workspace_id = channel.workspace_id " +
		"INNER JOIN channel_user ON channel.channel_id = channel_user.channel_id " +
		"WHERE workspace_user.workspace_id = ? " +
		"AND workspace_user.user_id = ? " +
		"AND channel_user.user_id = ? ;"
	_, err := ws.db.Exec(sql, workspaceId, userId, userId)
	if err != nil {
		ws.Log.Info("Failed to remove user from workspace", zap.Error(err))
		return err
	}
	ws.Log.Info("Successfully remove 1 user from workspace.", zap.String("user_id", userId))
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

func (ws *WorkspaceQuery) GetWorkspaceByWorkspaceId(workspaceId string) (listing.Workspace, error) {
	sql, args, err := sq.Select("workspace_id", "workspace_name, creator_id").From("workspace").Where(sq.Eq{"workspace_id": workspaceId}).ToSql()
	if err != nil {
		ws.Log.Warn("Unable to create select workspace query.")
		return listing.Workspace{}, err
	}
	var workspace listing.Workspace
	err = ws.db.Get(&workspace, sql, args...)
	if err != nil {
		ws.Log.Warn("Unable to select workspace from db")
		return listing.Workspace{}, err
	}
	return workspace, nil
}

func (ws *WorkspaceQuery) GetWorkspacesByUserId(userId string) ([]listing.Workspace, error) {
	sql, args, err := sq.Select("workspace_user.workspace_id", "workspace_name, creator_id").From("workspace_user").
		Where(sq.Eq{"user_id": userId}).
		InnerJoin("workspace as w on workspace_user.workspace_id = w.workspace_id").
		ToSql()
	if err != nil {
		ws.Log.Warn("Unable to create select workspace list query")
		return nil, err
	}
	var Workspaces []listing.Workspace
	err = ws.db.Select(&Workspaces, sql, args...)
	if err != nil {
		ws.Log.Warn("Unable to select workspace list from db")
		return nil, err
	}
	return Workspaces, nil
}

func (ws *WorkspaceQuery) AddUserToWorkspace(workspaceId string, userIdList []string) error {
	sqlBuilder := sq.Insert("workspace_user").Columns("workspace_id", "user_id")
	for _, id := range userIdList {
		sqlBuilder = sqlBuilder.Values(workspaceId, id)
	}
	sql, args, err := sqlBuilder.ToSql()
	if err != nil {
		ws.Log.Warn("Fail to create add user to workspace query.", zap.Error(err))
		return err
	}
	res, err := ws.db.Exec(sql, args...)
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
