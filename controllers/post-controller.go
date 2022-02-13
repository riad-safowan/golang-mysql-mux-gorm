package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/riad-safowan/GOLang-SQL/models"
	"github.com/riad-safowan/GOLang-SQL/utils"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all post called")
	posts := models.GetAllPosts()
	res, _ := json.Marshal(posts)
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
	book := &models.Post{}
	utils.ParseBody(r, book)
	book.CreatePost()
	GetPosts(w, r)
}

func UpdatePostByID(w http.ResponseWriter, r *http.Request) {

}
func DeletePostByID(w http.ResponseWriter, r *http.Request) {

}
