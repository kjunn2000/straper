package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	cr.HandleFunc("/create", server.CreateChannel(as, ls, cs)).Methods("POST")
	cr.HandleFunc("/join", server.JoinChannel(as, ls, cs)).Methods("POST")
	cr.HandleFunc("/update", server.UpdateChannel(es)).Methods("POST")
	cr.HandleFunc("/delete/{channel_id}", server.DeleteChannel(ls, ds, cs)).Methods("POST")
	cr.HandleFunc("/leave/{channel_id}", server.LeaveChannel(ds)).Methods("POST")
	cr.HandleFunc("/{channel_id}/messages", server.GetChannelMessages(cs)).Methods("GET")
	cr.Use(middleware.TokenVerifier(server.tokenMaker))
}

type CreateChannelRequest struct {
	WorkspaceId string `json:"workspace_id"`
	ChannelName string `json:"channel_name"`
}

type JoinChannelRequest struct {
	WorkspaceId string `json:"workspace_id"`
	ChannelId   string `json:"channel_id"`
}

func (server *Server) CreateChannel(as adding.Service, ls listing.Service, cs chatting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var cq CreateChannelRequest
		err := json.NewDecoder(r.Body).Decode(&cq)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		userId, err := server.getUserIdFromToken(r)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		if _, err := ls.GetWorkspaceByWorkspaceId(r.Context(), cq.WorkspaceId); err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		channel, err := as.CreateChannel(r.Context(), cq.WorkspaceId, cq.ChannelName, userId)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, channel, "")
	}
}

func (server *Server) JoinChannel(as adding.Service, ls listing.Service, cs chatting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var cq JoinChannelRequest
		err := json.NewDecoder(r.Body).Decode(&cq)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		userId, err := server.getUserIdFromToken(r)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		channel, err := ls.VerifyAndGetChannel(r.Context(), cq.WorkspaceId, cq.ChannelId)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		if err := as.AddUserToChannel(r.Context(), cq.ChannelId, []string{userId}); err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, channel, "")
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

func (server *Server) DeleteChannel(ls listing.Service, ds deleting.Service, cs chatting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := server.getUserIdFromToken(r)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		vars := mux.Vars(r)
		channelId, ok := vars["channel_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "channel.id.not.found")
			return
		}
		if err := ls.VerfiyDeleteChannel(r.Context(), channelId, userId); err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		if err := cs.DeleteSeaweedfsMessagesByChannelId(r.Context(), channelId); err != nil {
			rest.AddResponseToResponseWritter(w, nil, "failed.to.delete.files")
			return
		}
		if err := ds.DeleteChannel(r.Context(), channelId); err != nil {
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
		userId, err := server.getUserIdFromToken(r)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		err = ds.LeaveChannel(r.Context(), channelId, userId)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}

func (server *Server) GetChannelMessages(cs chatting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		channelId, ok := vars["channel_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "channel.id.not.found")
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
		userId, err := server.getUserIdFromToken(r)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		msgs, err := cs.GetChannelMessages(r.Context(), channelId, userId, limit, offset)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, msgs, "")
	}
}
