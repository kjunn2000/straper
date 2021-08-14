package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/http/rest"
)

var tokenType string = "bearer"

type TokenPayload struct{}

func TokenVerifier(tokenMaker auth.Maker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			v := r.Header.Get("Authorization")

			if v == "" {
				w.WriteHeader(http.StatusForbidden)
				rest.AddResponseToResponseWritter(w, nil, "access.token.not.found")
				return
			}

			fields := strings.Fields(v)
			if len(fields) != 2 || strings.ToLower(fields[0]) != tokenType {
				w.WriteHeader(http.StatusForbidden)
				rest.AddResponseToResponseWritter(w, nil, "invalid.access.token")
				return
			}

			payload, err := tokenMaker.VerifyToken(fields[1])
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				rest.AddResponseToResponseWritter(w, nil, "invalid.access.token")
				return
			}
			r = r.Clone(context.WithValue(r.Context(), TokenPayload{}, payload))

			next.ServeHTTP(w, r)
		})
	}
}
