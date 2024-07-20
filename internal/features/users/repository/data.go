package repository

import (
	"news-app-be23/internal/features/articles"
	"news-app-be23/internal/features/comments"
	"news-app-be23/internal/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string             `json:"username"`
	Password string             `json:"password"`
	Email    string             `json:"email"`
	Articles []articles.Article `gorm:"foreignKey:UserID"`
	Comments []comments.Comment `gorm:"foreignKey:UserID"`
}

func (u *User) toUserEntity() users.User {
	return users.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}
}

func toUserData(input users.User) User {
	return User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}
}
