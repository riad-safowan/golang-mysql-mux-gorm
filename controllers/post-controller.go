package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/riad-safowan/GOLang-SQL/models"
	"github.com/riad-safowan/GOLang-SQL/models/response"
	"github.com/riad-safowan/GOLang-SQL/utils"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := models.GetAllPosts()

	responselist := []response.Post{}
	for _, v := range posts {
		user := models.User{}
		user.GetUserByIdFromDB(v.UserId)
		name := *user.FirstName + " " + *user.LastName
		responselist = append(responselist, response.Post{ID: int(v.ID), Text: v.Text, UserId: v.UserId, UserName: name, UserImageUrl: *user.ImageUrl, CreatedAt: v.CreatedAt})
	}

	var baseResponse = &models.BaseResponse{}
	baseResponse.Data = responselist
	baseResponse.Status = 200
	baseResponse.Message = "success"

	res, _ := json.Marshal(baseResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["id"]
	ID, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	post, _ := models.GetPostById(ID)

	res, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	post := &models.Post{}
	utils.ParseBody(r, post)
	post.CreatePost()
	GetPosts(w, r)
}

func UpdatePostByID(w http.ResponseWriter, r *http.Request) {

}
func DeletePostByID(w http.ResponseWriter, r *http.Request) {

}
