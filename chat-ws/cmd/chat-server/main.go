package main

import (
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
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

func setUpRoutes(log *zap.Logger) (*mux.Router, error) {
	db, err := sqlx.Connect("mysql", "root:password@(localhost:3306)/straperdb?parseTime=true")
	if err != nil {
		log.Warn("Unable to connect mysql database.", zap.Error(err))
		return nil, err
	}

	mr := mux.NewRouter().PathPrefix("/api/v1").Subrouter()

	workspaceStore := storage.NewWorkspaceStore(db, log)
	channelStore := storage.NewChannelStore(db, log)
	userStore := storage.NewUserStore(log, db)

	rc := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	redisClient := rdb.NewRedisClient(rc, log)

	accountService := account.NewService(log, userStore)
	authService := auth.NewService(log, userStore)
	addingService := adding.NewService(log, workspaceStore, channelStore)
	chattingService := chatting.NewService(log, channelStore, &redisClient)
	listingService := listing.NewService(log, workspaceStore, channelStore)
	deletingService := deleting.NewService(log, workspaceStore, channelStore)

	handler.SetUpAuthRouter(mr, authService)
	handler.SetUpAccountRouter(mr, accountService)
	handler.SetUpWorkspaceRouter(mr, addingService, listingService, deletingService)
	handler.SetUpChannelRouter(mr, addingService, chattingService)
	handler.SetUpConnectionRouter(mr, log, chattingService)

	return mr, nil
}

func main() {

	log, _ := zap.NewDevelopment()
	mr, err := setUpRoutes(log)
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
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("Server is running.", zap.String("port", ":8080"))

	err = srv.ListenAndServe()

	if err != nil {
		log.Warn("Unable to start server.")
		return
	}
}
