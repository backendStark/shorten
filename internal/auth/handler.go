package auth

import (
	"net/http"
	"shorten/configs"
	"shorten/pkg/jwt"
	"shorten/pkg/req"
	"shorten/pkg/res"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
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

		email, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			res.JSON(w, http.StatusUnauthorized, err.Error())
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(email)
		if err != nil {
			res.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		resp := LoginResponse{
			Token: token,
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

		email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)

		if err != nil {
			res.JSON(w, http.StatusBadRequest, err.Error())
			return
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(email)
		if err != nil {
			res.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		resp := RegisterResponse{
			Token: token,
		}

		res.JSON(w, http.StatusOK, resp)
	}
}
