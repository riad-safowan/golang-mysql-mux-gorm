package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riad-safowan/GOLang-SQL/routes"
	"github.com/riad-safowan/GOLang-SQL/utils"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterAuthRoutes(r)
	routes.RegisterPostRoutes(r)
	routes.RegisterCommentRoutes(r)
	routes.RegisterImageUpload(r)

	// http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))
	log.Fatal(http.ListenAndServe(utils.BASEURL, r))

}
