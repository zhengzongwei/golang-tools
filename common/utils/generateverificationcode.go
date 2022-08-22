package utils

/*
 * @Title readconfig.go
 * @Description 生成验证码
 * @Author  zongweizheng 2022/8/22 16:45
 * @Update
 */

import (
	"math/rand"
	"time"
)

var (
	number        = []byte("1234567890")
	letters       = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numberLetters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

var typeMap = map[string][]byte{
	"number":        number,
	"letter":        letters,
	"letter_number": numberLetters,
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateVerificationCode(verificationCodeType string, verificationCodeLen int) string {
	b := make([]byte, verificationCodeLen)
	items := typeMap[verificationCodeType]

	for i := range b {
		b[i] = items[rand.Intn(len(number))]
	}
	return string(b)
}
