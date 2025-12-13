package middleware

import (
	"net/http"
)

func IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// authedHeader := r.Header.Get("Authorization")
		// token := strings.TrimPrefix(authedHeader, "Bearer ")

		next.ServeHTTP(w, r)
	})
}
