package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"go.uber.org/zap"
)

func (q *Queries) CreateCardComment(ctx context.Context, comment *chatting.CardComment) error {
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

func (q *Queries) GetCardComments(ctx context.Context, cardId string) ([]chatting.CardComment, error) {
	var cardComments []chatting.CardComment
	sql, arg, err := sq.Select("comment_id", "type", "card_id", "creator_id", "content", "file_name",
		"file_type", "created_date").
		From("card_comment").Where(sq.Eq{"cardId": cardId}).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return []chatting.CardComment{}, err
	}
	err = q.db.Select(&cardComments, sql, arg...)
	if err != nil {
		return []chatting.CardComment{}, err
	}
	return cardComments, nil
}
