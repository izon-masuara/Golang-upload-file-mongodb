package main

import (
	mongoDB "Golang-upload-file-mongodb/db"
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = mongoDB.ConnectDB()

type Data struct {
	ID   primitive.ObjectID
	Name string
}

func image(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
	switch r.Method {
	case "GET":
		cur, err := collection.Find(ctx, bson.M{})
		if err != nil {
			panic(err)
		}
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			var row Data
			err := cur.Decode(&row)
			if err != nil {
				panic(err)
			}
			fmt.Println(row)
		}

		fmt.Fprintf(w, "Hello world")
	case "POST":
		fmt.Fprintf(w, "Ini adalah %v", r.FormValue("message"))
		file, handler, err := r.FormFile("image")
		if err != nil {
			fmt.Println(err)
		}

		defer file.Close()
		fmt.Println(handler.Filename)
		fmt.Println(handler.Size)
		fmt.Println(handler.Header)
	}
}

func main() {
	mongoDB.ConnectDB()
	http.HandleFunc("/", image)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// nodemon --exec go run main.go --signal SIGTERM
