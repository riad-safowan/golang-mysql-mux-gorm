package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {

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

	temp, err := ioutil.TempFile("image-server", "upload-*.png")
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

	fmt.Fprintf(w, "Successfully uploaded file as\n", temp.Name())

}
func GetProfilePicture(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	vars := mux.Vars(r)
    key := vars["name"]
	var url = "image-server/profile/"+key
	http.ServeFile(w, r, url)
}