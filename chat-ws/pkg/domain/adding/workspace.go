package adding

type Workspace struct {
	Id   string `json:"workspace_id" db:"workspace_id"`
	Name string `json:"workspace_name" db:"workspace_name"`
	CreatorId string `db:"creator_id"`
}