package response

import "time"

type Post struct {
	ID           int       `json:"post_id"`
	Text         string    `json:"text"`
	UserId       int       `json:"user_id"`
	UserName     string    `json:"user_name"`
	UserImageUrl string    `json:"user_image_url"`
	CreatedAt    time.Time `json:"created_at"`
	ImageUrl     string    `json:"image_url"`
	Likes        int       `json:"likes"`
	Comments     int       `json:"comments"`
	IsLiked      bool      `json:"is_liked"`
}
