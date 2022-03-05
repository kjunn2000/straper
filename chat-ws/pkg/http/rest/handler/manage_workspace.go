package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/admin"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/deleting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/editing"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/listing"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
)

func (server *Server) SetUpManageWorkspaceRouter(mr *mux.Router, as admin.Service, es editing.Service,
	ls listing.Service, ds deleting.Service, cs chatting.Service) {
	r := mr.PathPrefix("/admin/protected/workspace").Subrouter()
	r.HandleFunc("/read/{workspace_id}", server.GetWorkspace(as)).Methods("GET")
	r.HandleFunc("/list", server.GetPaginationWorkspace(as)).Methods("GET")
	r.HandleFunc("/update", server.UpdateWorkspace(es)).Methods("POST")
	r.HandleFunc("/delete/{workspace_id}", server.DeleteWorkspace(ls, ds, cs, true)).Methods("POST")
	r.HandleFunc("/remove/{workspace_id}/{user_id}", server.RemoveUserFromWorkspace(ds)).Methods("POST")
	r.HandleFunc("/channel/update", server.UpdateChannel(es)).Methods("POST")
	r.HandleFunc("/channel/delete/{channel_id}", server.DeleteChannel(ls, ds, cs, true)).Methods("POST")
	r.Use(middleware.TokenVerifier(server.tokenMaker))
	r.Use(middleware.UpdateLastSeen(server.statusLogger))
}

func (server *Server) GetWorkspace(as admin.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workspaceId, ok := vars["workspace_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "invalid.workspace.id")
			return
		}
		workspace, err := as.GetWorkspace(r.Context(), workspaceId)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, workspace, "")
	}
}

func (server *Server) GetPaginationWorkspace(as admin.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 64)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, "invalid.limit")
			return
		}
		cursor := r.URL.Query().Get("cursor")
		isNext, err := strconv.ParseBool(r.URL.Query().Get("isNext"))
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, "invalid.is.next.attribute")
			return
		}
		searchStr := r.URL.Query().Get("searchStr")
		param := admin.PaginationWorkspacesParam{
			Limit:     limit,
			Cursor:    cursor,
			IsNext:    isNext,
			SearchStr: searchStr,
		}
		users, err := as.GetPaginationWorkspaces(r.Context(), param)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, users, "")
	}
}

func (server *Server) RemoveUserFromWorkspace(ds deleting.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workspaceId, ok := vars["workspace_id"]
		if !ok {
			rest.AddResponseToResponseWritter(rw, nil, "Workspace ID not found.")
			return
		}
		userId, ok := vars["user_id"]
		if !ok {
			rest.AddResponseToResponseWritter(rw, nil, "User ID not found.")
			return
		}
		err := ds.LeaveWorkspace(r.Context(), workspaceId, userId)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, nil, "")
	}
}
