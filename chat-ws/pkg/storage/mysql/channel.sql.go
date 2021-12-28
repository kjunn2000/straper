package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/editing"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/listing"
	"go.uber.org/zap"
)

func (q *Queries) CreateChannel(ctx context.Context, channel adding.Channel) error {
	sql, args, err := sq.Insert("channel").Columns("channel_id", "channel_name", "workspace_id", "creator_id", "created_date").
		Values(channel.ChannelId, channel.ChannelName, channel.WorkspaceId, channel.CreatorId, channel.CreatedDate).ToSql()
	if err != nil {
		q.log.Info("Unable to create insert channel sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to create new channel.", zap.String("workspace_id", channel.WorkspaceId))
		return err
	}
	return nil
}

func (q *Queries) AddUserToChannel(ctx context.Context, channelId string, userIdList []string) error {
	sqlBuilder := sq.Insert("channel_user").Columns("channel_id", "user_id")
	for _, id := range userIdList {
		sqlBuilder = sqlBuilder.Values(channelId, id)
	}
	sql, args, err := sqlBuilder.ToSql()
	if err != nil {
		q.log.Info("Unable to create add user to channel sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to add user to channel.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) GetChannelByChannelId(ctx context.Context, channelId string) (listing.Channel, error) {
	sql, args, err := sq.Select("channel_id, channel_name, workspace_id, creator_id,created_date").From("channel").
		Where(sq.Eq{"channel_id": channelId}).ToSql()
	if err != nil {
		q.log.Info("Unable to create select channel sql.", zap.Error(err))
		return listing.Channel{}, err
	}
	var channel listing.Channel
	err = q.db.Get(&channel, sql, args...)
	if err != nil {
		q.log.Info("Unable to create select channel sql.", zap.Error(err))
		return listing.Channel{}, err
	}
	return channel, nil
}

func (q *Queries) GetChannelsByUserId(ctx context.Context, userId string) ([]listing.Channel, error) {
	sql, args, err := sq.Select("channel.channel_id, channel_name, workspace_id, creator_id,created_date").
		From("channel").
		InnerJoin("channel_user as cu on channel.channel_id = cu.channel_id").
		Where(sq.Eq{"cu.user_id": userId}).
		OrderBy("created_date").
		ToSql()

	if err != nil {
		q.log.Info("Unable to create select channel sql.", zap.Error(err))
		return nil, err
	}

	var channels []listing.Channel
	err = q.db.Select(&channels, sql, args...)
	if err != nil {
		q.log.Info("Failed to select channel by workspace id.", zap.Error(err))
		return nil, err
	}
	return channels, nil
}

func (q *Queries) GetUserListByChannelId(ctx context.Context, channelId string) ([]chatting.User, error) {
	sql, args, err := sq.Select("user_id").From("channel_user").Where(sq.Eq{"channel_id": channelId}).ToSql()
	if err != nil {
		q.log.Info("Unable to create select client list sql.", zap.Error(err))
		return nil, err
	}
	var clientList []chatting.User
	err = q.db.Select(&clientList, sql, args...)
	if err != nil {
		q.log.Info("Failed to select client list.", zap.Error(err))
		return nil, err
	}
	return clientList, err
}

func (q *Queries) GetDefaultChannel(ctx context.Context, workspaceId string) (listing.Channel, error) {
	sql, args, err := sq.Select("channel_id, channel_name, workspace_id, creator_id, created_date").From("channel").
		Where(sq.Eq{"workspace_id": workspaceId}).
		Where(sq.Eq{"channel_name": "General"}).
		ToSql()
	if err != nil {
		q.log.Info("Unable to create select default channel sql.", zap.Error(err))
		return listing.Channel{}, err
	}
	var channel listing.Channel
	err = q.db.Get(&channel, sql, args...)
	if err != nil {
		q.log.Info("Failed to select default channel - General from db.", zap.Error(err))
		return listing.Channel{}, err
	}
	return channel, nil
}

func (q *Queries) GetDefaultChannelByWorkspaceId(ctx context.Context, workspaceId string) (adding.Channel, error) {
	sql, args, err := sq.Select("channel_id, channel_name, workspace_id, creator_id").From("channel").
		Where(sq.Eq{"workspace_id": workspaceId}).
		Where(sq.Eq{"channel_name": "General"}).
		ToSql()
	if err != nil {
		q.log.Info("Unable to create select default channel sql.", zap.Error(err))
		return adding.Channel{}, err
	}
	var channel adding.Channel
	err = q.db.Get(&channel, sql, args...)
	if err != nil {
		q.log.Info("Failed to select default channel - General from db.", zap.Error(err))
		return adding.Channel{}, err
	}
	return channel, nil
}

func (q *Queries) UpdateChannel(ctx context.Context, channel editing.Channel) error {
	sql, args, err := sq.Update("channel").
		Set("channel_name", channel.ChannelName).
		Where(sq.Eq{"channel_id": channel.ChannelId}).ToSql()
	if err != nil {
		q.log.Info("Failed to create update channel sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to update channel.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) DeleteChannel(ctx context.Context, channelId string) error {
	sql, args, err := sq.Delete("channel").Where(sq.Eq{"channel_id": channelId}).ToSql()
	if err != nil {
		q.log.Info("Unable to create delete channel sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to delete channel.", zap.Error(err))
		return err
	}
	return nil
}

func (q *Queries) RemoveUserFromChannel(ctx context.Context, channelId string, userId string) error {
	sql, args, err := sq.Delete("channel_user").
		Where(sq.Eq{"user_id": userId}).
		Where(sq.Eq{"channel_id": channelId}).
		ToSql()
	if err != nil {
		q.log.Info("Unable to create remove user from channels sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to remove user from channels.", zap.Error(err))
		return err
	}
	return nil
}
