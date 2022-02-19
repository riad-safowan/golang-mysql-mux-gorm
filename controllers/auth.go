package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/riad-safowan/GOLang-SQL/helpers"
	"github.com/riad-safowan/GOLang-SQL/models"
	"github.com/riad-safowan/GOLang-SQL/models/response"
	"github.com/riad-safowan/GOLang-SQL/utils"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPass string, providedPass string) (passIsValid bool, msg string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(providedPass))
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(userPass, providedPass)
		return false, fmt.Sprint("email or password is incorrect")
	} else {
		return true, ""
	}

}

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := utils.ParseBody(r, &user); err != nil {
		http.Error(w, "unable to marshal json", http.StatusBadRequest)
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if user.IsUserExist() {
		http.Error(w, "The email or phonenumber already exist", http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, _ := helpers.GenerateAllToken(*user.Email, *user.FirstName, *user.LastName, *user.UserType)
	user.AccessToken = &accessToken
	user.RefreshToken = &refreshToken

	user.InsertToDb()
	var baseResponse = &models.BaseResponse{}
	baseResponse.Data = getLoginResponse(user)
	baseResponse.Status = 200
	baseResponse.Message = "success"

	res, _ := json.Marshal(baseResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func Login(w http.ResponseWriter, r *http.Request) {

}

func getLoginResponse(user models.User) response.LoginResponse {
	var loginResponse = response.LoginResponse{}
	b, _ := json.Marshal(&user)
	json.Unmarshal(b, &loginResponse)
	loginResponse.ID = user.ID
	return loginResponse
}
