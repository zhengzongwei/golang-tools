/*
 * Copyright (c)2025. zhengzongwei
 * golang-tools is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *         http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 */

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
