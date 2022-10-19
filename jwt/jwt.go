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

func CheckToken(tokenString string, secret string) (bool, error) {
	if secret == "" {
		return false, errors.New("secret is empty")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(secret), nil
	})
	if token == nil {
		return false, errors.New("token is illegal")
	}
	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				return false, errors.New("token format is incorrect")
			} else if v.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return false, errors.New("token is expired")
			} else {
				return false, errors.New("token is error")
			}
		}
	}
	return true, nil
}

func ParseToken(token string, secret string) (jwt.MapClaims, error) {
	if secret == "" {
		return nil, errors.New("secret is empty")
	}
	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenObj.Claims.(jwt.MapClaims); ok && tokenObj.Valid {
		return claims, nil
	}
	return nil, errors.New("parse token error")
}
