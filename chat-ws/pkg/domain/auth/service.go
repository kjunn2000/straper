package auth

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"go.uber.org/zap"
)

type Repository interface {
	FindUserByUsername(username string) (*User, error)
}

type Service interface {
	Login(user User) (LoginResponseModel, error)
	RefreshToken(refreshToken string) (string, error)
}

type LoginResponseModel struct {
	AccessToken  string
	RefreshToken string
	User *User
}

type Claims struct {
	UserId   string
	Username string
	Role     string
	jwt.StandardClaims
}

type service struct {
	log *zap.Logger
	ar  Repository
}

func NewService(log *zap.Logger, ar Repository) *service {
	return &service{
		log: log,
		ar:  ar,
	}
}

func generateNewClaims(expiredAt time.Time, userId string, username string, role string) Claims {
	return Claims{
		UserId:   userId,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  jwt.Now(),
			ExpiresAt: jwt.At(expiredAt),
		},
	}
}

func (as *service) Login(user User) (LoginResponseModel, error) {

	u, err := as.ar.FindUserByUsername(user.Username)

	if err == sql.ErrNoRows {
		as.log.Info("User not found")
		return LoginResponseModel{}, errors.New("user.not.found")
	} else if u.Password != user.Password {
		as.log.Info("Invalid credential")
		return LoginResponseModel{}, errors.New("invalid.credential")
	}
	atc := generateNewClaims(time.Now().Add(time.Minute*10), u.UserId, u.Username, u.Role)
	rtc := generateNewClaims(time.Now().Add(time.Minute*45), u.UserId, u.Username, u.Role)
	att := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)
	rtt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)

	ats, aerr := att.SignedString(SecretKey)
	rts, rerr := rtt.SignedString(SecretKey)

	if aerr != nil || rerr != nil {
		as.log.Warn("Unable to sign jwt token.", zap.Error(aerr), zap.Error(rerr))
		return LoginResponseModel{}, err
	}
	return LoginResponseModel{
			AccessToken: "Bearer " + ats,
			RefreshToken: "Bearer " + rts,
			User: u,
		}, nil
}

func (as *service) RefreshToken(refreshToken string) (string, error) {

	claims, err := ExtractClaimsFromTokenStr(refreshToken)
	if err != nil {
		as.log.Info("Unable to extract claims from refresh token.")
		return "", err
	}

	atc := generateNewClaims(time.Now().Add(time.Minute*10), claims.UserId, claims.Username, claims.Role)

	att := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)

	ats, err := att.SignedString(SecretKey)

	if err != nil {
		as.log.Warn("Unable to sign jwt token.", zap.Error(err))
		return "", err
	}

	return "Bearer " + ats, nil
}

func ExtractClaimsFromTokenStr(tokenStr string) (Claims, error) {

	i := strings.Index(tokenStr, "Bearer ")
	if i == -1 || i != 0 {
		return Claims{}, errors.New("invalid.token.format")
	}
	tokenStr = tokenStr[7:]

	c := Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, &c,
		func(t *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})

	if err != nil || !token.Valid || c.ExpiresAt.Time.Before(time.Now()) {
		return Claims{}, err
	}
	return c, nil

}
