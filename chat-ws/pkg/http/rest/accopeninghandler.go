package rest

import (
	"encoding/json"
	"net/http"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain"
	"go.uber.org/zap"
)

type AccOpeningHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type accOpeningHandler struct {
	log *zap.Logger
	us  domain.UserService
}

func NewAccOpeningHandler(log *zap.Logger, us domain.UserService) *accOpeningHandler {
	return &accOpeningHandler{
		log: log,
		us:  us,
	}
}

func (a *accOpeningHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)
	err := a.us.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
