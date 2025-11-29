package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	handler := &AuthHandler{}
	req := httptest.NewRequest(http.MethodGet, "/auth/login", nil)
	w := httptest.NewRecorder()

	handler.Login().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", w.Code)
	}
}

func TestRegister(t *testing.T) {
	handler := &AuthHandler{}
	req := httptest.NewRequest(http.MethodGet, "/auth/register", nil)
	w := httptest.NewRecorder()

	handler.Register().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", w.Code)
	}
}

func TestNewAuthHandler(t *testing.T) {
	mux := http.NewServeMux()
	NewAuthHandler(mux)

	tests := []struct {
		path string
	}{
		{"/auth/login"},
		{"/auth/register"},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(http.MethodGet, tt.path, nil)
		w := httptest.NewRecorder()

		// mux.ServeHTTP will route the request to the registered handler
		mux.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("path %s: expected status OK; got %v", tt.path, w.Code)
		}
	}
}
