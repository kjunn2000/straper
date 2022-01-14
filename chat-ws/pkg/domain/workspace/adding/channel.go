package adding

import "time"

type Channel struct {
	ChannelId   string    `json:"channel_id" db:"channel_id"`
	ChannelName string    `json:"channel_name" db:"channel_name"`
	WorkspaceId string    `json:"workspace_id" db:"workspace_id"`
	CreatorId   string    `json:"creator_id" db:"creator_id"`
	IsDefault   bool      `json:"is_default" db:"is_default"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
}

func NewChannel(channelId string, channelName string, workspaceId string,
	creatorId string, isDefault bool, createdDate time.Time) Channel {
	return Channel{
		ChannelId:   channelId,
		ChannelName: channelName,
		WorkspaceId: workspaceId,
		CreatorId:   creatorId,
		IsDefault:   isDefault,
		CreatedDate: createdDate,
	}
}
