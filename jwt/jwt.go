package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

//
// GenerateToken
//  @Description: generate jwt token
//  @param data : user claims
//  @param secret : your own secret
//  @return string
//  @return error
//
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

//
// CheckToken
//  @Description: check jwt token
//  @param tokenString : jwt to be checked
//  @param secret : your own secret
//  @return bool
//  @return error
//
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

//
// ParseToken
//  @Description:
//  @param token : jwt to be parsed
//  @param secret : your own secret
//  @return jwt.MapClaims : user claims
//  @return error
//
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
