package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
)

func SetUpAuthRouter(mr *mux.Router, as auth.Service) {
	ar := mr.PathPrefix("/auth").Subrouter()
	ar.HandleFunc("/login", Login(as)).Methods("POST")
	ar.HandleFunc("/refresh-token", RefreshToken(as)).Methods("POST")
}

type LoginResponseModal struct {
	AccessToken string     `json:"access_token"`
	Identity    *auth.User `json:"identity"`
}

func Login(as auth.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := auth.User{}
		json.NewDecoder(r.Body).Decode(&user)
		if user.Username == "" || user.Password == "" {
			rest.AddResponseToResponseWritter(w, nil, "Invalid credential.")
			return
		}
		loginResponse, err := as.Login(user)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		refreshTokenCookie := &http.Cookie{
			Name:     "refresh_token",
			Value:    loginResponse.RefreshToken,
			Expires:  time.Now().Add(time.Minute * 45),
			HttpOnly: true,
			Secure:   false,
			Path:     "/",
		}
		http.SetCookie(w, refreshTokenCookie)
		loginResponseModal := LoginResponseModal{
			AccessToken: loginResponse.AccessToken,
			Identity:    loginResponse.User,
		}
		rest.AddResponseToResponseWritter(w, loginResponseModal, "")
	}
}

func RefreshToken(as auth.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rt, err := r.Cookie("refresh_token")
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		v := rt.Value

		if v == "" {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		ats, err := as.RefreshToken(v)
		if err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}

		rest.AddResponseToResponseWritter(w, ats, "")
	}
}
