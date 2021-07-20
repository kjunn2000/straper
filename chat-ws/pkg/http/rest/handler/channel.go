package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
)

func NewChannelRouter(as adding.Service, cs chatting.Service) *mux.Router {
	mr := mux.NewRouter()
	mr.HandleFunc("", CreateChannel(as, cs)).Methods("POST")
	mr.Use(middleware.JwtTokenVerifier)
	return mr
}

type ChannelRequest struct {
	WorkspaceId string `json:"workspace_id"`
	ChannelName string `json:"channel_name"`
}

func CreateChannel(as adding.Service, cs chatting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var cq ChannelRequest
		err := json.NewDecoder(r.Body).Decode(&cq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		_, err = as.CreateChannel(cq.WorkspaceId, cq.ChannelName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}
