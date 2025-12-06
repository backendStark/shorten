package auth

import (
	"fmt"
	"net/http"
	"shorten/configs"
	"shorten/pkg/req"
	"shorten/pkg/res"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}

		fmt.Println(body)

		resp := LoginResponse{
			Token: "LoginToken",
		}

		res.JSON(w, 200, resp)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}

		fmt.Println(body)

		resp := RegisterResponse{
			Token: "RegisterToken",
		}

		res.JSON(w, 200, resp)
	}
}
