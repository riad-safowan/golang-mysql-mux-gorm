package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riad-safowan/GOLang-SQL/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterAuthRoutes(r)
	routes.RegisterPostRoutes(r)
	routes.RegisterCommentRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("192.168.31.215:9090", r))

}
