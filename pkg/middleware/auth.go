package middleware

import (
	"context"
	"net/http"
	"shorten/configs"
	"shorten/pkg/jwt"
	"shorten/pkg/res"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func writeUnauthed(w http.ResponseWriter) {
	res.JSON(w, http.StatusUnauthorized, "Unauthorized")
}

func IsAuth(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer ") {
			writeUnauthed(w)
			return
		}
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)

		if !isValid {
			writeUnauthed(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, ContextEmailKey, data.Email)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
