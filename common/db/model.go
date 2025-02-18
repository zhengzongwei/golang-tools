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

package db

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Name   string `gorm:"not null;unique"`
	Passwd string `gorm:"not null"`
	Status int64  `gorm:"defaule:0"`
}

type Log struct {
	gorm.Model
	Id       int64  `gorm:"not null;unique"`
	Level    string `gorm:"not null"`
	Message  string `gorm:"not null"`
	source   string `gorm:"not null"`
	userId   string `gorm:""`
	Ip       string `gorm:""`
	MetaData string `gorm:""`
}
