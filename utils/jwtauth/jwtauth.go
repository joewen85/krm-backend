package jwtauth

import (
	"errors"
	"krm-backend/config"
	"krm-backend/utils/logs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJwt(username string) (string, error) {
	claims := JwtClaims{
		username,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(config.ExpireTime))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "joe",
			Subject:   "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(config.SigningKey))
	return ss, err
}

func ParseJwt(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SigningKey), nil
	})
	if err != nil {
		logs.Error(nil, "token解析错误")
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token claims验证失败")
	}
}
