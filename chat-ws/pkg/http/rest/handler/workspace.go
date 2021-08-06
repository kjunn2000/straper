package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/deleting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/listing"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
)

func SetUpWorkspaceRouter(mr *mux.Router, as adding.Service, ls listing.Service, ds deleting.Service) {
	wr := mr.PathPrefix("/protected/workspace").Subrouter()
	wr.HandleFunc("/create", CreateWorkspace(as, ls)).Methods("POST")
	wr.HandleFunc("/join", JoinWorkspace(as, ls)).Methods("POST")
	wr.HandleFunc("/list", GetWorkspaces(ls)).Methods("GET")
	wr.HandleFunc("/delete/{workspace_id}", DeleteWorkspace(ds)).Methods("POST")
	wr.HandleFunc("/leave/{workspace_id}", LeaveWorkspace(ds)).Methods("POST")
	wr.Use(middleware.JwtTokenVerifier)
}

func CreateWorkspace(as adding.Service, ls listing.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		var workspace adding.Workspace
		err := json.NewDecoder(r.Body).Decode(&workspace)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		userId, err := extractUserId(r)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		workspace, err = as.CreateWorkspace(workspace, userId)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		w, err := ls.GetWorkspaceByWorkspaceId(workspace.Id)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, w, "")
	}
}

func JoinWorkspace(as adding.Service, ls listing.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		var workspace adding.Workspace
		err := json.NewDecoder(r.Body).Decode(&workspace)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		userId, err := extractUserId(r)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		err = as.AddUserToWorkspace(workspace.Id, []string{userId})
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		w, err := ls.GetWorkspaceByWorkspaceId(workspace.Id)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, w, "")
	}
}

func GetWorkspaces(ls listing.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		userId, err := extractUserId(r)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		workspaceList, err := ls.GetWorkspaceData(userId)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, workspaceList, "")
	}
}

func DeleteWorkspace(ds deleting.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["workspace_id"]
		if !ok {
			rw.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(rw, nil, "Id not found.")
			return
		}
		err := ds.DeleteWorkspace(id)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, nil, "")
	}
}

func LeaveWorkspace(ds deleting.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["workspace_id"]
		if !ok {
			rest.AddResponseToResponseWritter(rw, nil, "Id not found.")
			return
		}
		userId, err := extractUserId(r)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, "Id not found.")
			return
		}
		err = ds.LeaveWorkspace(id, userId)
		if err != nil {
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

func extractUserId(r *http.Request) (string, error) {
	accessToken := r.Header.Get("Authorization")
	claims, err := auth.ExtractClaimsFromTokenStr(accessToken)
	if err != nil {
		return "", err
	}
	return claims.UserId, nil
}
