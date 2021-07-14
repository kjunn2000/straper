package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain"
	"go.uber.org/zap"
)

type AuthHandler interface {
	LoginHandler(w http.ResponseWriter, r *http.Request)
	RefreshTokenHandler(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	log *zap.Logger
	as  domain.AuthService
}

func NewAuthHandler(log *zap.Logger, as domain.AuthService) *authHandler {
	return &authHandler{
		log: log,
		as:  as,
	}
}

func (a *authHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	user := domain.User{}
	json.NewDecoder(r.Body).Decode(&user)
	if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	loginResponse, err := a.as.Login(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

}

func (a *authHandler) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	rt, err := r.Cookie("refresh_token")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	v := rt.Value

	if v == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ats, err := a.as.RefreshToken(v)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accessTokenCookie := &http.Cookie{
		Name:    "access_token",
		Value:   ats,
		Expires: time.Now().Add(time.Minute * 10),
	}

	http.SetCookie(w, accessTokenCookie)
}
