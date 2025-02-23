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

import (
	"fmt"
	"golang-tools/common/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDb(confPath string) (db *gorm.DB) {
	conf := utils.GetYamlConfig(confPath)
	fmt.Printf("%#v", conf.MYSQL)
	//var DSN = "#{user}:#{password}@tcp({#host}:#{port})/{name}?charset=utf8mb4&parseTime=True&loc=Local"
	var DSN = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	targetDsn := fmt.Sprintf(DSN, conf.MYSQL.User, conf.MYSQL.Password, conf.MYSQL.Host, conf.MYSQL.Port, conf.MYSQL.Name)
	//fmt.Println(targetDsn)
	db, err := gorm.Open(mysql.Open(targetDsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	InitModel(db)

	return db
	//defer db.Close()
	//return db
	//defer db.Close()
}

func InitModel(db *gorm.DB) {
	//defer db.Close()
	err := db.AutoMigrate(&UserInfo{})
	if err != nil {
		log.Printf("ERROR!,Migrate Model Failed %s ", err)
	}
}

func CreateUser(db *gorm.DB) {
	user := UserInfo{Name: "zhengzongwei", Passwd: "zhengzongwei", Status: 1}
	user.Passwd = utils.EncryptedPassword(user.Passwd)

	db.Create(&user)
}

//func init() {
//	InitDb("../conf/demo.yaml")
//}
