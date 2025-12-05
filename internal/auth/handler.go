package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shorten/configs"
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
		fmt.Println(handler.Config.Auth.Token)
		fmt.Println("You send GET response to /auth/login route")



		// fmt.Println(payload)

		resp := LoginResponse{
			Token: "123",
		}

		res.JSON(w, 200, resp)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("You send GET response to /auth/register route")
	}
}

