package handler

import (
	"news-app-be23/internal/features/users"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func ToModelUsers(r SignUpRequest) users.User {
	return users.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}
