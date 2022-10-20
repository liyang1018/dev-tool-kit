package jwt

import (
	"testing"
	"time"
)

var expireTime = 1 * time.Minute
var secret = "MyOwnSecret"

func TestGenerateToken(t *testing.T) {
	claims := map[string]interface{}{
		"username": "vito",
		"password": "123456",
	}
	token, err := GenerateToken(claims, secret)
	if err != nil {
		t.Errorf("GenerateToken error %s", err)
	}
	t.Logf("token: %s", token)
}

func TestCheckToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6IjEyMzQ1NiIsInVzZXJuYW1lIjoidml0byJ9.9FrzUgLatYXHe4Qj29JUdn-Db3jYM34-rfO7BUvzhqo"
	ok, err := CheckToken(tokenString, secret)
	if !ok || err != nil {
		t.Errorf("CheckToken error %s", err)
	}
	t.Logf("CheckToken ok %v", ok)
}
func TestParseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6IjEyMzQ1NiIsInVzZXJuYW1lIjoidml0byJ9.9FrzUgLatYXHe4Qj29JUdn-Db3jYM34-rfO7BUvzhqo"
	m, err := ParseToken(tokenString, secret)
	if err != nil {
		t.Errorf("ParseToken error %s", err)
	}
	for k, v := range m {
		t.Logf("key %s value %v", k, v)
	}
}
