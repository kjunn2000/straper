package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kjunn2000/straper/chat-ws/configs"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(ctx context.Context, req LoginRequest) (LoginResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}

type LoginRequest struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
	User         LoginResponseUser
}

type LoginResponseUser struct {
	UserId   string `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email" validate:"email"`
	PhoneNo  string `json:"phone_no" db:"phone_no"`
	Role     string `json:"role" db:"role"`
}

type service struct {
	log        *zap.Logger
	ar         Repository
	tokenMaker Maker
	config     configs.Config
}

func NewService(log *zap.Logger, ar Repository, tokenMaker Maker, config configs.Config) *service {
	return &service{
		log:        log,
		ar:         ar,
		tokenMaker: tokenMaker,
		config:     config,
	}
}

func (as *service) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {

	u, err := as.ar.GetUserCredentialByUsername(ctx, req.Username)

	if err == sql.ErrNoRows {
		return LoginResponse{}, errors.New("user.not.found")
	} else if err = as.comparePassword(u.Password, req.Password); err != nil {
		return LoginResponse{}, errors.New("invalid.credential")
	} else if u.Status != "ACTIVE" {
		return LoginResponse{}, errors.New("invalid.account.status")
	}

	accessToken, err := as.tokenMaker.CreateToken(u.UserId, u.Username, as.config.AccessTokenDuration)
	if err != nil {
		return LoginResponse{}, err
	}

	refreshToken, err := as.tokenMaker.CreateToken(u.UserId, u.Username, as.config.RefreshTokenDuration)
	if err != nil {
		return LoginResponse{}, err
	}

	loginResponseUser := LoginResponseUser{
		UserId:   u.UserId,
		Username: u.Username,
		Email:    u.Email,
		PhoneNo:  u.PhoneNo,
		Role:     u.Role,
	}

	return LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         loginResponseUser,
	}, nil
}

func (as *service) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	payload, err := as.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		return "", err
	}

	accessToken, err := as.tokenMaker.CreateToken(payload.UserId, payload.Username, as.config.AccessTokenDuration)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (as *service) comparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
