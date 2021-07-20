package adding

type Channel struct {
	ChannelId   string `json:"channel_id" db:"channel_id"`
	ChannelName string `json:"channel_name" db:"channel_name"`
	WorkspaceId string `json:"workspace_id" db:"workspace_id"`
}

func NewChannel(channelId, channelName, workspaceId string) *Channel {
	return &Channel{
		ChannelId:   channelId,
		ChannelName: channelName,
		WorkspaceId: workspaceId,
	}
}
