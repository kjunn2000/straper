package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/deleting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/editing"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/listing"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
)

func (server *Server) SetUpWorkspaceRouter(mr *mux.Router, as adding.Service, ls listing.Service, es editing.Service, ds deleting.Service) {
	wr := mr.PathPrefix("/protected/workspace").Subrouter()
	wr.HandleFunc("/create", server.CreateWorkspace(as)).Methods("POST")
	wr.HandleFunc("/join/{workspace_id}", server.JoinWorkspace(as)).Methods("POST")
	wr.HandleFunc("/list", server.GetWorkspaces(ls)).Methods("GET")
	wr.HandleFunc("/update", server.UpdateWorkspace(es)).Methods("POST")
	wr.HandleFunc("/delete/{workspace_id}", server.DeleteWorkspace(ds)).Methods("POST")
	wr.HandleFunc("/leave/{workspace_id}", server.LeaveWorkspace(ds)).Methods("POST")
	wr.Use(middleware.TokenVerifier(server.tokenMaker))
}

func (server *Server) CreateWorkspace(as adding.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		var workspace adding.Workspace
		err := json.NewDecoder(r.Body).Decode(&workspace)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		payloadVal := r.Context().Value(middleware.TokenPayload{})
		if payloadVal == nil {
			rest.AddResponseToResponseWritter(rw, nil, "payload.not.found.in.context")
			return
		}
		payload, ok := payloadVal.(*auth.Payload)
		if !ok {
			rest.AddResponseToResponseWritter(rw, nil, "invalid.payload.in.context")
			return
		}
		_, err = as.CreateWorkspace(r.Context(), workspace, payload.UserId)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, nil, "")
	}
}

func (server *Server) JoinWorkspace(as adding.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workspaceId, ok := vars["workspace_id"]
		if !ok {
			rest.AddResponseToResponseWritter(rw, nil, "workspace.id.not.found")
			return
		}
		payloadVal := r.Context().Value(middleware.TokenPayload{})
		if payloadVal == nil {
			rest.AddResponseToResponseWritter(rw, nil, "payload.not.found.in.context")
			return
		}
		payload, ok := payloadVal.(*auth.Payload)
		if !ok {
			rest.AddResponseToResponseWritter(rw, nil, "invalid.payload.in.context")
			return
		}
		err := as.AddUserToWorkspace(r.Context(), workspaceId, []string{payload.UserId})
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, nil, "")
	}
}

func (server *Server) GetWorkspaces(ls listing.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		payloadVal := r.Context().Value(middleware.TokenPayload{})
		if payloadVal == nil {
			rest.AddResponseToResponseWritter(rw, nil, "payload.not.found.in.context")
			return
		}
		payload, ok := payloadVal.(*auth.Payload)
		if !ok {
			rest.AddResponseToResponseWritter(rw, nil, "invalid.payload.in.context")
			return
		}
		workspaceList, err := ls.GetWorkspaceData(r.Context(), payload.UserId)
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

func (server *Server) DeleteWorkspace(ds deleting.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["workspace_id"]
		if !ok {
			rw.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(rw, nil, "Id not found.")
			return
		}
		err := ds.DeleteWorkspace(r.Context(), id)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, nil, "")
	}
}

func (server *Server) LeaveWorkspace(ds deleting.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["workspace_id"]
		if !ok {
			rest.AddResponseToResponseWritter(rw, nil, "Id not found.")
			return
		}
		payloadVal := r.Context().Value(middleware.TokenPayload{})
		if payloadVal == nil {
			rest.AddResponseToResponseWritter(rw, nil, "payload.not.found.in.context")
			return
		}
		payload, ok := payloadVal.(*auth.Payload)
		if !ok {
			rest.AddResponseToResponseWritter(rw, nil, "invalid.payload.in.context")
			return
		}
		err := ds.LeaveWorkspace(r.Context(), id, payload.UserId)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, nil, "")
	}
}
