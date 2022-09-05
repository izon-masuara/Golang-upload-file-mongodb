package main

import (
	"fmt"
	"log"
	"net/http"
)

type test_struct struct {
	Test string
}

func image(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
	switch r.Method {
	case "GET":
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
	http.HandleFunc("/", image)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// nodemon --exec go run main.go --signal SIGTERM
