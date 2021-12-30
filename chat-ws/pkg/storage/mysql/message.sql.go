package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"go.uber.org/zap"
)

func (q *Queries) CreateMessage(ctx context.Context, message *chatting.Message) error {
	sql, arg, err := sq.Insert("message").
		Columns("message_id", "type", "channel_id", "creator_name", "content", "created_date").
		Values(message.MessageId, message.Type, message.ChannelId, message.CreatorName, message.Content, message.CreatedDate).
		ToSql()
	if err != nil {
		q.log.Warn("Failed to create message query.")
		return err
	}
	_, err = q.db.Exec(sql, arg...)
	if err != nil {
		q.log.Info("Failed to create message to db.", zap.String("error", err.Error()))
		return err
	}
	return nil
}
