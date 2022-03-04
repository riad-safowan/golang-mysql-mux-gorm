package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/riad-safowan/GOLang-SQL/models"
	"github.com/riad-safowan/GOLang-SQL/models/response"
	"github.com/riad-safowan/GOLang-SQL/utils"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	myUserId := context.Get(r, "user_id").(int)
	posts := models.GetAllPosts()

	responselist := []response.Post{}
	for _, v := range posts {
		user := models.User{}
		user.GetUserByIdFromDB(v.UserId)
		name := *user.FirstName + " " + *user.LastName
		isliked:=models.IsLiked(myUserId, int(v.ID))
		responselist = append(responselist, response.Post{ID: int(v.ID), Text: v.Text, UserId: v.UserId, UserName: name, UserImageUrl: *user.ImageUrl, CreatedAt: v.CreatedAt, ImageUrl: v.ImageUrl, Likes: v.Likes, Comments: v.Comments, Isliked: isliked})
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
	post.UserId = context.Get(r, "user_id").(int)

	post.CreatePost()

	var baseResponse = &models.BaseResponse{}
	baseResponse.Data = post
	baseResponse.Status = 200
	baseResponse.Message = "success"

	res, _ := json.Marshal(baseResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePostByID(w http.ResponseWriter, r *http.Request) {

}
func DeletePostByID(w http.ResponseWriter, r *http.Request) {

}
func UploadPostImage(w http.ResponseWriter, r *http.Request) {
	userId := context.Get(r, "user_id").(int)
	vars := mux.Vars(r)
	id := vars["postId"]
	postId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

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

	var path = "image-server/post"
	var name = strconv.Itoa(userId)

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
	name = strings.Split(temp.Name(), "\\")[1]
	ImageUrl := "http://" + utils.BASEURL + "/images/post/" + name

	models.UpdatePostImageUrl(int(postId), ImageUrl)

	var baseResponse = &models.BaseResponse{}
	baseResponse.Data = response.ImageResponse{ImageUrl: ImageUrl}
	baseResponse.Status = 200
	baseResponse.Message = "success"

	res, _ := json.Marshal(baseResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPostImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	vars := mux.Vars(r)
	key := vars["name"]
	var url = "image-server/post/" + key
	http.ServeFile(w, r, url)
}

func Like(w http.ResponseWriter, r *http.Request) {
	userId := context.Get(r, "user_id").(int)
	vars := mux.Vars(r)
	id := vars["id"]
	postId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	like := models.Like{PostId: int(postId), UserId: userId, UserName: models.GetUserNameById(userId)}
	l := like.CreateLike()
	if l.IsLiked {
		models.IncrementLikes(int(postId))
	} else {
		models.DecrementLikes(int(postId))
	}

	GetPostByID(w, r)
}
