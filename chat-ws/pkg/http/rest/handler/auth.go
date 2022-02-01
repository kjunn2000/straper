package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
)

func (server *Server) SetUpAuthRouter(mr *mux.Router, as auth.Service) {
	ar := mr.PathPrefix("/auth").Subrouter()
	ar.HandleFunc("/login", server.Login(as)).Methods("POST")
	ar.HandleFunc("/refresh-token", server.RefreshToken(as)).Methods("POST")
}

type LoginResponseModal struct {
	AccessToken string                 `json:"access_token"`
	User        auth.LoginResponseUser `json:"user"`
}

func (server *Server) Login(as auth.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := auth.LoginRequest{}
		json.NewDecoder(r.Body).Decode(&user)
		if user.Username == "" || len(user.Username) < 4 || user.Password == "" {
			rest.AddResponseToResponseWritter(w, nil, "invalid.credential")
			return
		}
		loginResponse, err := as.Login(r.Context(), user)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		refreshTokenCookie := &http.Cookie{
			Name:     "refresh_token",
			Value:    loginResponse.RefreshToken,
			Expires:  time.Now().Add(server.config.RefreshTokenDuration),
			HttpOnly: true,
			Secure:   false,
			Path:     "/",
		}
		http.SetCookie(w, refreshTokenCookie)
		loginResponseModal := LoginResponseModal{
			AccessToken: loginResponse.AccessToken,
			User:        loginResponse.User,
		}
		rest.AddResponseToResponseWritter(w, loginResponseModal, "")
	}
}

func (server *Server) RefreshToken(as auth.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rt, err := r.Cookie("refresh_token")
		if err != nil {
			fmt.Println("not refresh token")
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		v := rt.Value

		if v == "" {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		ats, err := as.RefreshToken(r.Context(), v)
		if err != nil {
			fmt.Println("fail to refresh token")
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}

		rest.AddResponseToResponseWritter(w, ats, "")
	}
}
