package main

import (
	"Golang-upload-file-mongodb/controller"
	mongoDB "Golang-upload-file-mongodb/db"
	"fmt"
	"net/http"
)

func main() {
	mongoDB.ConnectDB()
	http.HandleFunc("/", controller.Image)
	fmt.Println("server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// nodemon --exec go run main.go --signal SIGTERM
