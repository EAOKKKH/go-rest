package utils

import (
	"go-rest/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(login string, cfg *config.Config) (string, error) {
	claims := &JwtClaims{
		Login: login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(cfg.Server.JwtExpire))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(cfg.Server.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
