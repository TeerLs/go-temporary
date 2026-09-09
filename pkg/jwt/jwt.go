package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	SecretKey string
}

func NewJWT(secretKey string) *JWT {
	return &JWT{
		SecretKey: secretKey,
	}
}

func (j *JWT) Create(phone string) (string, error) {
	claims := jwt.MapClaims{
		"phone": phone,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}