package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	http.ListenAndServe(":8080", router())
}

func helloWorld(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to Go!")
}

func router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloWorld).Methods("GET")

	return router
}