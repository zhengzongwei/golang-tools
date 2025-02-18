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
// @Title           utils.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/10/23 14:09

package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func EncryptedPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	encodePWD := string(hash)
	fmt.Println(encodePWD)
	return encodePWD
}

func DecryptPassword(password string) bool {
	// TODO 从数据库中获取已加密的密码
	encodePWD := "$2a$10$SMYNM68w6rl4gWqVzGksuufEVHGTziLg6UiECB608KvNdSxpGjAO2"

	err := bcrypt.CompareHashAndPassword([]byte(encodePWD), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
