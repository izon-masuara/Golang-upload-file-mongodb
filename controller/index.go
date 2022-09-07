package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	mongoDB "Golang-upload-file-mongodb/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = mongoDB.ConnectDB()
var ctx = mongoDB.Ctx

type Data struct {
	ID   primitive.ObjectID
	Name string
}

func Image(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
	switch r.Method {
	case "GET":
		buf, err := ioutil.ReadFile("files/images.jpeg")

		if err != nil {

			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "image/png")
		w.Write(buf)

	case "POST":
		if err := r.ParseMultipartForm(1024); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		alias := r.FormValue("alias")

		uploadedFile, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer uploadedFile.Close()

		dir, err := os.Getwd()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filename := handler.Filename
		if alias != "" {
			filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
		}

		fileLocation := filepath.Join(dir, "files", filename)
		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, uploadedFile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("done"))
	}
}
