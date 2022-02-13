package routes

import (
	"github.com/gorilla/mux"
	"github.com/riad-safowan/GOLang-SQL/controllers"
)

var RegisterPostRoutes = func(router *mux.Router) {
	router.HandleFunc("/post", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/post/{id}", controllers.GetPostByID).Methods("GET")
	router.HandleFunc("/post/{id}", controllers.UpdatePostByID).Methods("PUT")
	router.HandleFunc("/post/{id}", controllers.DeletePostByID).Methods("DELETE")
}

var RegisterCommentRoutes = func(router *mux.Router) {
	router.HandleFunc("/comment", controllers.CreateComment).Methods("POST")
	router.HandleFunc("/comments", controllers.GetComments).Methods("GET")
	router.HandleFunc("/comment/{id}", controllers.GetCommentByID).Methods("GET")
	router.HandleFunc("/comments/{id}", controllers.GetCommentsByPostID).Methods("GET")
	router.HandleFunc("/comment/{id}", controllers.UpdateCommentByID).Methods("PUT")
	router.HandleFunc("/comment/{id}", controllers.DeleteCommentByID).Methods("DELETE")
}
