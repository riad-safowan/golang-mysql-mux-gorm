package models

import (
	"github.com/jinzhu/gorm"
	"github.com/riad-safowan/GOLang-SQL/config"
)

var db *gorm.DB

type Post struct {
	gorm.Model
	Text     string `json:"text"validate:"required"`
	UserId   int    `json:"user_id"`
	ImageUrl string `json:"image_url"`
	Likes    int    `json:"likes"`
	Comments int    `json:"comments"`
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
	db.Order("updated_at desc").Find(&Posts)
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

func UpdatePostImageUrl(id int, url string) {
	db.Model(&Post{}).Where("id=?", id).Updates(Post{ImageUrl: url})
}

func IncrementLikes(id int) {
	db.Exec("UPDATE posts SET likes = likes + 1 WHERE id = ?", id)
}
func IncrementComments(id int) {
	db.Exec("UPDATE posts SET comments = comments + 1 WHERE id = ?", id)
}
func DecrementLikes(id int) {
	db.Exec("UPDATE posts SET likes = likes - 1 WHERE id = ?", id)
}
func DecrementComments(id int) {
	db.Exec("UPDATE posts SET comments = comments - 1 WHERE id = ?", id)
}
