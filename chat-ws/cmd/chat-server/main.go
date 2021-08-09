package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kjunn2000/straper/chat-ws/configs"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/handler"
	"github.com/kjunn2000/straper/chat-ws/pkg/storage/mysql"

	_ "github.com/golang/mock/mockgen/model"
	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewDevelopment()

	config, err := configs.LoadConfig(".")

	if err != nil {
		log.Warn("Unable to load config.", zap.Error(err))
	}
	store, err := getSQLStore(log, config)
	if err != nil {
		return
	}
	server, err := handler.NewServer(log, config, store)
	if err != nil {
		return
	}
	server.StartServer()
}

func getSQLStore(log *zap.Logger, config configs.Config) (mysql.Store, error) {
	connStr := config.DBUser + ":" + config.DBPassword + config.DBSource
	db, err := sqlx.Connect(config.DBDriver, connStr)
	if err != nil {
		log.Warn("Unable to connect mysql database.", zap.Error(err))
		return &mysql.SQLStore{}, err
	}
	store := mysql.NewStore(db, log)
	return store, nil
}
