package listing

type Workspace struct {
	Id          string    `json:"workspace_id" db:"workspace_id"`
	Name        string    `json:"workspace_name" db:"workspace_name"`
	CreatorId   string    `json:"creator_id" db:"creator_id"`
	ChannelList []Channel `json:"channel_list"`
}
