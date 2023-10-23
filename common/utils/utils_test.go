// Package          utils
// @Title           utils_test.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/10/23 14:15

package utils

import (
	"fmt"
	"testing"
)

var pwd = "pwd123"

func TestEncryptedPassword(t *testing.T) {
	password := EncryptedPassword(pwd)
	fmt.Print(password)
}

func TestDecryptPassword(t *testing.T) {
	checkPWD := DecryptPassword(pwd)
	fmt.Println(checkPWD)
}
