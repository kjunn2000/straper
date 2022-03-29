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
	br.HandleFunc("/card/comments/{card_id}", server.GetCardComments(bs)).Methods("GET")
	br.Use(middleware.TokenVerifier(server.tokenMaker))
	br.Use(middleware.UpdateLastSeen(server.statusLogger))
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

func (server *Server) GetCardComments(bs board.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cardId, ok := vars["card_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "card.id.not.found")
			return
		}
		param := board.PaginationCommentParam{
			Cursor: r.URL.Query().Get("cursor"),
		}
		msgs, err := bs.GetCardComments(r.Context(), cardId, param)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, msgs, "")
	}
}
