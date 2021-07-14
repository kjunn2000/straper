package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain"
)

func JwtTokenVerifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		v := r.Header.Get("Authorization")

		if v == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		c := &domain.Claims{}
		token, err := jwt.ParseWithClaims(v, c,
			func(t *jwt.Token) (interface{}, error) {
				return domain.SecretKey, nil
			})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
