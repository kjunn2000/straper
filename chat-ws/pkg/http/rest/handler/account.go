package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
)

func SetUpAccountRouter(mr *mux.Router, as account.Service) {
	ar := mr.PathPrefix("/account").Subrouter()
	ar.HandleFunc("/opening", Register(as)).Methods("POST")
}

func Register(as account.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		validate := validator.New()
		var user account.User
		json.NewDecoder(r.Body).Decode(&user)
		err := validate.Struct(user)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		err = as.Register(user)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}
