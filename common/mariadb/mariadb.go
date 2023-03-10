package mariadb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 定义 数据库连接信息

type DbConn struct {
	Host     string
	UserName string
	Password string
	Port     string
	DbName   string
}

func init_db() {
	dbInfo := DbConn{}
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbInfo.UserName, dbInfo.Password, dbInfo.Host, dbInfo.Port, dbInfo.DbName))

	if err != nil {
		panic(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
}
