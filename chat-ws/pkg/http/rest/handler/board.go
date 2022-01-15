package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
)

func (server *Server) SetUpBoardRouter(mr *mux.Router) {
	wr := mr.PathPrefix("/protected/board").Subrouter()
	wr.HandleFunc("/{workspace_id}", server.GetBoardData()).Methods("GET")
	wr.Use(middleware.TokenVerifier(server.tokenMaker))
}

func (server *Server) GetBoardData() func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r)
		// workspaceId, ok := vars["workspace_id"]
		// rest.AddResponseToResponseWritter(rw, nil, "")
	}
}
