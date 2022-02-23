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
	br.HandleFunc("/list/{workspace_id}", server.GetIssues(bs)).Methods("GET")
	br.HandleFunc("/update", server.UpdateIssue(bs)).Methods("POST")
	br.HandleFunc("/delete/{issue_id}", server.DeleteIssue(bs)).Methods("POST")

	br.HandleFunc("/epic-link/option/{workspace_id}", server.GetEpicLinkOptions(bs)).Methods("GET")
	br.HandleFunc("/assignee/option/{workspace_id}", server.GetAssigneeOptions(bs)).Methods("GET")

	br.HandleFunc("/attachments/upload", server.AddIssueAttachments(bs)).Methods("POST")
	br.HandleFunc("/attachment/delete/{fid}", server.DeleteIssueAttachment(bs)).Methods("POST")
	br.HandleFunc("/attachment/download/{fid}", server.DownloadAttachment(bs)).Methods("GET")
	br.Use(middleware.TokenVerifier(server.tokenMaker))
}

func (server *Server) CreateIssue(bs bug.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		userId, err := server.getUserIdFromToken(r)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		var issue bug.Issue
		json.NewDecoder(r.Body).Decode(&issue)
		issue.Reporter = userId
		issue.Status = "ACTIVE"
		issue, err = bs.CreateIssue(r.Context(), issue)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, issue, "")
	}
}

func (server *Server) AddIssueAttachments(bs bug.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(32 << 2)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		issueId := r.PostFormValue("issue_id")
		fileTypes := r.PostForm["types"]
		files := r.MultipartForm.File["files"]

		attachments, err := bs.AddIssueAttachments(r.Context(), issueId, fileTypes, files)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, attachments, "")
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
		var issue bug.Issue
		json.NewDecoder(r.Body).Decode(&issue)

		issue, err := bs.UpdateIssue(r.Context(), issue)
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

func (server *Server) DeleteIssueAttachment(bs bug.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fid, ok := vars["fid"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "fid.not.found")
			return
		}
		err := bs.DeleteIssueAttachment(r.Context(), fid)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}

func (server *Server) GetEpicLinkOptions(bs bug.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workspaceId, ok := vars["workspace_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "workspace.id.not.found")
			return
		}
		options, err := bs.GetEpicLinkOptions(r.Context(), workspaceId)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, options, "")
	}
}

func (server *Server) GetAssigneeOptions(bs bug.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workspaceId, ok := vars["workspace_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "workspace.id.not.found")
			return
		}
		options, err := bs.GetAssigneeOptions(r.Context(), workspaceId)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, options, "")
	}
}

func (server *Server) DownloadAttachment(bs bug.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fid, ok := vars["fid"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "fid.not.found")
			return
		}
		bytes, err := bs.GetAttachment(r.Context(), fid)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		w.Write(bytes)
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}
