package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/deleting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/editing"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/listing"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
)

func (server *Server) SetUpChannelRouter(mr *mux.Router, as adding.Service, ls listing.Service, es editing.Service,
	ds deleting.Service, cs chatting.Service) {
	cr := mr.PathPrefix("/protected/channel").Subrouter()
	cr.HandleFunc("/create", server.CreateChannel(as, cs)).Methods("POST")
	cr.HandleFunc("/join/{channel_id}", server.JoinChannel(as, cs)).Methods("POST")
	cr.HandleFunc("/update", server.UpdateChannel(es)).Methods("POST")
	cr.HandleFunc("/delete/{channel_id}", server.DeleteChannel(ds)).Methods("POST")
	cr.HandleFunc("/leave/{channel_id}", server.LeaveChannel(ds)).Methods("POST")
	cr.Use(middleware.TokenVerifier(server.tokenMaker))
}

type CreateChannelRequest struct {
	WorkspaceId string `json:"workspace_id"`
	ChannelName string `json:"channel_name"`
}

func (server *Server) CreateChannel(as adding.Service, cs chatting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var cq CreateChannelRequest
		err := json.NewDecoder(r.Body).Decode(&cq)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		payloadVal := r.Context().Value(middleware.TokenPayload{})
		if payloadVal == nil {
			rest.AddResponseToResponseWritter(w, nil, "payload.not.found.in.context")
			return
		}
		payload, ok := payloadVal.(*auth.Payload)
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "invalid.payload.in.context")
			return
		}
		channel, err := as.CreateChannel(r.Context(), cq.WorkspaceId, cq.ChannelName, payload.UserId)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, channel, "")
	}
}

func (server *Server) JoinChannel(as adding.Service, cs chatting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		channelId, ok := vars["channel_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "channel.id.not.found")
			return
		}
		payloadVal := r.Context().Value(middleware.TokenPayload{})
		if payloadVal == nil {
			rest.AddResponseToResponseWritter(w, nil, "payload.not.found.in.context")
			return
		}
		payload, ok := payloadVal.(*auth.Payload)
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "invalid.payload.in.context")
			return
		}
		err := as.AddUserToChannel(r.Context(), channelId, []string{payload.UserId})
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}

func (server *Server) UpdateChannel(es editing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var channel editing.Channel
		err := json.NewDecoder(r.Body).Decode(&channel)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		err = es.UpdateChannel(r.Context(), channel)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}

func (server *Server) DeleteChannel(ds deleting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		payloadVal := r.Context().Value(middleware.TokenPayload{})
		if payloadVal == nil {
			rest.AddResponseToResponseWritter(w, nil, "payload.not.found.in.context")
			return
		}
		payload, ok := payloadVal.(*auth.Payload)
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "invalid.payload.in.context")
			return
		}
		vars := mux.Vars(r)
		channelId, ok := vars["channel_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "channel.id.not.found")
			return
		}
		err := ds.DeleteChannel(r.Context(), channelId, payload.UserId)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}

func (server *Server) LeaveChannel(ds deleting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		channelId, ok := vars["channel_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "channel.id.not.found")
			return
		}
		payloadVal := r.Context().Value(middleware.TokenPayload{})
		if payloadVal == nil {
			rest.AddResponseToResponseWritter(w, nil, "payload.not.found.in.context")
			return
		}
		payload, ok := payloadVal.(*auth.Payload)
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "invalid.payload.in.context")
			return
		}
		err := ds.LeaveChannel(r.Context(), channelId, payload.UserId)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}
