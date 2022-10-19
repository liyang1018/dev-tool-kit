package jwt

import (
	"testing"
	"time"
)

var expireTime = 1 * time.Minute
var secret = "MyOwnSecret"

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(nil, secret)
	if err != nil {
		t.Errorf("error %s", err)
	}
	if token == "" {
		t.Errorf("token is empty")
	}
}
