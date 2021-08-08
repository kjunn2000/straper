package mysql

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/kjunn2000/straper/chat-ws/configs"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

var testDb *sqlx.DB
var store *Store
var log *zap.Logger

func TestMain(m *testing.M) {
	log, _ = zap.NewDevelopment()
	config, err := configs.LoadConfig("../../../")
	if err != nil {
		log.Fatal("Unable to load config")
	}
	connStr := config.DBUser + ":" + config.DBPassword + config.DBSource
	testDb, err = sqlx.Connect(config.DBDriver, connStr)
	if err != nil {
		log.Fatal("Failed to open conn.", zap.Error(err))
	}
	store = NewStore(testDb, log)
	os.Exit(m.Run())
}
