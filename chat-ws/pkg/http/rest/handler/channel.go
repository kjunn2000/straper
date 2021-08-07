package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
)

func SetUpChannelRouter(mr *mux.Router, as adding.Service, cs chatting.Service) {
	cr := mr.PathPrefix("/protected/channel").Subrouter()
	cr.HandleFunc("/create", CreateChannel(as, cs)).Methods("POST")
	cr.Use(middleware.JwtTokenVerifier)
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
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		accessToken := r.Header.Get("Authorization")
		claims, err := auth.ExtractClaimsFromTokenStr(accessToken)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		channel, err := as.CreateChannel(cq.WorkspaceId, cq.ChannelName, claims.UserId)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, channel, "")
	}
}
