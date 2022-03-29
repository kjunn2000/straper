package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"go.uber.org/zap"
)

func (q *Queries) CreateMessage(ctx context.Context, message *chatting.Message) error {
	sql, arg, err := sq.Insert("message").
		Columns("message_id", "type", "channel_id", "creator_id", "content", "file_name", "file_type", "created_date").
		Values(message.MessageId, message.Type,
			message.ChannelId, message.CreatorId,
			message.Content, message.FileName,
			message.FileType, message.CreatedDate).
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

func (q *Queries) GetAllChannelMessages(ctx context.Context, channelId string) ([]chatting.Message, error) {
	msgs := make([]chatting.Message, 0)
	sql, arg, err := sq.Select("message_id", "type", "channel_id", "creator_id", "content", "file_name", "file_type", "created_date").
		From("message").Where(sq.Eq{"channel_id": channelId}).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return []chatting.Message{}, err
	}
	err = q.db.Select(&msgs, sql, arg...)
	if err != nil {
		return []chatting.Message{}, err
	}
	return msgs, nil
}

func (q *Queries) GetAllChannelMessagesByWorkspaceId(ctx context.Context, workspaceId string) ([]chatting.Message, error) {
	msgs := make([]chatting.Message, 0)
	sql, arg, err := sq.Select("message_id", "type", "message.channel_id", "message.creator_id", "content", "file_name", "file_type", "message.created_date").
		From("message").Join("channel c on message.channel_id = c.channel_id").
		Where(sq.Eq{"c.workspace_id": workspaceId}).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return []chatting.Message{}, err
	}
	err = q.db.Select(&msgs, sql, arg...)
	if err != nil {
		return []chatting.Message{}, err
	}
	return msgs, nil
}

func (q *Queries) GetChannelMessages(ctx context.Context, channelId string, param chatting.PaginationMessagesParam) ([]chatting.Message, error) {
	msgs := make([]chatting.Message, 0)
	sb := sq.Select("message_id", "type", "channel_id", "creator_id", "content", "file_name", "file_type", "created_date").
		From("message").Where(sq.Eq{"channel_id": channelId}).OrderBy("created_date desc")
	if param.Cursor == "" {
		sb = sb.OrderBy("created_date desc", "message_id")
	} else {
		sb = sb.Where(sq.Or{
			sq.And{
				sq.Eq{"created_date": param.CreatedTime},
				sq.Gt{"message_id": param.Id}},
			sq.Lt{"created_date": param.CreatedTime}}).
			OrderBy("created_date desc", "message_id")
	}
	sql, arg, err := sb.Limit(uint64(25)).ToSql()
	if err != nil {
		q.log.Warn("Failed to create select sql.")
		return []chatting.Message{}, err
	}
	err = q.db.Select(&msgs, sql, arg...)
	if err != nil {
		return []chatting.Message{}, err
	}
	return msgs, nil
}

func (q *Queries) EditMessage(ctx context.Context, params chatting.EditChatMessageParams) error {
	sql, args, err := sq.Update("message").
		Set("content", params.Content).
		Where(sq.Eq{"message_id": params.MessageId}).
		ToSql()
	if err != nil {
		q.log.Info("Failed to create update message sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to update message.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) DeleteMessage(ctx context.Context, messageId string) error {
	sql, args, err := sq.Delete("message").
		Where(sq.Eq{"message_id": messageId}).
		ToSql()
	if err != nil {
		q.log.Info("Unable to create delete sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to delete chat message.", zap.Error(err))
		return err
	}
	return nil
}
