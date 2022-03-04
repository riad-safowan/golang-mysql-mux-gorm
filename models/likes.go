package models

import (
	"github.com/jinzhu/gorm"
	"github.com/riad-safowan/GOLang-SQL/config"
)

type Like struct {
	gorm.Model        //id, time
	PostId     int    `gorm:"foreign_key"json:"post_id"`
	UserId     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	IsLiked    bool
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&Like{})
}

func (c *Like) CreateLike() *Like {
	var like Like
	result := db.Where("user_id=? AND post_id=?", c.UserId, c.PostId).First(&like)
	// result := db.Exec("UPDATE likes SET is_Liked = !is_liked WHERE user_id = ? AND post_id = ?", c.UserId, c.PostId)
	if result.RowsAffected == 0 {
		c.IsLiked = true
		db.NewRecord(c)
		db.Create(&c)
	} else {
		// result=db.Model(&Like{}).Updates(Like{IsLiked: !like.IsLiked}).Where("user_id=? AND post_id=?", c.UserId, c.PostId)
		db.Exec("UPDATE likes SET is_Liked = ? WHERE user_id = ? AND post_id = ?", !like.IsLiked, c.UserId, c.PostId)
		c.IsLiked = !like.IsLiked
	}
	return c
}

func GetAllLikes() []Like {
	var Likes []Like
	db.Find(&Likes)
	return Likes
}

func GetLikeById(Id int64) *Like {
	var Like Like
	db.Where("ID=?", Id).Find(&Like)
	return &Like
}

func GetLikesByPostId(Id int64) []Like {
	var Likes []Like
	db.Where("post_id=?", Id).Find(&Likes)
	return Likes
}

func DeleteLikeById(Id int64) Like {
	var Like Like
	db.Where("ID=?", Id).Delete(Like)
	return Like
}
func DeleteLikeByIdID(UserID int64, PostId int) Like {
	var Like Like
	db.Where("user_id=? AND post_id=?", UserID, PostId).Delete(Like)
	return Like
}
func IsLiked(UserID int, PostId int) bool {
	var liked = Like{}
	db.Select("is_liked").Where("user_id=? AND post_id=?", UserID, PostId).First(&liked)
	return liked.IsLiked
}
