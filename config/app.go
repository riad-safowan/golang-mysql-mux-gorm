package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/troikasoft?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to db")
	db = d
}

func GetDB() *gorm.DB{
	if db == nil {
		Connect()
	}
	return db
}
