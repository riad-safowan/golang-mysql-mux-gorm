package config

import (
	"fmt"

	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/troikasoft?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := "host=ec2-3-222-204-187.compute-1.amazonaws.com user=qwunujnjqzytpj password=cbb3ed7b6e5c6d736cc8a04b7266f478270550edc38849aacdd8493ecd398190 dbname=ddsnsmg54h89a9 port=5432"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to db")
	// d.LogMode(true)
	db = d
}

func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}
	return db
}
