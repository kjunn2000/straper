package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/configs"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/admin"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/board"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/bug"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/dblog"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/pagination"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/websocket"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/deleting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/editing"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/listing"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
	"github.com/kjunn2000/straper/chat-ws/pkg/seaweedfs"
	"go.uber.org/zap"

	rdb "github.com/kjunn2000/straper/chat-ws/pkg/redis"
	"github.com/kjunn2000/straper/chat-ws/pkg/storage/mysql"
)

type Server struct {
	log             *zap.Logger
	config          configs.Config
	store           mysql.Store
	httpServer      *http.Server
	tokenMaker      auth.Maker
	statusLogger    dblog.StatusLogger
	redisClient     rdb.RedisClient
	seaweedfsClient chatting.SeaweedfsClient
}

func NewServer(log *zap.Logger, config configs.Config, store mysql.Store) (*Server, error) {

	tokenMaker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Warn("Unable to create token maker.", zap.Error(err))
	}

	statusLogger := dblog.NewStatusLogger(log, store)

	rc := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err = rc.Set(context.Background(), "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	redisClient := rdb.NewRedisClient(rc)

	seaweedClient := seaweedfs.NewSeaweedfsClient(log)

	srv := &http.Server{
		Addr:         config.ServerAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server := &Server{
		log:             log,
		httpServer:      srv,
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		statusLogger:    statusLogger,
		redisClient:     &redisClient,
		seaweedfsClient: &seaweedClient,
	}

	server.SetServerRoute()
	return server, nil
}

func (server *Server) StartServer() {

	server.log.Info("Server is running.", zap.String("port", server.config.ServerAddress))

	err := server.httpServer.ListenAndServe()
	if err != nil {
		server.log.Warn("Unable to start server.", zap.Error(err))
		return
	}
}

func (server *Server) SetServerRoute() (*mux.Router, error) {

	mr := mux.NewRouter().PathPrefix("/api/v1").Subrouter()

	paginationService := pagination.Service{}
	accountService := account.NewService(server.log, server.store, server.config)
	// accountService.SeedUserAccount()
	authService := auth.NewService(server.log, server.store, server.tokenMaker, server.config)
	addingService := adding.NewService(server.log, server.store)
	// addingService.SeedWorkspaces()
	chattingService := chatting.NewService(server.log, server.store, server.seaweedfsClient, paginationService)
	boardService := board.NewService(server.log, server.store, server.seaweedfsClient, paginationService)
	listingService := listing.NewService(server.log, server.store)
	editingService := editing.NewService(server.log, server.store)
	deletingService := deleting.NewService(server.log, server.store, server.seaweedfsClient)
	bugService := bug.NewService(server.log, server.store, server.seaweedfsClient)
	websocketService := websocket.NewService(server.log, server.redisClient, chattingService, boardService)
	adminService := admin.NewService(server.log, server.store, paginationService)
	websocketService.SetUpWSServer(context.Background())

	server.SetUpAuthRouter(mr, authService)
	server.SetUpAccountRouter(mr, accountService)
	server.SetUpWorkspaceRouter(mr, addingService, listingService, editingService, deletingService, chattingService)
	server.SetUpChannelRouter(mr, addingService, listingService, editingService, deletingService, chattingService)
	server.SetUpBoardRouter(mr, boardService)
	server.SetUpBugRouter(mr, bugService)
	server.SetUpWebsocketRouter(mr, websocketService)
	server.SetUpManageUserRouter(mr, adminService)
	server.SetUpManageWorkspaceRouter(mr, adminService, editingService, listingService, deletingService, chattingService)

	server.httpServer.Handler = getCORSHandler()(mr)
	return mr, nil
}

func getCORSHandler() func(http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:3001"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type", "Accept", "Authorization"}),
		handlers.AllowCredentials(),
	)
}

func (server *Server) getUserIdFromToken(r *http.Request) (string, error) {
	payloadVal := r.Context().Value(middleware.TokenPayload{})
	if payloadVal == nil {
		return "", errors.New("payload.not.found.in.context")
	}
	payload, ok := payloadVal.(*auth.Payload)
	if !ok {
		return "", errors.New("invalid.payload.in.context")
	}
	return payload.UserId, nil
}
