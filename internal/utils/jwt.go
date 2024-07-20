package utils

import (
	"os"
	"time"

	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JwtUtilityInterface interface {
	GenerateToken(userID uint) (string, error)
	DecodeToken(token *jwt.Token) uint
}

type jwtUtility struct{}

func NewJwtUtility() JwtUtilityInterface {
	return &jwtUtility{}
}

func (jt *jwtUtility) GenerateToken(userID uint) (string, error) {
	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		return "", fmt.Errorf("JWT secret key not found in environment variables")
	}
	var claims = jwt.MapClaims{}
	claims["id"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Second * 60).Unix()

	var process = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := process.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return result, nil
}

func (jt *jwtUtility) DecodeToken(token *jwt.Token) uint {
	var claims = token.Claims.(jwt.MapClaims)
	var userID = claims["id"].(float64)
	return uint(userID)
}
