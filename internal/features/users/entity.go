package users

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       uint
	Username string
	Password string
	Email    string
}

type Handler interface {
	SignUp() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type Services interface {
	SignUp(newUser User) error
	Login(username string, password string) (string, error)
}

type Query interface {
	SignUp(newUser User) error
	Login(username string) (User, error)
}
