package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain"
	"go.uber.org/zap"
)

var Upgrader websocket.Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type ChannelHandler interface {
	CreateChannel(w http.ResponseWriter, r *http.Request)
	DeleteChannel(w http.ResponseWriter, r *http.Request)
}

type channelHandler struct {
	log *zap.Logger
	cs  domain.ChannelService
}

func NewChannelHandler(log *zap.Logger, cs domain.ChannelService) *channelHandler {
	return &channelHandler{
		log: log,
		cs:  cs,
	}
}

type ChannelRequest struct {
	workspaceId string `json:"workspaceId"`
	channelName string `json:"channelName"`
}

func (ch *channelHandler) CreateChannel(w http.ResponseWriter, r *http.Request) {
	var cq ChannelRequest
	err := json.NewDecoder(r.Body).Decode(&cq)
	if err != nil {
		ch.log.Warn("Unable to decode create channel request.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = ch.cs.CreateChannel(cq.workspaceId, cq.channelName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (ch *channelHandler) handleUpgrade(w http.ResponseWriter, r *http.Request) {
	var cq ChannelRequest
	err := json.NewDecoder(r.Body).Decode(&cq)
	if err != nil {
		ch.log.Warn("Unable to decode create channel request.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		ch.log.Warn("Cannot upgrade to websocket connection : %s", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ch.log.Info("Successful created websocket connection.")
}
