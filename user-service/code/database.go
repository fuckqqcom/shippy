package code

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func CreateConnection() (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"abc123456",
		"47.75.105.53",
		"demo")
	return gorm.Open("mysql", dns)
}
