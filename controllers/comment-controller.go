package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/riad-safowan/GOLang-SQL/models"
	"github.com/riad-safowan/GOLang-SQL/models/response"
	"github.com/riad-safowan/GOLang-SQL/utils"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	myUserId := context.Get(r, "user_id").(int)
	vars := mux.Vars(r)
	id := vars["postId"]
	postId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	comment := &models.Comment{UserId: myUserId, PostId: int(postId)}
	utils.ParseBody(r, comment)
	comment.CreateComment()
	models.IncrementComments(comment.PostId)
	GetCommentsByPostID(w, r)
}
func GetComments(w http.ResponseWriter, r *http.Request) {

}
func GetCommentByID(w http.ResponseWriter, r *http.Request) {

}
func GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["postId"]
	postId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	comments := models.GetCommentByPostId(postId)
	responselist := []response.Comment{}
	for _, comment := range comments {
		user := models.User{}
		user.GetUserByIdFromDB(comment.UserId)
		name := *user.FirstName + " " + *user.LastName
		responselist = append(responselist, response.Comment{Id: int(comment.ID), Text: comment.Text, UserId: int(user.ID), PostId: comment.PostId, UserName: name, UserImgUrl: user.ImageUrl})
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
func UpdateCommentByID(w http.ResponseWriter, r *http.Request) {

}
func DeleteCommentByID(w http.ResponseWriter, r *http.Request) {

}
