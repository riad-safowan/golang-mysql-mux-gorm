package response

type Comment struct {
	Id int `json:"comment_id"`
	Text string `json:"text"`
	UserId int `json:"user_id"`
	PostId     int `json:"post_id"`
	UserName string `json:"user_name"`
	UserImgUrl string `json:"user_img_url"`
}
