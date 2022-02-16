package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/bug"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
)

func (server *Server) SetUpBugRouter(mr *mux.Router, bs bug.Service) {
	br := mr.PathPrefix("/protected/issue").Subrouter()
	br.HandleFunc("/create", server.CreateIssue(bs)).Methods("POST")
	br.HandleFunc("/list/read/{workspace_id}", server.GetIssues(bs)).Methods("GET")
	br.HandleFunc("/update", server.UpdateIssue(bs)).Methods("POST")
	br.HandleFunc("/delete/{issue_id}", server.DeleteIssue(bs)).Methods("POST")
	br.Use(middleware.TokenVerifier(server.tokenMaker))
}

func (server *Server) CreateIssue(bs bug.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		var issue bug.Issue
		json.NewDecoder(r.Body).Decode(&issue)

		issue, err := bs.CreateIssue(r.Context(), issue)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, issue, "")
	}
}

func (server *Server) GetIssues(bs bug.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workspaceId, ok := vars["workspace_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "workspace.id.not.found")
			return
		}
		limit, err := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 64)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, "invalid.limit")
			return
		}
		offset, err := strconv.ParseUint(r.URL.Query().Get("offset"), 10, 64)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, "invalid.offset")
			return
		}
		issues, err := bs.GetIssuesByWorkspaceId(r.Context(), workspaceId, limit, offset)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, issues, "")
	}
}

func (server *Server) UpdateIssue(bs bug.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		var param bug.UpdateIssueParam
		json.NewDecoder(r.Body).Decode(&param)

		issue, err := bs.UpdateIssue(r.Context(), param)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, issue, "")
	}
}

func (server *Server) DeleteIssue(bs bug.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		issueId, ok := vars["issue_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "issue.id.not.found")
			return
		}
		err := bs.DeleteIssue(r.Context(), issueId)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}
