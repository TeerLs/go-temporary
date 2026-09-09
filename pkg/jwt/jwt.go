package jwt

import (
	"errors"
	"time"

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
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) Verify(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.SecretKey), nil
	})
}

func (j *JWT) IsValid(tokenString string) (string, error) {
	token, err := j.Verify(tokenString)
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if phone, ok := claims["phone"].(string); ok {
			return phone, nil
		}
	}
	return "", errors.New("invalid token")
}