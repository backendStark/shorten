package user

import "shorten/pkg/db"

type UserRepository struct {
	databsse *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.databsse.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.databsse.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
