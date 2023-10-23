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
