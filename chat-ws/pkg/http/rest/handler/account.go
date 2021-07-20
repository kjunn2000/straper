package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
)

func NewAccountRouter(us account.Service) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/opening", Register(us)).Methods("POST")
	return router
}

func Register(us account.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user account.User
		json.NewDecoder(r.Body).Decode(&user)
		err := us.Register(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}
