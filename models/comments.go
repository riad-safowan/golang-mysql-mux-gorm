package models

import (
	"github.com/jinzhu/gorm"
	"github.com/riad-safowan/GOLang-SQL/config"
)

type Comment struct {
	gorm.Model        //id, time
	PostId     int    `gorm:"foreign_key"json:"post_id"`
	Text       string `gorm:""json:"text"`
	Writer     string `json:"writer"`
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&Comment{})
}

func (c *Comment) CreateComment() *Comment {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func GetAllComments() []Comment {
	var Comments []Comment
	db.Find(&Comments)
	return Comments
}

func GetCommentById(Id int64) *Comment {
	var comment Comment
	db.Where("ID=?", Id).Find(&comment)
	return &comment
}
func GetCommentByPostId(Id int64) []Comment {
	var comments []Comment
	db.Where("post_id=?", Id).Find(&comments)
	return comments
}

func DeleteCommentById(Id int64) Comment {
	var comment Comment
	db.Where("ID=?", Id).Delete(comment)
	return comment
}
