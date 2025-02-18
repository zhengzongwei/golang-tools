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
	numberLetters = []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
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
		b[i] = items[rand.Intn(len(items))]
	}
	return string(b)
}
