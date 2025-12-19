package jwt_test

import (
	"shorten/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "a@ya.ru"
	jwtService := jwt.NewJWT("secret")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})

	if err != nil {
		t.Fatal(err)
	}

	if token == "" {
		t.Fatal("Token is empty")
	}

	isValid, data := jwtService.Parse(token)

	if !isValid {
		t.Fatal("Token is invalid")
	}

	if data.Email != email {
		t.Fatalf("Expected email %s, got %s", email, data.Email)
	}
}
