package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/context"
	"github.com/riad-safowan/GOLang-SQL/models"
	"github.com/riad-safowan/GOLang-SQL/utils"
)

func UpdateProfilePicture(w http.ResponseWriter, r *http.Request) {
	email := context.Get(r, "email").(string)
	var user models.User
	user.GetUserFromDB(&email)
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("image")
	if err != nil {
		println("Error retrieving file from form-data: ", err)
		return
	}
	defer file.Close()
	println("Uploaded File: ", handler.Filename)
	println("File size: ", handler.Size)
	println("MIME Header: ", handler.Header)

	var path = "image-server/profile"
	var name = strconv.Itoa(int(user.ID))

	temp, err := ioutil.TempFile(path, name+"-*.jpg")
	if err != nil {
		println(err)
		return
	}
	defer temp.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		println(err)
	}
	temp.Write(fileBytes)

	oldurl := models.GetImageUrl(email)
	oldname := strings.TrimPrefix(oldurl, "http://"+utils.BASEURL +"/images/")
	err = os.Remove("./image-server/profile/" + oldname)
	if err != nil {
		println("Image not found")
	}

	name = strings.Split(temp.Name(), "\\")[1]
	var url = "http://"+utils.BASEURL + "/images/" + name
	user.ImageUrl = url

	models.UpdateImageUrl(url, user.ID)

	var baseResponse = &models.BaseResponse{}
	baseResponse.Data = getLoginResponse(user)
	baseResponse.Status = 200
	baseResponse.Message = "success"

	res, _ := json.Marshal(baseResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
