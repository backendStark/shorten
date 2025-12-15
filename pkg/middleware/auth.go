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

func IsAuth(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)

		if !isValid {
			res.JSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, ContextEmailKey, data.Email)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
