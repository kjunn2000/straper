package mysql

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/listing"
	"go.uber.org/zap"
)

func (q *Queries) CreateChannel(channel *adding.Channel) error {
	sql, args, err := sq.Insert("channel").Columns("channel_id", "channel_name", "workspace_id").
		Values(channel.ChannelId, channel.ChannelName, channel.WorkspaceId).ToSql()
	if err != nil {
		q.log.Info("Unable to create insert channel sql.", zap.Error(err))
		return err
	}
	_, err = q.db.Exec(sql, args...)
	if err != nil {
		q.log.Info("Failed to create new channel.", zap.String("workspace_id", channel.WorkspaceId))
		return err
	}
	q.log.Info("Successfully create a new channel.", zap.String("channel_name", channel.ChannelName))
	return nil
}

func (q *Queries) AddUserToChannel(channelId string, userIdList []string) error {
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
	q.log.Info("Successfully add user to channel.", zap.String("channel_id", channelId))
	return nil
}

func (q *Queries) GetAllChannelByUserAndWorkspaceId(userId, workspaceId string) ([]listing.Channel, error) {
	sql, args, err := sq.Select("channel.channel_id, channel_name").
		From("channel").
		InnerJoin("channel_user on channel.channel_id = channel_user.channel_id").
		Where(sq.Eq{"workspace_id": workspaceId}).
		Where(sq.Eq{"user_id": userId}).
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

func (q *Queries) GetClientListByChannelId(channelId string) ([]chatting.Client, error) {
	sql, args, err := sq.Select("user_id").From("channel_user").Where(sq.Eq{"channel_id": channelId}).ToSql()
	if err != nil {
		q.log.Info("Unable to create select client list sql.", zap.Error(err))
		return nil, err
	}
	var clientList []chatting.Client
	err = q.db.Select(&clientList, sql, args...)
	if err != nil {
		q.log.Info("Failed to select client list.", zap.Error(err))
		return nil, err
	}
	return clientList, err
}

func (q *Queries) GetAllChannelByWorkspaceId(workspaceId string) ([]listing.Channel, error) {
	sql, args, err := sq.Select("channel.channel_id, channel_name").
		From("channel").
		Where(sq.Eq{"workspace_id": workspaceId}).
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

func (q *Queries) GetDefaultChannelByWorkspaceId(workspaceId string) (adding.Channel, error) {
	sql, args, err := sq.Select("channel_id").From("channel").
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

func (q *Queries) DeleteChannel(channelId string) error {
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
	q.log.Info("Successfully to delete 1 channel.", zap.String("channel_id", channelId))
	return nil
}

func (q *Queries) RemoveUserFromChannelList(channelIdList []string, userId string) error {
	sql, args, err := sq.Delete("channel_user").
		Where(sq.Eq{"user_id": userId}).
		Where(sq.Eq{"channel_id": channelIdList}).
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
	q.log.Info("Successfully remove 1 user from channels.", zap.String("user_id", userId))
	return nil
}
