package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ksarsecha/movie_rental/domain"
	"github.com/ksarsecha/movie_rental/http/utils"
	"net/http"
	"strconv"
)

var movies []domain.Movie

func CreateMovieHandler(response http.ResponseWriter, request *http.Request) {
	var movie domain.Movie
	err := utils.ParseRequest(request, &movie)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte(err.Error()))
		return
	}

	movies = append(movies, movie)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	movieBytes, _ := json.Marshal(movie)

	size, _ := response.Write(movieBytes)
	response.Header().Set("content-Length", strconv.Itoa(size))
}

func GetMovieHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	movieName := vars["name"]

	for _, movie := range movies {
		if movie.Name == movieName {
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusOK)
			movieBytes, _ := json.Marshal(movie)

			size, _ := response.Write(movieBytes)
			response.Header().Set("content-Length", strconv.Itoa(size))
			break
		}
	}
}
