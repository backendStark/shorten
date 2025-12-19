package auth_test

import (
	"shorten/internal/auth"
	"shorten/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func (repo MockUserRepository) Create(user *user.User) (*user.User, error) {
	return user, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "a@ya.ru"
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(initialEmail, "1", "Вася")

	if err != nil {
		t.Fatal(err)
	}

	if email != initialEmail {
		t.Fatalf("Expected email %s, got %s", initialEmail, email)
	}
}
