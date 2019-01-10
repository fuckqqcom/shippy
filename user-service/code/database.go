package code

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func CreateConnection() (*gorm.DB, error) {

	host := "127.0.0.1"
	port := "3306"
	user := "root"
	password := "fpf"
	DBName := "test"

	return gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, host, port, DBName,
		),
	)

	// 线上数据
	// return gorm.Open("mysql",
	// 	fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 		"root",
	// 		"abc123456",
	// 		"47.75.105.53",
	// 		"demo"),
	// )
}
