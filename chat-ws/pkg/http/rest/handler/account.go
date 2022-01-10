package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest/middleware"
	zxcvbn "github.com/nbutton23/zxcvbn-go"
)

type ResetPasswordRequest struct {
	Email string `json:"email" db:"email" validate:"email"`
}

func (server *Server) SetUpAccountRouter(mr *mux.Router, as account.Service) {
	validate := validator.New()

	validate.RegisterValidation("password", validatePassword)
	ar := mr.PathPrefix("/account").Subrouter()
	pr := mr.PathPrefix("/protected/account").Subrouter()
	ar.HandleFunc("/create", server.Register(as, validate)).Methods("POST")
	pr.HandleFunc("/read/{user_id}", server.GetAccount(as)).Methods("GET")
	pr.HandleFunc("/update", server.UpdateAccount(as)).Methods("POST")
	pr.HandleFunc("/delete/{user_id}", server.DeleteAccount(as)).Methods("POST")
	ar.HandleFunc("/email/verify/{token_id}", server.ValidateVerifyEmailToken(as)).Methods("POST")
	ar.HandleFunc("/reset-password/create", server.ResetPasswordRequest(as)).Methods("POST")
	ar.HandleFunc("/password/update", server.UpdatePassword(as)).Methods("POST")
	pr.Use(middleware.TokenVerifier(server.tokenMaker))
}

func (server *Server) Register(as account.Service, validate *validator.Validate) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user account.CreateUserParam
		json.NewDecoder(r.Body).Decode(&user)
		err := validate.Struct(user)
		if err != nil {
			switch ExtractFieldFromValidationMsg(err.Error()) {
			case "Username":
				err = errors.New("invalid.username.format")
			case "Email":
				err = errors.New("invalid.email.format")
			case "PhoneNo":
				err = errors.New("invalid.phone.no.format")
			case "Password":
				err = errors.New("password.too.weak")
			}
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

func (server *Server) ValidateVerifyEmailToken(as account.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tokenId, ok := vars["token_id"]
		if !ok {
			rest.AddResponseToResponseWritter(w, nil, "invalid.token.id")
			return
		}
		err := as.ValidateVerifyEmailToken(r.Context(), tokenId)
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

func (server *Server) ResetPasswordRequest(as account.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var resetPasswordRequest ResetPasswordRequest
		json.NewDecoder(r.Body).Decode(&resetPasswordRequest)
		if err := as.ResetAccountPassword(r.Context(), resetPasswordRequest.Email); err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}

func (server *Server) UpdatePassword(as account.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var updatePasswordParam account.UpdatePasswordParam
		json.NewDecoder(r.Body).Decode(&updatePasswordParam)
		if err := as.UpdateAccountPassword(r.Context(), updatePasswordParam); err != nil {
			rest.AddResponseToResponseWritter(w, nil, err.Error())
			return
		}
		rest.AddResponseToResponseWritter(w, nil, "")
	}
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	score := zxcvbn.PasswordStrength(password, []string{})
	return score.Score >= 2
}
