package string

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

//
// Encrypt
//  @Description: encrypt by md5
//  @param input
//  @return output
//
func Encrypt(input []byte) (output string) {
	b := input
	h := md5.New()
	return hex.EncodeToString(h.Sum(b))
}

//
// GenerateRandNumStr
//  @Description: generate random number string
//  @param n length of string
//  @return string
//
func GenerateRandNumStr(n int) string {
	var letterRunes = []rune("0123456789")
	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//
// GenerateRandLetterStr
//  @Description: generate random letter string
//  @param n length of string
//  @return string
//
func GenerateRandLetterStr(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
