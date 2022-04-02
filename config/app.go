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
	dsn := "host=ec2-54-173-77-184.compute-1.amazonaws.com user=mrzdwtxhakntkt password=ee0ba5528140c65ea1e54a8d3fa5eb54fafe27c7e4873e8a57340392f46aac16 dbname=dbqpg8ei3ic0qi port=5432"
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
