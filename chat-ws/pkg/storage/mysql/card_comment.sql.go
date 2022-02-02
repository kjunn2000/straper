package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/board"
	"go.uber.org/zap"
)

func (q *Queries) CreateCardComment(ctx context.Context, comment *board.CardComment) error {
	sql, arg, err := sq.Insert("card_comment").
		Columns("comment_id", "type", "card_id", "creator_id", "content", "file_name",
			"file_type", "created_date").
		Values(comment.CommentId, comment.Type, comment.CardId, comment.CreatorId,
			comment.Content, comment.FileName, comment.FileType, comment.CreatedDate).
		ToSql()
	if err != nil {
		q.log.Warn("Failed to create card comment query.")
		return err
	}
	_, err = q.db.Exec(sql, arg...)
	if err != nil {
		q.log.Info("Failed to create card comment to db.", zap.String("error", err.Error()))
		return err
	}
	return nil
}

func (q *Queries) GetCardComments(ctx context.Context, cardId string, limit, offset uint64) ([]board.CardComment, error) {
	var cardComments []board.CardComment
	sql, arg, err := sq.Select("comment_id", "type", "card_id", "creator_id", "content", "file_name",
		"file_type", "created_date").
		From("card_comment").Where(sq.Eq{"card_id": cardId}).OrderBy("created_date desc").Limit(limit).Offset(offset).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return []board.CardComment{}, err
	}
	err = q.db.Select(&cardComments, sql, arg...)
	if err != nil {
		return []board.CardComment{}, err
	}
	return cardComments, nil
}

func (q *Queries) DeleteCardComment(ctx context.Context, commentId string) error {
	sql, args, err := sq.Delete("card_comment").
		Where(sq.Eq{"comment_id": commentId}).
		ToSql()
	if err != nil {
		q.log.Info("Unable to create delete sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to delete comment from card.", zap.Error(err))
		return err
	}
	return nil
}
