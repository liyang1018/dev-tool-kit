package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(data interface{}, secret string) (string, error) {
	if secret == "" {
		return "", errors.New("secret is empty")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	mapData, ok := data.(map[string]interface{})
	if !ok {
		return "", errors.New("input data isn't map type")
	}
	for k, v := range mapData {
		claims[k] = v
	}
	if err := claims.Valid(); err != nil {
		return "", err
	}
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
