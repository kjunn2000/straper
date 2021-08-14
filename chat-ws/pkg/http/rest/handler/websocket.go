package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
	"go.uber.org/zap"
)

var Upgrader websocket.Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (server *Server) SetUpConnectionRouter(mr *mux.Router, cs chatting.Service) {
	cr := mr.PathPrefix("/protected").Subrouter()
	cr.HandleFunc("/ws", server.HandleUpgrade(cs))
	cr.Use(middleware.TokenVerifier(server.tokenMaker))
}

func (server *Server) HandleUpgrade(cs chatting.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		conn, err := Upgrader.Upgrade(rw, r, nil)
		if err != nil {
			server.log.Warn("Cannot upgrade to websocket connection.", zap.Error(err))
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		payloadVal := r.Context().Value(middleware.TokenPayload{})
		if payloadVal == nil {
			server.log.Warn("Cannot get payload data.")
			return
		}
		payload, ok := payloadVal.(*auth.Payload)
		if !ok {
			server.log.Warn("Cannot cast to payload strct.")
			return
		}
		err = cs.SetUpUserConnection(r.Context(), payload.UserId, conn)
		if err != nil {
			server.log.Warn("Connection cannot save to redis cache.", zap.Error(err))
			return
		}
		server.log.Info("Successful open websocket connection.", zap.String("user_id", payload.UserId))
	}
}
