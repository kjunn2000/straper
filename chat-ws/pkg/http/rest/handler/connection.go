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

func SetUpConnectionRouter(mr *mux.Router, log *zap.Logger, cs chatting.Service) {
	cr := mr.PathPrefix("/protected/connection").Subrouter()
	cr.HandleFunc("/upgrade", HandleUpgrade(log, cs)).Methods("POST")
	cr.Use(middleware.JwtTokenVerifier)
}

func HandleUpgrade(log *zap.Logger, cs chatting.Service) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		conn, err := Upgrader.Upgrade(rw, r, nil)
		if err != nil {
			log.Warn("Cannot upgrade to websocket connection.", zap.Error(err))
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		accessToken := r.Header.Get("Authorization")
		claims, err := auth.ExtractClaimsFromTokenStr(accessToken)
		if err != nil {
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		err = cs.SaveConnectionToCache(claims.UserId, conn)
		if err != nil {
			log.Warn("Connection cannot save to redis cache.", zap.Error(err))
			rest.AddResponseToResponseWritter(rw, nil, err.Error())
			return
		}
		log.Info("Successful created websocket connection.")
		rest.AddResponseToResponseWritter(rw, nil, "")
	}
}