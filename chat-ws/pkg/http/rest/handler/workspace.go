package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
)

func NewWorkspaceRouter(as adding.Service) *mux.Router {
	mr := mux.NewRouter()
	mr.HandleFunc("/create", CreateWorkspace(as)).Methods("POST")
	mr.Use(middleware.JwtTokenVerifier)
	return mr
}

func CreateWorkspace(as adding.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		var workspace adding.Workspace
		err := json.NewDecoder(r.Body).Decode(&workspace)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		accessToken := r.Header.Get("Authorization")
		claims, err := auth.ExtractClaimsFromTokenStr(accessToken)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		err = as.CreateWorkspace(workspace, claims.UserId)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, nil, "")
	}
}

// func (w *workspaceHandler) EditWorkspace(rw http.ResponseWriter, r *http.Request) {
// 	var ws domain.Workspace
// 	err := json.NewDecoder(r.Body).Decode(&ws)
// 	if err != nil {
// 		rw.WriteHeader(http.StatusBadRequest)
// 		rest.AddResponseToResponseWritter(rw, nil, err.Error())
// 		return
// 	}
// 	err = w.ws.EditWorkspace(ws)
// 	if err != nil {
// 		rw.WriteHeader(http.StatusBadRequest)
// 		rest.AddResponseToResponseWritter(rw, nil, err.Error())
// 		return
// 	}
// 	rest.AddResponseToResponseWritter(rw, nil, "")
// }

// func (w *workspaceHandler) DeleteWorkspace(rw http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, ok := vars["id"]
// 	if !ok {
// 		rw.WriteHeader(http.StatusBadRequest)
// 		rest.AddResponseToResponseWritter(rw, nil, "Id not found.")
// 		return
// 	}
// 	err := w.ws.DeleteWorkspace(id)
// 	if err != nil {
// 		rw.WriteHeader(http.StatusBadRequest)
// 		rest.AddResponseToResponseWritter(rw, nil, err.Error())
// 		return
// 	}
// 	rest.AddResponseToResponseWritter(rw, nil, "")
// }

// func (w *workspaceHandler) GetWorkspaces(rw http.ResponseWriter, r *http.Request) {
// 	accessToken := r.Header.Get("Authorization")
// 	claims, err := auth.ExtractClaimsFromTokenStr(accessToken)
// 	if err != nil {
// 		rw.WriteHeader(http.StatusBadRequest)
// 		rest.AddResponseToResponseWritter(rw, nil, err.Error())
// 		return
// 	}
// 	workspaceList, err := w.ws.GetWorkspaces(claims.UserId)
// 	if err != nil {
// 		rw.WriteHeader(http.StatusBadRequest)
// 		rest.AddResponseToResponseWritter(rw, nil, err.Error())
// 		return
// 	}
// 	responseData := WorkspacListModel{
// 		WorkspaceList: workspaceList,
// 	}
// 	rest.AddResponseToResponseWritter(rw, responseData, "")
// }
