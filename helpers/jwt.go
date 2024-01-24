package helpers

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const secretKey = "your-secret-key"

func GenerateJWTToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"exp": time.Now().Add(time.Hour * 168).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func ParseJWTToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	return token, err
}
