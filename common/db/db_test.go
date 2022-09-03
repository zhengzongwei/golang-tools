package db

import (
	"testing"
)

func TestInitDb(t *testing.T) {
	db := InitDb("../conf/demo.yaml")
	CreateUser(db)
	//fmt.Printf("%#v", conf)
}
