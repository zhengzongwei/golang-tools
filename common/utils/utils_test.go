// Package          utils
// @Title           utils_test.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/10/23 14:15

package utils

import (
	"fmt"
	"testing"
)

func TestEncryptedPassword(t *testing.T) {
	password := EncryptedPassword("pwd123")
	fmt.Print(password)
}

func TestDecryptPassword(t *testing.T) {
	password := "pwd12"
	checkPWD := DecryptPassword(password)
	fmt.Println(checkPWD)
}
