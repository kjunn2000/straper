package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain"
	"go.uber.org/zap"
)

type WorkspaceHanlder interface {
	CreateWorkspace(rw http.ResponseWriter, r *http.Request)
	EditWorkspace(rw http.ResponseWriter, r *http.Request)
	DeleteWorkspace(rw http.ResponseWriter, r *http.Request)
	GetWorkspaces(rw http.ResponseWriter, r *http.Request)
	GetWorkspace(rw http.ResponseWriter, r *http.Request)
}

type workspaceHandler struct {
	ws  domain.WorkspaceService
	log *zap.Logger
}

func NewWorkspaceHandler(ws domain.WorkspaceService, log *zap.Logger) *workspaceHandler {
	return &workspaceHandler{
		ws:  ws,
		log: log,
	}
}

func (w *workspaceHandler) CreateWorkspace(rw http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// name, ok := vars["name"]
	// if !ok {
	// 	w.log.Warn("Unable to extract name from create workplace request")
	// 	rw.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// err := w.ws.CreateWorkspace(name)
	// if err != nil {
	// 	rw.WriteHeader(http.StatusBadRequest)
	// }
}

func (w *workspaceHandler) EditWorkspace(rw http.ResponseWriter, r *http.Request) {
	var ws domain.Workspace
	err := json.NewDecoder(r.Body).Decode(&ws)
	if err != nil {
		w.log.Warn("Unable to decode json workspace object.")
		return
	}
	err = w.ws.EditWorkspace(ws)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
}

func (w *workspaceHandler) DeleteWorkspace(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	fmt.Println(id)
	if !ok {
		w.log.Warn("Unable to extract id from delete workplace request")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err := w.ws.DeleteWorkspace(id)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
}

func (w *workspaceHandler) GetWorkspace(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.log.Warn("Unable to extract id from get workplace request")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	workspace, err := w.ws.GetWorkspace(id)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(rw).Encode(&workspace)
	if err != nil {
		w.log.Warn("Unable to encode workspace into json format.")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}
