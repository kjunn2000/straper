package main

import (
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/kjunn2000/straper/chat-ws/configs"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/adding"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/chatting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/deleting"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/listing"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/handler"
	storage "github.com/kjunn2000/straper/chat-ws/pkg/storage/mysql"
	rdb "github.com/kjunn2000/straper/chat-ws/pkg/storage/redis"
	"go.uber.org/zap"
)

func setUpRoutes(log *zap.Logger, config configs.Config) (*mux.Router, error) {
	connStr := config.DBUser + ":" + config.DBPassword + config.DBSource
	db, err := sqlx.Connect(config.DBDriver, connStr)
	if err != nil {
		log.Warn("Unable to connect mysql database.", zap.Error(err))
		return nil, err
	}

	mr := mux.NewRouter().PathPrefix("/api/v1").Subrouter()

	store := storage.NewStore(db, log)

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

	handler.SetUpAuthRouter(mr, authService)
	handler.SetUpAccountRouter(mr, accountService)
	handler.SetUpWorkspaceRouter(mr, addingService, listingService, deletingService)
	handler.SetUpChannelRouter(mr, addingService, chattingService)
	handler.SetUpConnectionRouter(mr, log, chattingService)

	return mr, nil
}

func main() {

	log, _ := zap.NewDevelopment()
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Warn("Unable to load config.", zap.Error(err))
		return
	}
	mr, err := setUpRoutes(log, config)
	if err != nil {
		log.Warn("Unable to set up route.")
		return
	}

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type", "Accept", "Authorization"}),
		handlers.AllowCredentials(),
	)

	srv := &http.Server{
		Handler:      corsHandler(mr),
		Addr:         config.ServerAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("Server is running.", zap.String("port", config.ServerAddress))

	err = srv.ListenAndServe()

	if err != nil {
		log.Warn("Unable to start server.")
		return
	}
}
