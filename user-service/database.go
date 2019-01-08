package main

import (
	"fmt"
	"github.com/xormplus/xorm"
	"os"
)

func CreateConnection() (*xorm.Engine, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PWD")
	dbName := os.Getenv("DB_NAME")
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pwd,
		host,
		dbName)
	return xorm.NewEngine("mysql", dns)
}
