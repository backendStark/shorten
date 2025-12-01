package auth

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
