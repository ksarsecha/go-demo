package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", handler).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to Go!")
}
