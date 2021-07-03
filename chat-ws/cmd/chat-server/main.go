package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	storage "github.com/kjunn2000/straper/chat-ws/pkg/storage/mysql"
	"go.uber.org/zap"
)

func setUpRoutes(log *zap.Logger) error {
	db, err := sqlx.Connect("mysql", "root:password@(localhost:3306)/straperdb")
	if err != nil {
		log.Warn("Unable to connect mysql database.", zap.Error(err))
		return err
	}
	wstore := storage.NewWorkspaceStore(db, log)
	ws := domain.NewWorkspaceService(wstore, log)
	wh := rest.NewWorkspaceHandler(ws, log)

	r := mux.NewRouter()
	wr := r.PathPrefix("/workspace").Subrouter()
	wr.HandleFunc("/{name}", wh.CreateWorkspace).Methods("POST")
	wr.HandleFunc("/{id}", wh.DeleteWorkspace).Methods("DELETE")
	wr.HandleFunc("/{id}", wh.GetWorkspace).Methods("GET")
	wr.HandleFunc("", wh.EditWorkspace).Methods("PUT")
	http.Handle("/", r)
	return nil
}

func main() {

	log, _ := zap.NewDevelopment()
	err := setUpRoutes(log)
	if err != nil {
		log.Warn("Unable to set up route.")
		return
	}
	port := ":9090"
	log.Info("Server is running.", zap.String("port", port))
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Warn("Unable to start server.")
		return
	}
}
