package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

var (
	SecretKey = []byte("secretkey")
)

type Users map[string]*User

type User struct {
	UserId      string    `json:"user_id" db:"user_id"`
	Username    string    `json:"username" db:"username"`
	Password    string    `json:"password" db:"password"`
	Role        string    `json:"role" db:"role"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
}

type Claims struct {
	Username string
	jwt.StandardClaims
}
