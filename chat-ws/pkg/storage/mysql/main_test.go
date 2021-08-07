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
	config, err := configs.LoadConfig("../../../")
	if err != nil {
		log.Fatal("Unable to load config")
	}
	testDb, err = sqlx.Connect(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Failed to open conn.")
	}
	log, _ = zap.NewDevelopment()
	store = NewStore(testDb, log)
	os.Exit(m.Run())
}
