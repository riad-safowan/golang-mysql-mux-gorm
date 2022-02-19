package models

import (
	"github.com/jinzhu/gorm"
	"github.com/riad-safowan/GOLang-SQL/config"
)

type User struct {
	gorm.Model
	FirstName    *string `json:"first_name" validate:"required,min=2,max=100"`
	LastName     *string `json:"last_name" validate:"required"`
	Password     *string `json:"password" validate:"required,min=6,max=100"`
	Email        *string `json:"email" validate:"email,required"`
	PhoneNumber  *string `json:"phone_number" validate:"required"`
	AccessToken  *string `json:"access_token" gorm:"size:1000"`
	RefreshToken *string `json:"refresh_token"`
	UserType     *string `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) IsUserExist() bool {
	var user = User{}
	db.Where("email = ? OR phone_number=?", u.Email, u.PhoneNumber).First(&user)
	println(&user)
	if user.ID > 0 {
		println(user.FirstName)
		return true
	} else {
		return false
	}
}

func (u *User) InsertToDb() {
	db.Create(&u)
}
