package listing

type Channel struct {
	ChannelId   string `json:"channel_id" db:"channel_id"`
	ChannelName string `json:"channel_name" db:"channel_name"`
}