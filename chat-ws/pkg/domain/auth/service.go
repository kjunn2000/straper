package auth

import (
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

	if err != nil || u.Password != user.Password {
		as.log.Info("User not found")
		return LoginResponseModel{}, err
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
	return LoginResponseModel{AccessToken: ats, RefreshToken: rts}, nil
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

	return ats, nil
}

func ExtractClaimsFromTokenStr(tokenStr string) (Claims, error) {

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
