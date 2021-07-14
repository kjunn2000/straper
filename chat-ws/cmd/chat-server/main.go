package main

import (
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	storage "github.com/kjunn2000/straper/chat-ws/pkg/storage/mysql"
	"go.uber.org/zap"
)

func setUpRoutes(log *zap.Logger) (*mux.Router, error) {
	db, err := sqlx.Connect("mysql", "root:password@(localhost:3306)/straperdb?parseTime=true")
	if err != nil {
		log.Warn("Unable to connect mysql database.", zap.Error(err))
		return nil, err
	}

	mr := mux.NewRouter()

	// wstore := storage.NewWorkspaceStore(db, log)
	// ws := domain.NewWorkspaceService(wstore, log)
	// wh := rest.NewWorkspaceHandler(ws, log)
	// wr := mr.PathPrefix("/api/v1/workspace").Subrouter()
	// wr.HandleFunc("/{name}", wh.CreateWorkspace).Methods("POST")
	// wr.HandleFunc("/{id}", wh.DeleteWorkspace).Methods("DELETE")
	// wr.HandleFunc("/{id}", wh.GetWorkspace).Methods("GET")
	// wr.HandleFunc("", wh.EditWorkspace).Methods("PUT")

	ustore := storage.NewUserStore(log, db)
	as := domain.NewAuthService(log, ustore)
	ah := rest.NewAuthHandler(log, as)
	ar := mr.PathPrefix("/api/auth").Subrouter()
	ar.HandleFunc("/login", ah.LoginHandler)
	ar.HandleFunc("/refresh-token", ah.RefreshTokenHandler)

	us := domain.NewUserService(log, ustore)
	aoh := rest.NewAccOpeningHandler(log, us)
	mr.HandleFunc("/api/account/opening", aoh.Register)

	// mr.Use(middleware.JwtTokenVerifier)

	return mr, nil
}

func main() {

	log, _ := zap.NewDevelopment()
	mr, err := setUpRoutes(log)
	if err != nil {
		log.Warn("Unable to set up route.")
		return
	}

	srv := &http.Server{
		Handler:      mr,
		Addr:         "127.0.0.1:9090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("Server is running.", zap.String("port", ":9090"))

	err = srv.ListenAndServe()

	if err != nil {
		log.Warn("Unable to start server.")
		return
	}
}
