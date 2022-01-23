package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/board"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
)

func (server *Server) SetUpBoardRouter(mr *mux.Router, bs board.Service) {
	br := mr.PathPrefix("/protected/board").Subrouter()
	br.HandleFunc("/{workspace_id}", server.GetBoardData(bs)).Methods("GET")
	br.Use(middleware.TokenVerifier(server.tokenMaker))
}

func (server *Server) GetBoardData(bs board.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workspaceId, ok := vars["workspace_id"]
		if !ok {
			rest.AddResponseToResponseWritter(rw, nil, "workspace.id.not.found")
			return
		}
		taskBoardResp, err := bs.GetTaskBoardData(r.Context(), workspaceId)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(rw, taskBoardResp, "")
	}
}
