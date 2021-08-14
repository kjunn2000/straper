package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
)

func (server *Server) SetUpAccountRouter(mr *mux.Router, as account.Service) {
	ar := mr.PathPrefix("/account").Subrouter()
	ar.HandleFunc("/create", server.Register(as)).Methods("POST")
	ar.HandleFunc("/read/{user_id}", server.GetAccount(as)).Methods("GET")
	ar.HandleFunc("/update", server.UpdateAccount(as)).Methods("POST")
	ar.HandleFunc("/delete/{user_id}", server.DeleteAccount(as)).Methods("POST")
}

func (server *Server) Register(as account.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		validate := validator.New()
		var user account.CreateUserParam
		json.NewDecoder(r.Body).Decode(&user)
		err := validate.Struct(user)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		err = as.Register(r.Context(), user)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}

func (server *Server) GetAccount(as account.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId, ok := vars["user_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "invalid.user.id")
			return
		}
		user, err := as.GetUserByUserId(r.Context(), userId)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, user, "")
	}
}

func (server *Server) UpdateAccount(as account.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		validate := validator.New()
		var user account.UpdateUserParam
		json.NewDecoder(r.Body).Decode(&user)
		err := validate.Struct(user)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		err = as.UpdateUser(r.Context(), user)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}

func (server *Server) DeleteAccount(as account.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId, ok := vars["user_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "invalid.user.id")
			return
		}
		err := as.DeleteUser(r.Context(), userId)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}
