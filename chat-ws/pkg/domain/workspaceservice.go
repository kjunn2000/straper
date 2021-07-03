package domain

import (
	"go.uber.org/zap"
)

type WorkspaceService interface {
	CreateWorkspace(name string) error
	EditWorkspace(w Workspace) error
	DeleteWorkspace(id string) error
	GetWorkspaces() ([]Workspace, error)
	GetWorkspace(id string) (Workspace, error)
}

type WorkspaceRepository interface {
	CreateWorkspace(w Workspace) error
	EditWorkspace(w Workspace) error
	DeleteWorkspace(id string) error
	GetWorkspaces() ([]Workspace, error)
	GetWorkspace(id string) (Workspace, error)
}

type workspaceService struct {
	s   WorkspaceRepository
	log *zap.Logger
}

func NewWorkspaceService(r WorkspaceRepository, log *zap.Logger) *workspaceService {
	return &workspaceService{
		s:   r,
		log: log,
	}
}

func (ws *workspaceService) CreateWorkspace(name string) error {
	newWorkspace := Workspace{
		Id:   "",
		Name: name,
	}
	return ws.s.CreateWorkspace(newWorkspace)
}

func (ws *workspaceService) EditWorkspace(w Workspace) error {
	return ws.s.EditWorkspace(w)
}

func (ws *workspaceService) DeleteWorkspace(id string) error {
	return ws.s.DeleteWorkspace(id)
}

func (ws *workspaceService) GetWorkspace(id string) (Workspace, error) {
	return ws.s.GetWorkspace(id)
}

func (ws *workspaceService) GetWorkspaces() ([]Workspace, error) {
	return ws.s.GetWorkspaces()
}
