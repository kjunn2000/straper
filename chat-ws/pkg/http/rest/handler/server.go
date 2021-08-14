package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/configs"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/deleting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/editing"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/workspace/listing"
	"go.uber.org/zap"

	"github.com/kjunn2000/straper/chat-ws/pkg/storage/mysql"
	rdb "github.com/kjunn2000/straper/chat-ws/pkg/storage/redis"
)

type Server struct {
	log         *zap.Logger
	config      configs.Config
	store       mysql.Store
	httpServer  *http.Server
	tokenMaker  auth.Maker
	redisClient rdb.RedisClient
}

func NewServer(log *zap.Logger, config configs.Config, store mysql.Store) (*Server, error) {

	tokenMaker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Warn("Unable to create token maker.", zap.Error(err))
	}

	rc := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	redisClient := rdb.NewRedisClient(rc, log)

	srv := &http.Server{
		Addr:         config.ServerAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server := &Server{
		log:         log,
		httpServer:  srv,
		config:      config,
		store:       store,
		tokenMaker:  tokenMaker,
		redisClient: &redisClient,
	}

	server.SetServerRoute()
	return server, nil
}

func (server *Server) StartServer() {

	server.log.Info("Server is running.", zap.String("port", server.config.ServerAddress))

	err := server.httpServer.ListenAndServe()
	if err != nil {
		server.log.Warn("Unable to start server.")
		return
	}
}

func (server *Server) SetServerRoute() (*mux.Router, error) {

	mr := mux.NewRouter().PathPrefix("/api/v1").Subrouter()

	accountService := account.NewService(server.log, server.store)
	authService := auth.NewService(server.log, server.store, server.tokenMaker, server.config)
	addingService := adding.NewService(server.log, server.store)
	chattingService := chatting.NewService(server.log, server.store, server.redisClient)
	chattingService.SetUpWSServer(context.Background())
	listingService := listing.NewService(server.log, server.store)
	editingService := editing.NewService(server.log, server.store)
	deletingService := deleting.NewService(server.log, server.store)

	server.SetUpAuthRouter(mr, authService)
	server.SetUpAccountRouter(mr, accountService)
	server.SetUpWorkspaceRouter(mr, addingService, listingService, editingService, deletingService)
	server.SetUpChannelRouter(mr, addingService, listingService, editingService, deletingService, chattingService)
	server.SetUpConnectionRouter(mr, chattingService)

	server.httpServer.Handler = getCORSHandler()(mr)
	return mr, nil
}

func getCORSHandler() func(http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type", "Accept", "Authorization"}),
		handlers.AllowCredentials(),
	)
}
