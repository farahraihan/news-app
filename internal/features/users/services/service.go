package services

import (
	"errors"
	"news-app-be23/internal/features/users"
	"news-app-be23/internal/utils"
)

type UserServices struct {
	qry users.Query
	pu  utils.PasswordUtilityInterface
	jt  utils.JwtUtilityInterface
}

func NewUserService(q users.Query, p utils.PasswordUtilityInterface, j utils.JwtUtilityInterface) users.Services {
	return &UserServices{
		qry: q,
		pu:  p,
		jt:  j,
	}
}

func (us *UserServices) SignUp(newData users.User) error {
	processPw, err := us.pu.GeneratePassword(newData.Password)
	if err != nil {
		return errors.New("input data tidak valid, data tidak bisa diproses")
	}
	newData.Password = string(processPw)
	err = us.qry.SignUp(newData)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat mengolah data")
	}
	return nil
}

func (us *UserServices) Login(username string, password string) (string, error) {
	result, err := us.qry.Login(username)
	if err != nil {
		return "", errors.New("terjadi kesalahan pada server saat login")
	}
	err = us.pu.CheckPassword([]byte(password), []byte(result.Password))
	if err != nil {
		return "", errors.New("input data tidak valid, data tidak bisa diproses")
	}
	token, err := us.jt.GenerateToken(result.ID)
	if err != nil {
		return "", errors.New("terjadi kesalahan pada saat generate token")
	}
	return token, nil
}
