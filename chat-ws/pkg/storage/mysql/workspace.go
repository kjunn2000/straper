package storage

type Workspace struct {
	id   string `db:"workspace_id"`
	name string `db:"workspace_name"`
}