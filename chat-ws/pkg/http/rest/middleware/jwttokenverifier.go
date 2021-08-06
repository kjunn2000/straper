package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
)

func JwtTokenVerifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		v := r.Header.Get("Authorization")

		if v == "" {
			w.WriteHeader(http.StatusForbidden)
			rest.AddResponseToResponseWritter(w, nil, "access.token.not.found")
			return
		}

		i := strings.Index(v, "Bearer ")
		if i == -1 || i != 0 {
			w.WriteHeader(http.StatusForbidden)
			rest.AddResponseToResponseWritter(w, nil, "incorrect.token.format")
			return
		}
		v = v[7:]

		c := &auth.Claims{}
		token, err := jwt.ParseWithClaims(v, c,
			func(t *jwt.Token) (interface{}, error) {
				return auth.SecretKey, nil
			})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			rest.AddResponseToResponseWritter(w, nil, "Unauthorized request.")
			return
		}
		next.ServeHTTP(w, r)
	})
}
