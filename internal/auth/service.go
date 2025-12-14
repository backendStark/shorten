package auth

import (
	"errors"
	"shorten/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(UserRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository}
}

func (service *AuthService) Register(email, password, name string) (*user.User, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return nil, errors.New(ErrUserExists)
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &user.User{Email: email, Password: string(hashPassword), Name: name}

	_, err = service.UserRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
