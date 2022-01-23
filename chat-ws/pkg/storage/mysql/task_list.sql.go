package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/board"
	"go.uber.org/zap"
)

func (q *Queries) CreateTaskList(ctx context.Context, taskList board.TaskList) error {
	sql, args, err := sq.Insert("task_list").Columns("list_id", "list_name", "board_id", "order_index").
		Values(taskList.ListId, taskList.ListName, taskList.BoardId, taskList.OrderIndex).ToSql()
	if err != nil {
		q.log.Info("Unable to create insert task list sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to create new task list.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) GetTaskListsByBoardId(ctx context.Context, boardId string) ([]board.TaskList, error) {
	sql, args, err := sq.Select("list_id", "list_name", "board_id", "order_index").From("task_list").
		Where(sq.Eq{"board_id": boardId}).
		OrderBy("order_index").
		ToSql()
	if err != nil {
		q.log.Info("Unable to create select task list sql.", zap.Error(err))
		return []board.TaskList{}, err
	}
	var taskLists []board.TaskList
	err = q.db.Select(&taskLists, sql, args...)
	if err != nil {
		q.log.Info("Unable to create select task list sql.", zap.Error(err))
		return []board.TaskList{}, err
	}
	return taskLists, nil
}

func (q *Queries) UpdateTaskList(ctx context.Context, taskList board.UpdateListParams) error {
	sql, args, err := sq.Update("task_list").
		Set("list_name", taskList.ListName).
		Where(sq.Eq{"list_id": taskList.ListId}).ToSql()
	if err != nil {
		q.log.Info("Failed to create update task list sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to update task list.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) UpdateTaskListOrder(ctx context.Context, listId string, orderIndex int) error {
	sql, args, err := sq.Update("task_list").
		Set("order_index", orderIndex).
		Where(sq.Eq{"list_id": listId}).ToSql()
	if err != nil {
		q.log.Info("Failed to create update task list order sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to update task list order.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) DeleteTaskList(ctx context.Context, listId string) error {
	sql, args, err := sq.Delete("task_list").Where(sq.Eq{"list_id": listId}).ToSql()
	if err != nil {
		q.log.Info("Unable to create delete task list sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to delete task list.", zap.Error(err))
		return err
	}
	return nil
}
