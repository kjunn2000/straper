package mysql

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/admin"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/editing"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/listing"
	"go.uber.org/zap"
)

func (q *Queries) CreateWorkspace(ctx context.Context, w adding.Workspace) error {
	sql, args, err := sq.Insert("workspace").Columns("workspace_id", "workspace_name", "creator_id", "created_date").
		Values(w.Id, w.Name, w.CreatorId, w.CreatedDate).ToSql()
	if err != nil {
		q.log.Warn("Unable to create insert workspace query.")
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Unable to execute insert workspace query.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) AddUserToWorkspace(ctx context.Context, workspaceId string, userIdList []string) error {
	sqlBuilder := sq.Insert("workspace_user").Columns("workspace_id", "user_id")
	for _, id := range userIdList {
		sqlBuilder = sqlBuilder.Values(workspaceId, id)
	}
	sql, args, err := sqlBuilder.ToSql()
	if err != nil {
		q.log.Warn("Fail to create add user to workspace query.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Unable to execute add user to workspace query.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) GetWorkspace(ctx context.Context, workspaceId string) (admin.Workspace, error) {
	workspace, err := q.GetWorkspaceByAdmin(ctx, workspaceId)
	if err != nil {
		return admin.Workspace{}, err
	}
	channelList, err := q.GetWorkspaceChannelsByAdmin(ctx, workspaceId)
	if err != nil {
		return admin.Workspace{}, err
	}
	userList, err := q.GetWorkspaceUsersByAdmin(ctx, workspaceId)
	if err != nil {
		return admin.Workspace{}, err
	}
	workspace.ChannelList = channelList
	workspace.UserList = userList
	return workspace, nil
}

func (q *Queries) GetWorkspaceByWorkspaceId(ctx context.Context, workspaceId string) (listing.Workspace, error) {
	sql, args, err := sq.Select("workspace_id", "workspace_name, creator_id, created_date").
		From("workspace").
		Where(sq.Eq{"workspace_id": workspaceId}).
		OrderBy("created_date").
		ToSql()

	if err != nil {
		q.log.Warn("Unable to create select workspace query.")
		return listing.Workspace{}, err
	}
	var workspace listing.Workspace
	err = q.db.Get(&workspace, sql, args...)
	if err != nil {
		q.log.Info("Unable to select workspace from db")
		return listing.Workspace{}, err
	}
	return workspace, nil
}

func (q *Queries) GetWorkspaceByAdmin(ctx context.Context, workspaceId string) (admin.Workspace, error) {
	sql, args, err := sq.Select("workspace_id", "workspace_name, creator_id, created_date").
		From("workspace").
		Where(sq.Eq{"workspace_id": workspaceId}).
		OrderBy("created_date").
		ToSql()

	if err != nil {
		q.log.Warn("Unable to create select workspace query.")
		return admin.Workspace{}, err
	}
	var workspace admin.Workspace
	err = q.db.Get(&workspace, sql, args...)
	if err != nil {
		q.log.Info("Unable to select workspace from db")
		return admin.Workspace{}, err
	}
	return workspace, nil
}

func (q *Queries) GetWorkspacesByUserId(ctx context.Context, userId string) ([]listing.Workspace, error) {
	sql, args, err := sq.Select("workspace_user.workspace_id", "workspace_name, creator_id").From("workspace_user").
		Where(sq.Eq{"user_id": userId}).
		InnerJoin("workspace as w on workspace_user.workspace_id = w.workspace_id").
		OrderBy("created_date").
		ToSql()
	if err != nil {
		q.log.Warn("Unable to create select workspace list query")
		return nil, err
	}
	var Workspaces []listing.Workspace
	err = q.db.Select(&Workspaces, sql, args...)
	if err != nil {
		q.log.Info("Unable to select workspace list from db")
		return nil, err
	}
	return Workspaces, nil
}

func (q *Queries) GetWorkspacesByCursor(ctx context.Context, param admin.PaginationWorkspacesParam) ([]admin.WorkspaceSummary, error) {
	var workspaces []admin.WorkspaceSummary
	sb := sq.Select("workspace_id", "workspace_name, creator_id, created_date").
		From("workspace").
		Where(sq.Or{
			sq.Like{"workspace_id": fmt.Sprintf("%%%s%%", param.SearchStr)},
			sq.Like{"workspace_name": fmt.Sprintf("%%%s%%", param.SearchStr)},
			sq.Like{"creator_id": fmt.Sprintf("%%%s%%", param.SearchStr)},
		})
	if param.Cursor == "" {
		sb = sb.OrderBy("created_date desc", "workspace_id")
	} else if param.Cursor != "" && param.IsNext {
		sb = sb.Where(sq.Or{
			sq.And{
				sq.Eq{"created_date": param.CreatedTime},
				sq.Gt{"workspace_id": param.Id}},
			sq.Lt{"created_date": param.CreatedTime}}).
			OrderBy("created_date desc", "workspace_id")
	} else {
		sb = sb.Where(sq.Or{
			sq.And{
				sq.Eq{"created_date": param.CreatedTime},
				sq.Lt{"workspace_id": param.Id}},
			sq.Gt{"created_date": param.CreatedTime}}).
			OrderBy("created_date", "workspace_id desc")
	}
	sql, arg, err := sb.Limit(uint64(param.Limit)).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return []admin.WorkspaceSummary{}, err
	}
	err = q.db.Select(&workspaces, sql, arg...)
	if err != nil {
		return []admin.WorkspaceSummary{}, err
	}
	if param.Cursor != "" && !param.IsNext {
		for i, j := 0, len(workspaces)-1; i < j; i, j = i+1, j-1 {
			workspaces[i], workspaces[j] = workspaces[j], workspaces[i]
		}
	}
	return workspaces, nil
}

func (q *Queries) GetWorkspacesCount(ctx context.Context, searchStr string) (int, error) {
	var count int
	sql, arg, err := sq.Select("COUNT(workspace_id)").
		From("workspace").
		Where(sq.Or{
			sq.Like{"workspace_id": fmt.Sprintf("%%%s%%", searchStr)},
			sq.Like{"workspace_name": fmt.Sprintf("%%%s%%", searchStr)},
			sq.Like{"creator_id": fmt.Sprintf("%%%s%%", searchStr)},
		}).
		ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return 0, err
	}
	err = q.db.Get(&count, sql, arg...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (q *Queries) GetWorkspaceUserCount(ctx context.Context, workspaceId string) (int, error) {
	sql, args, err := sq.Select("COUNT(w.workspace_id) as total_users").
		From("workspace w").
		InnerJoin("workspace_user wc on w.workspace_id = wc.workspace_id").
		Where(sq.Eq{"w.workspace_id": workspaceId}).
		ToSql()
	if err != nil {
		q.log.Warn("Unable to create count workspace users query")
		return 0, err
	}
	var totalUser int
	err = q.db.Get(&totalUser, sql, args...)
	if err != nil {
		q.log.Info("Unable to count workspace users from db")
		return 0, err
	}
	return totalUser, nil
}

func (q *Queries) GetWorkspaceChannelCount(ctx context.Context, workspaceId string) (int, error) {
	sql, args, err := sq.Select("COUNT(workspace_id) as total_channels").
		From("channel").
		Where(sq.Eq{"workspace_id": workspaceId}).
		ToSql()
	if err != nil {
		q.log.Warn("Unable to create count workspace channels query")
		return 0, err
	}
	var totalChannel int
	err = q.db.Get(&totalChannel, sql, args...)
	if err != nil {
		q.log.Info("Unable to count workspace channel from db")
		return 0, err
	}
	return totalChannel, nil
}

func (q *Queries) UpdateWorkspace(ctx context.Context, workspace editing.Workspace) error {
	sql, args, err := sq.Update("workspace").Set("workspace_name", workspace.Name).Where(sq.Eq{"workspace_id": workspace.Id}).ToSql()
	if err != nil {
		q.log.Warn("Failed to create update workspace query.")
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to update workspace.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) DeleteWorkspace(ctx context.Context, id string) error {
	sql, args, err := sq.Delete("workspace").Where(sq.Eq{"workspace_id": id}).ToSql()
	if err != nil {
		q.log.Warn("Unable to create delete workspace query.")
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Unable to delete workspace.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) RemoveUserFromWorkspace(ctx context.Context, workspaceId, userId string) error {
	sql := "DELETE workspace_user, channel_user FROM workspace_user " +
		"INNER JOIN channel ON workspace_user.workspace_id = channel.workspace_id " +
		"INNER JOIN channel_user ON channel.channel_id = channel_user.channel_id " +
		"WHERE workspace_user.workspace_id = ? " +
		"AND workspace_user.user_id = ? " +
		"AND channel_user.user_id = ? ;"
	_, err := q.db.Exec(sql, workspaceId, userId, userId)
	if err != nil {
		q.log.Info("Failed to remove user from workspace", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) GetUserListByWorkspaceId(ctx context.Context, workspaceId string) ([]websocket.UserData, error) {
	sql, args, err := sq.Select("user_id").From("workspace_user").
		Where(sq.Eq{"workspace_id": workspaceId}).ToSql()
	if err != nil {
		q.log.Info("Unable to create select client list sql.", zap.Error(err))
		return nil, err
	}
	var clientList []websocket.UserData
	err = q.db.Select(&clientList, sql, args...)
	if err != nil {
		q.log.Info("Failed to select client list.", zap.Error(err))
		return nil, err
	}
	return clientList, err
}
