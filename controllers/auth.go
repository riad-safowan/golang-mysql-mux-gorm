package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/riad-safowan/GOLang-SQL/helpers"
	"github.com/riad-safowan/GOLang-SQL/models"
	"github.com/riad-safowan/GOLang-SQL/models/response"
	"github.com/riad-safowan/GOLang-SQL/utils"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4) //14 for high security
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPass string, providedPass string) (passIsValid bool, msg string) {

	err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(providedPass))

	if err != nil {
		return false, fmt.Sprint("email or password is incorrect")
	} else {
		return true, ""
	}

}

func SignUp(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "The email already exist", http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, _ := helpers.GenerateAllToken(int(user.ID), *user.Email, *user.FirstName, *user.LastName, *user.UserType)
	user.AccessToken = &accessToken
	user.RefreshToken = &refreshToken
	*user.Password = HashPassword(*user.Password)

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
	var user models.User
	var foundUser models.User
	if err := utils.ParseBody(r, &user); err != nil {
		http.Error(w, "unable to marshal json", http.StatusBadRequest)
		return
	}

	foundUser.GetUserFromDB(user.Email)
	if !(foundUser.ID > 0) {
		http.Error(w, "Incorrect email", http.StatusBadRequest)
		return
	}

	passisvalid, _ := VerifyPassword(*foundUser.Password, *user.Password)
	if !passisvalid {
		http.Error(w, "Incorrect password", http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, _ := helpers.GenerateAllToken(int(foundUser.ID), *foundUser.Email, *foundUser.FirstName, *foundUser.LastName, *foundUser.UserType)

	models.UpdateAllTokens(accessToken, refreshToken, foundUser.ID)

	foundUser.AccessToken = &accessToken
	foundUser.RefreshToken = &refreshToken

	var baseResponse = &models.BaseResponse{}
	baseResponse.Data = getLoginResponse(foundUser)
	baseResponse.Status = 200
	baseResponse.Message = "success"
	res, _ := json.Marshal(baseResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func getLoginResponse(user models.User) response.LoginResponse {
	var loginResponse = response.LoginResponse{}
	b, _ := json.Marshal(&user)
	json.Unmarshal(b, &loginResponse)
	loginResponse.ID = user.ID
	return loginResponse
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	clientToken := r.Header.Get("Authorization")
	if clientToken == "" {
		clientToken = r.Header.Get("token")
	} else if strings.HasPrefix(clientToken, "Bearer ") {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		clientToken = splitToken[1]
	} else {
		http.Error(w, "invalid authorization token", http.StatusUnauthorized)
		return
	}

	if clientToken == "" {
		http.Error(w, "No Authorization header provided", http.StatusUnauthorized)
		return
	}
	// handle access token
	claims, err := helpers.ValidateToken(clientToken)

	if err != "" || claims.Token_type != "refresh_token" {
		http.Error(w, err, http.StatusUnauthorized)
		return
	}

	var foundUser models.User
	foundUser.GetUserFromDB(&claims.Email)
	if !(foundUser.ID > 0) {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, _ := helpers.GenerateAllToken(int(foundUser.ID), *foundUser.Email, *foundUser.FirstName, *foundUser.LastName, *foundUser.UserType)

	models.UpdateAllTokens(accessToken, refreshToken, foundUser.ID)

	foundUser.AccessToken = &accessToken
	foundUser.RefreshToken = &refreshToken

	var baseResponse = &models.BaseResponse{}
	baseResponse.Data = response.Token{RefreshToken: refreshToken, AccessToken: accessToken}
	baseResponse.Status = 200
	baseResponse.Message = "success"
	res, _ := json.Marshal(baseResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
