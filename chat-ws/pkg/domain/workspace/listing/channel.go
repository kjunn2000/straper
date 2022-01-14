package listing

import "time"

type Channel struct {
	ChannelId    string    `json:"channel_id" db:"channel_id"`
	ChannelName  string    `json:"channel_name" db:"channel_name"`
	WorkspaceId  string    `json:"workspace_id" db:"workspace_id"`
	CreatorId    string    `json:"creator_id" db:"creator_id"`
	CreatedDate  time.Time `json:"created_date" db:"created_date"`
	IsDefault    bool      `json:"is_default" db:"is_default"`
	LastAccessed time.Time `json:"last_accessed" db:"last_accessed"`
}
