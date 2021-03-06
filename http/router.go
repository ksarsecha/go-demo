package http

import (
	"github.com/gorilla/mux"
	"github.com/ksarsecha/movie_rental/handler"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", handler.HelloWorld).Methods("GET")
	router.HandleFunc("/movie", handler.CreateMovieHandler).Methods("POST")
	router.HandleFunc("/movie/{name}", handler.GetMovieHandler).Methods("GET")

	assetsDir := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(assetsDir))
	router.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return router
}
