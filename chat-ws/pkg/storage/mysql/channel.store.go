package mysql

import (
	"context"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"go.uber.org/zap"
)

func (s *SQLStore) CreateNewChannel(ctx context.Context, channel adding.Channel, userId string) (adding.Channel, error) {
	err := s.execTx(func(q *Queries) error {
		err := q.CreateChannel(ctx, channel)
		if err != nil {
			return err
		}
		err = q.AddUserToChannel(ctx, channel.ChannelId, []string{userId})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		s.log.Info("Failed to create a new channel.", zap.Error(err))
		return adding.Channel{}, err
	}
	return channel, nil
}
