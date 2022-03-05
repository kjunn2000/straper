package middleware

import (
	"fmt"
	"net/http"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain/auth"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/dblog"
)

func UpdateLastSeen(sl dblog.StatusLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			payloadVal := r.Context().Value(TokenPayload{})
			if payloadVal == nil {
				next.ServeHTTP(w, r)
			}
			payload, ok := payloadVal.(*auth.Payload)
			if !ok || payload == nil {
				next.ServeHTTP(w, r)
			}
			fmt.Println("update user access info ...." + payload.Username)
			sl.UpdateLastSeen(r.Context(), payload.CredentialId)
			next.ServeHTTP(w, r)
		})
	}
}
