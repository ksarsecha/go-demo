package handler

import (
	"encoding/json"
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
