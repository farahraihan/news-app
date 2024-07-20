package handler

import (
	"news-app-be23/internal/features/users"
	"news-app-be23/internal/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	srv users.Services
}

func NewUserController(s users.Services) users.Handler {
	return &UserController{
		srv: s,
	}
}

func (uc *UserController) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input SignUpRequest
		err := c.Bind(&input)
		if err != nil {
			c.Logger().Error("register parse error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}
		err = uc.srv.SignUp(ToModelUsers(input))
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "tidak valid") {
				errCode = 400
			}
			return c.JSON(500, helper.ResponseFormat(errCode, "server error", nil))
		}
		return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
		err := c.Bind(&input)
		if err != nil {
			c.Logger().Error("login parse error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}
		token, err := uc.srv.Login(input.Username, input.Password)
		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "tidak ditemukan") {
				errCode = 400
			}
			return c.JSON(500, helper.ResponseFormat(errCode, "server error", nil))
		}
		return c.JSON(200, helper.ResponseFormat(200, "success login", ToLoginReponse(token)))
	}
}
