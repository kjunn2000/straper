package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"go.uber.org/zap"
)

type AuthRepository interface {
	FindUserByUsername(username string) (*User, error)
}

type AuthService interface {
	Login(user User) (LoginResponseModel, error)
	RefreshToken(refreshToken string) (string, error)
}

type LoginResponseModel struct {
	AccessToken  string
	RefreshToken string
}

type authService struct {
	log *zap.Logger
	ar  AuthRepository
}

func NewAuthService(log *zap.Logger, ar AuthRepository) *authService {
	return &authService{
		log: log,
		ar:  ar,
	}
}

func (as *authService) Login(user User) (LoginResponseModel, error) {

	u, err := as.ar.FindUserByUsername(user.Username)

	if err != nil || u.Password != user.Password{
		as.log.Info("User not found")
		return LoginResponseModel{}, err
	}
	atc := Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  jwt.Now(),
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 10)),
		},
	}
	rtc := Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  jwt.Now(),
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 45)),
		},
	}
	att := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)
	rtt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)

	ats, aerr := att.SignedString(SecretKey)
	rts, rerr := rtt.SignedString(SecretKey)

	if aerr != nil || rerr != nil {
		as.log.Warn("Unable to sign jwt token.", zap.Error(aerr), zap.Error(rerr))
		return LoginResponseModel{}, err
	}
	return LoginResponseModel{AccessToken: ats, RefreshToken: rts}, nil
}

func (as *authService) RefreshToken(refreshToken string) (string, error) {
	c := &Claims{}
	token, err := jwt.ParseWithClaims(refreshToken, c,
		func(t *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})

	if err != nil || !token.Valid || c.ExpiresAt.Time.Before(time.Now()) {
		as.log.Info("Refresh token is not valid.", zap.Error(err))
		return "", err
	}

	atc := Claims{
		Username: c.Username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  jwt.Now(),
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 10)),
		},
	}

	att := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)

	ats, err := att.SignedString(SecretKey)

	if err != nil {
		as.log.Warn("Unable to sign jwt token.", zap.Error(err))
		return "", err
	}

	return ats, nil
}
