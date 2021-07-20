package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
)

func NewAuthRouter(as auth.Service) *mux.Router {
	mr := mux.NewRouter()
	mr.HandleFunc("/login", Login(as)).Methods("POST")
	mr.HandleFunc("/refresh-token", RefreshToken(as)).Methods("POST")
	return mr
}

func Login(as auth.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := auth.User{}
		json.NewDecoder(r.Body).Decode(&user)
		if user.Username == "" || user.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(w, nil, "Invalid credential.")
			return
		}
		loginResponse, err := as.Login(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		accessTokenCookie := &http.Cookie{
			Name:    "access_token",
			Value:   loginResponse.AccessToken,
			Expires: time.Now().Add(time.Minute * 10),
		}
		refreshTokenCookie := &http.Cookie{
			Name:    "refresh_token",
			Value:   loginResponse.RefreshToken,
			Expires: time.Now().Add(time.Minute * 45),
		}
		http.SetCookie(w, accessTokenCookie)
		http.SetCookie(w, refreshTokenCookie)
		rest.AddResponseToResponseWritter(w, loginResponse, "")
	}
}

func RefreshToken(as auth.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rt, err := r.Cookie("refresh_token")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		v := rt.Value

		if v == "" {
			w.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		ats, err := as.RefreshToken(v)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}

		accessTokenCookie := &http.Cookie{
			Name:    "access_token",
			Value:   ats,
			Expires: time.Now().Add(time.Minute * 10),
		}

		http.SetCookie(w, accessTokenCookie)
		rest.AddResponseToResponseWritter(w, accessTokenCookie, "")
	}
}
