package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/board"
	"go.uber.org/zap"
)

func (q *Queries) CreateBoard(ctx context.Context, board board.TaskBoard) error {
	sql, args, err := sq.Insert("task_board").Columns("board_id", "board_name", "workspace_id").
		Values(board.BoardId, board.BoardName, board.WorkspaceId).ToSql()
	if err != nil {
		q.log.Info("Unable to create insert task board sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to create new task board.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) GetTaskBoardByWorkspaceId(ctx context.Context, workspaceId string) (board.TaskBoard, error) {
	sql, args, err := sq.Select("board_id, board_name, workspace_id").From("task_board").
		Where(sq.Eq{"workspace_id": workspaceId}).ToSql()
	if err != nil {
		q.log.Info("Unable to create select task board sql.", zap.Error(err))
		return board.TaskBoard{}, err
	}
	var taskBoard board.TaskBoard
	err = q.db.Get(&taskBoard, sql, args...)
	if err != nil {
		q.log.Info("Unable to create select task board sql.", zap.Error(err))
		return board.TaskBoard{}, err
	}
	return taskBoard, nil
}

func (q *Queries) UpdateTaskBoard(ctx context.Context, board board.TaskBoard) error {
	sql, args, err := sq.Update("task_board").
		Set("board_name", board.BoardName).
		Where(sq.Eq{"board_id": board.BoardId}).ToSql()
	if err != nil {
		q.log.Info("Failed to create update task board sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to update task board.", zap.Error(err))
		return err
	}
	return nil
}
