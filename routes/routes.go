package routes

import (
	"github.com/gorilla/mux"
	"github.com/riad-safowan/GOLang-SQL/controllers"
	// "github.com/riad-safowan/GOLang-SQL/middleware"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/users/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
}

// var RegisterUserRoute = func(router *mux.Router) {
// 	router.HandleFunc("/user/refresh_token", controllers.RefreshToken).Methods("GET")
// 	router.HandleFunc("/users", middleware.Authenticate(controllers.GetUsers)).Methods("GET")
// 	router.HandleFunc("/user/:user_id", middleware.Authenticate(controllers.GetUser)).Methods("GET")
// }

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
