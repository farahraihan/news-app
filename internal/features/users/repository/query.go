package repository

import (
	"news-app-be23/internal/features/users"

	"gorm.io/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(connection *gorm.DB) users.Query {
	return &UserModel{
		db: connection,
	}
}

func (um *UserModel) Login(username string) (users.User, error) {
	var result User
	err := um.db.Where("username = ?", username).First(&result).Error
	if err != nil {
		return users.User{}, err
	}
	return result.toUserEntity(), nil
}

func (um *UserModel) SignUp(newUser users.User) error {
	err := um.db.Create(&newUser).Error
	return err
}
