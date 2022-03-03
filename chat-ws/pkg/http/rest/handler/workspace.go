package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/deleting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/editing"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/listing"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
)

func (server *Server) SetUpWorkspaceRouter(mr *mux.Router, as adding.Service, ls listing.Service, es editing.Service, ds deleting.Service, cs chatting.Service) {
	wr := mr.PathPrefix("/protected/workspace").Subrouter()
	wr.HandleFunc("/create", server.CreateWorkspace(as)).Methods("POST")
	wr.HandleFunc("/join/{workspace_id}", server.JoinWorkspace(as, ls)).Methods("POST")
	wr.HandleFunc("/list", server.GetWorkspaces(ls)).Methods("GET")
	wr.HandleFunc("/update", server.UpdateWorkspace(es)).Methods("POST")
	wr.HandleFunc("/delete/{workspace_id}", server.DeleteWorkspace(ls, ds, cs, false)).Methods("POST")
	wr.HandleFunc("/leave/{workspace_id}", server.LeaveWorkspace(ds)).Methods("POST")
	wr.Use(middleware.TokenVerifier(server.tokenMaker))
}

type AddWorkspaceParam struct {
	Name string `json:"workspace_name" db:"workspace_name"`
}

func (server *Server) CreateWorkspace(as adding.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		var workspace AddWorkspaceParam
		err := json.NewDecoder(r.Body).Decode(&workspace)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}

		userId, err := server.getUserIdFromToken(r)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		w, err := as.CreateWorkspace(r.Context(), workspace.Name, userId)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, w, "")
	}
}

func (server *Server) JoinWorkspace(as adding.Service, ls listing.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workspaceId, ok := vars["workspace_id"]
		if !ok {
			rest.AddResponseToResponseWritter(rw, nil, "workspace.id.not.found")
			return
		}

		userId, err := server.getUserIdFromToken(r)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		err = as.AddUserToWorkspace(r.Context(), workspaceId, []string{userId})
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		w, err := ls.GetWorkspaceByWorkspaceId(r.Context(), workspaceId)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, w, "")
	}
}

func (server *Server) GetWorkspaces(ls listing.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		userId, err := server.getUserIdFromToken(r)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		workspaceList, err := ls.GetWorkspaceData(r.Context(), userId)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, workspaceList, "")
	}
}

func (server *Server) UpdateWorkspace(as editing.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		var workspace editing.Workspace
		err := json.NewDecoder(r.Body).Decode(&workspace)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		err = as.UpdateWorkspace(r.Context(), workspace)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, nil, "")
	}
}

func (server *Server) DeleteWorkspace(ls listing.Service, ds deleting.Service, cs chatting.Service, isAdmin bool) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workspaceId, ok := vars["workspace_id"]
		if !ok {
			rest.AddResponseToResponseWritter(rw, nil, "Id not found.")
			return
		}
		if !isAdmin {
			userId, err := server.getUserIdFromToken(r)
			if err != nil {
				rest.AddResponseToResponseWritter(rw, nil, err.Error())
				return
			}
			if err := ls.VerfiyDeleteWorkspace(r.Context(), workspaceId, userId); err != nil {
				rest.AddResponseToResponseWritter(rw, nil, err.Error())
				return
			}
		}
		if err := cs.DeleteSeaweedfsMessagesByWorkspaceId(r.Context(), workspaceId); err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		if err := ds.DeleteWorkspace(r.Context(), workspaceId); err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, nil, "")
	}
}

func (server *Server) LeaveWorkspace(ds deleting.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workspaceId, ok := vars["workspace_id"]
		if !ok {
			rest.AddResponseToResponseWritter(rw, nil, "Workspace ID not found.")
			return
		}
		userId, err := server.getUserIdFromToken(r)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		err = ds.LeaveWorkspace(r.Context(), workspaceId, userId)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, nil, "")
	}
}
