package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

func main() {
	fmt.Println("Hello World")

	root := mux.NewRouter()
	root.HandleFunc("/", HomeHandler)

	log.Println("Server is starting on port :8088")
	http.ListenAndServe(":8088", root)
}
