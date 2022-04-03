package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	ws "github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	"go.uber.org/zap"
)

var Upgrader websocket.Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (server *Server) SetUpWebsocketRouter(mr *mux.Router, ws ws.Service) {
	cr := mr.PathPrefix("/protected").Subrouter()
	cr.HandleFunc("/ws/{userId}", server.HandleUpgrade(ws))
}

func (server *Server) HandleUpgrade(ws ws.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId, ok := vars["userId"]
		if !ok {
			rest.AddResponseToResponseWritter(rw, nil, "user.id.not.found")
			return
		}
		conn, err := Upgrader.Upgrade(rw, r, nil)
		if err != nil {
			server.log.Warn("Cannot upgrade to websocket connection.", zap.Error(err))
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		ws.SetUpUserConnection(r.Context(), userId, conn)
		server.log.Info("Successful open websocket connection.", zap.String("user_id", userId))
	}
}
