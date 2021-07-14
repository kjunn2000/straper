package domain

type Workspace struct {
	Id       string     `json:"workspaceId" db:"workspace_id"`
	Name     string     `json:"workspaceName" db:"workspace_name"`
	Channels []*Channel
}
