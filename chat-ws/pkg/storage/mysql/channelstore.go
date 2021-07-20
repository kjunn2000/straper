package mysql

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"go.uber.org/zap"
)

type ChannelStore struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewChannelStore(db *sqlx.DB, log *zap.Logger) *ChannelStore {
	return &ChannelStore{
		db:  db,
		log: log,
	}
}

func (cs *ChannelStore) CreateChannel(channel *adding.Channel) error {
	sql, args, err := sq.Insert("channel").Columns("channel_id", "channel_name", "workspace_id").
		Values(channel.ChannelId, channel.ChannelName, channel.WorkspaceId).ToSql()
	if err != nil {
		cs.log.Info("Unable to create insert channel sql.", zap.Error(err))
		return err
	}
	_, err = cs.db.Exec(sql, args...)
	if err != nil {
		cs.log.Info("Failed to create new channel.", zap.String("workspace_id", channel.WorkspaceId))
		return err
	}
	cs.log.Info("Successfully create a new channel.", zap.String("channel_name", channel.ChannelName))
	return nil
}

func (cs *ChannelStore) GetAllChannelByWorkspaceId(workspaceId string) (adding.Channel, error) {
	sql, args, err := sq.Select("*").From("channel").Where(sq.Eq{"workspace_id": workspaceId}).ToSql()
	if err != nil {
		cs.log.Info("Unable to create select channel sql.", zap.Error(err))
		return adding.Channel{}, err
	}
	var channel adding.Channel
	err = cs.db.Get(&channel, sql, args...)
	if err != nil {
		cs.log.Info("Failed to select channel by workspace id.", zap.Error(err))
		return adding.Channel{}, err
	}
	return channel, nil
}

func (cs *ChannelStore) AddUserToChannel(channelId string, userIdList []string) error {
	sqlBuilder := sq.Insert("channel_user").Columns("channel_id", "user_id")
	for _, id := range userIdList {
		sqlBuilder = sqlBuilder.Values(channelId, id)
	}
	sql, args, err := sqlBuilder.ToSql()
	if err != nil {
		cs.log.Info("Unable to create add user to channel sql.", zap.Error(err))
		return err
	}
	_, err = cs.db.Exec(sql, args...)
	if err != nil {
		cs.log.Info("Failed to add user to channel.", zap.Error(err))
		return err
	}
	cs.log.Info("Successfully add user to channel.", zap.String("channel_id", channelId))
	return nil
}

func (cs *ChannelStore) GetClientListByChannelId(channelId string) ([]chatting.Client, error) {
	sql, args, err := sq.Select("user_id").From("channel_user").Where(sq.Eq{"channel_id": channelId}).ToSql()
	if err != nil {
		cs.log.Info("Unable to create select client list sql.", zap.Error(err))
		return nil, err
	}
	var clientList []chatting.Client
	err = cs.db.Select(&clientList, sql, args...)
	if err != nil {
		cs.log.Info("Failed to select client list.", zap.Error(err))
		return nil, err
	}
	return clientList, err
}
