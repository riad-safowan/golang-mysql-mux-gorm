package models

import (
	"github.com/jinzhu/gorm"
	"github.com/riad-safowan/GOLang-SQL/config"
)

var db *gorm.DB

type Post struct {
	gorm.Model
	Text   string `gorm:""json:"text"`
	Writer string `json:"writer"`
	//id, time
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&Post{})
}

func (b *Post) CreatePost() *Post {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllPosts() []Post {
	var Posts []Post
	db.Find(&Posts)
	return Posts
}

func GetPostById(Id int64) (*Post, *gorm.DB) {
	var post Post
	db := db.Where("ID=?", Id).Find(&post)
	return &post, db
}

func DeletePostById(Id int64) Post {
	var post Post
	db.Where("ID=?", Id).Delete(post)
	return post
}
