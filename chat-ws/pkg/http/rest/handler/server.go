package handler

import (
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/configs"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/deleting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/listing"
	"go.uber.org/zap"

	"github.com/kjunn2000/straper/chat-ws/pkg/storage/mysql"
	rdb "github.com/kjunn2000/straper/chat-ws/pkg/storage/redis"
)

type Server struct {
	log        *zap.Logger
	httpServer *http.Server
	config     configs.Config
	store      mysql.Store
}

func NewServer(log *zap.Logger, config configs.Config, store mysql.Store) (*Server, error) {
	mr, err := setUpRoutes(log, store)
	if err != nil {
		log.Warn("Unable to set up route.")
		return &Server{}, err
	}

	srv := &http.Server{
		Handler:      getCORSHandler()(mr),
		Addr:         config.ServerAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server := &Server{
		log:        log,
		httpServer: srv,
		config:     config,
		store:      store,
	}
	return server, nil
}

func (s *Server) StartServer() {

	s.log.Info("Server is running.", zap.String("port", s.config.ServerAddress))

	err := s.httpServer.ListenAndServe()

	if err != nil {
		s.log.Warn("Unable to start server.")
		return
	}
}

func setUpRoutes(log *zap.Logger, store mysql.Store) (*mux.Router, error) {

	mr := mux.NewRouter().PathPrefix("/api/v1").Subrouter()

	rc := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	redisClient := rdb.NewRedisClient(rc, log)

	accountService := account.NewService(log, store)
	authService := auth.NewService(log, store)
	addingService := adding.NewService(log, store)
	chattingService := chatting.NewService(log, store, &redisClient)
	listingService := listing.NewService(log, store)
	deletingService := deleting.NewService(log, store)

	SetUpAuthRouter(mr, authService)
	SetUpAccountRouter(mr, accountService)
	SetUpWorkspaceRouter(mr, addingService, listingService, deletingService)
	SetUpChannelRouter(mr, addingService, chattingService)
	SetUpConnectionRouter(mr, log, chattingService)

	return mr, nil
}

func getCORSHandler() func(http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type", "Accept", "Authorization"}),
		handlers.AllowCredentials(),
	)
}
