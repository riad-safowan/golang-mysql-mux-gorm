package models

import (
	"github.com/jinzhu/gorm"
	"github.com/riad-safowan/GOLang-SQL/config"
)

type User struct {
	gorm.Model
	FirstName    *string `json:"first_name" validate:"required,min=2,max=100"`
	LastName     *string `json:"last_name" validate:"required"`
	ImageUrl     *string `json:"image_url"`
	Password     *string `json:"password" validate:"required,min=6,max=100"`
	Email        *string `json:"email" validate:"email,required"`
	AccessToken  *string `json:"access_token" gorm:"size:300"`
	RefreshToken *string `json:"refresh_token" gorm:"size:300"`
	UserType     *string `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) IsUserExist() bool {
	var user = User{}
	db.Where("email = ?", u.Email).First(&user)
	println(&user)
	if user.ID > 0 {
		println(user.FirstName)
		return true
	} else {
		return false
	}
}

func (u *User) GetUserFromDB(email *string) {
	db.Where("email = ? ", email).First(&u)
}
func (u *User) GetUserByIdFromDB(uid int) {
	db.Where("id = ?", uid).First(&u)
}
func GetUserNameById(uid int) string {
	var u = User{}
	db.Select("first_name , last_name").Where("id = ?", uid).First(&u)
	return *u.FirstName + " " + *u.LastName
}

func (u *User) InsertToDb() {
	db.Create(&u)
}

func UpdateAllTokens(signedAccessToken string, signedRefreshToken string, userId uint) {
	db.Model(&User{}).Where("id=?", userId).Updates(User{AccessToken: &signedAccessToken, RefreshToken: &signedRefreshToken})
}

func UpdateImageUrl(url string, userId uint) {
	db.Model(&User{}).Where("id=?", userId).Updates(User{ImageUrl: &url})
}

func GetImageUrl(email string) string {
	type Url struct {
		ImageUrl string
	}
	var urlContainer Url
	db.Model(&User{}).Where("email = ? ", email).Scan(&urlContainer)
	return urlContainer.ImageUrl
}
