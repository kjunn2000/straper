package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/admin"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
)

func (server *Server) SetUpManageUserRouter(mr *mux.Router, as admin.Service) {
	r := mr.PathPrefix("/admin/protected/user").Subrouter()
	r.HandleFunc("/read/{user_id}", server.GetUser(as)).Methods("GET")
	r.HandleFunc("/list", server.GetPaginationUsers(as)).Methods("GET")
	r.HandleFunc("/update", server.UpdateUser(as)).Methods("POST")
	r.HandleFunc("/delete/{user_id}", server.DeleteUser(as)).Methods("POST")
	r.Use(middleware.TokenVerifier(server.tokenMaker))
	r.Use(middleware.UpdateLastSeen(server.statusLogger))
}

func (server *Server) GetUser(as admin.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId, ok := vars["user_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "invalid.user.id")
			return
		}
		user, err := as.GetUser(r.Context(), userId)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, user, "")
	}
}

func (server *Server) GetPaginationUsers(as admin.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 64)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, "invalid.limit")
			return
		}
		cursor := r.URL.Query().Get("cursor")
		isNext, err := strconv.ParseBool(r.URL.Query().Get("isNext"))
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, "invalid.is.next.attribute")
			return
		}
		searchStr := r.URL.Query().Get("searchStr")
		param := admin.PaginationUsersParam{
			Limit:     limit,
			Cursor:    cursor,
			IsNext:    isNext,
			SearchStr: searchStr,
		}
		users, err := as.GetPaginationUsers(r.Context(), param)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, users, "")
	}
}

func (server *Server) UpdateUser(as admin.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user admin.UpdateUserParam
		json.NewDecoder(r.Body).Decode(&user)
		err := as.UpdateUser(r.Context(), user)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}

func (server *Server) DeleteUser(as admin.Service) func(http.ResponseWriter, *http.Request) {
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
