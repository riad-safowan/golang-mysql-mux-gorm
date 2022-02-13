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

func CreateComment(w http.ResponseWriter, r *http.Request) {
	comment := &models.Comment{}
	utils.ParseBody(r, comment)
	comment.CreateComment()

	comments := models.GetCommentByPostId(int64(comment.PostId))
	res, _ := json.Marshal(comments)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetComments(w http.ResponseWriter, r *http.Request) {

}
func GetCommentByID(w http.ResponseWriter, r *http.Request) {

}
func GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["id"]
	ID, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	comments := models.GetCommentByPostId(ID)
	res, _ := json.Marshal(comments)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateCommentByID(w http.ResponseWriter, r *http.Request) {

}
func DeleteCommentByID(w http.ResponseWriter, r *http.Request) {

}
