package auth

import (
	"net/http"
	"net/http/httptest"
	"shorten/configs"
	"testing"
)

func TestNewAuthHandler(t *testing.T) {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	NewAuthHandler(router, AuthHandlerDeps{
		Config: conf,
	})

	tests := []string{
		"/auth/login",
		"/auth/register",
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			req := httptest.NewRequest("POST", tt, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != 200 {
				t.Errorf("%s return no 200 status code", tt)
			}
		})
	}
}
