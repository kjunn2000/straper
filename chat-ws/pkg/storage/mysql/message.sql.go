package mysql

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"go.uber.org/zap"
)

func (q *Queries) CreateMessage(ctx context.Context, message *chatting.Message) error {
	sql, arg, err := sq.Insert("message").
		Columns("message_id", "type", "channel_id", "creator_name", "content", "file_name", "file_type", "created_date").
		Values(message.MessageId, message.Type,
			message.ChannelId, message.CreatorName,
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
	sql, arg, err := sq.Select("message_id", "type", "channel_id", "creator_name", "content", "file_name", "file_type", "created_date").
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
	sql, arg, err := sq.Select("message_id", "type", "message.channel_id", "creator_name", "content", "file_name", "file_type", "message.created_date").
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

func (q *Queries) GetChannelMessages(ctx context.Context, channelId string, limit, offset uint64) ([]chatting.Message, error) {
	msgs := make([]chatting.Message, 0)
	sql, arg, err := sq.Select("message_id", "type", "channel_id", "creator_name", "content", "file_name", "file_type", "created_date").
		From("message").Where(sq.Eq{"channel_id": channelId}).OrderBy("created_date desc").Limit(limit).Offset(offset).ToSql()
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

func (q *Queries) UpdateChannelAccessTime(ctx context.Context, channelId string, userId string) error {
	sql, args, err := sq.Update("channel_user").
		Set("last_accessed", time.Now()).
		Where(sq.And{
			sq.Eq{"channel_id": channelId},
			sq.Eq{"user_id": userId},
		}).ToSql()
	if err != nil {
		q.log.Info("Failed to create update last accessed sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to update last accessed.", zap.Error(err))
		return err
	}
	return nil
}
