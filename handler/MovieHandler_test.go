package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ksarsecha/movie_rental/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldCreateAMovie(t *testing.T) {
	movie := "{\"name\":\"TestMovieName\", \"description\": \"TestMovieDescription\"}"

	request, err := http.NewRequest(http.MethodPost, "/movie", bytes.NewBufferString(movie))
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()

	CreateMovieHandler(responseRecorder, request)

	assert.Equalf(t, http.StatusCreated, responseRecorder.Code,
		"Expected 201 Created, received %v. Reason: %s", responseRecorder.Code, responseRecorder.Body.String())
	assert.Equal(t, 1, len(movies))
}

func TestShouldReturnBadRequestForInvalidRequestBody(t *testing.T) {
	request, err := http.NewRequest(http.MethodPost, "/movie", nil)
	require.NoError(t, err)

	response := httptest.NewRecorder()

	CreateMovieHandler(response, request)

	assert.Equalf(t, http.StatusBadRequest, response.Code,
		"Expected 400 Bad request, received %v", response.Code)
}

func TestShouldReturnMovie(t *testing.T) {
	movies = append(movies, domain.Movie{
		Name:        "TestName",
		Description: "TestDescription",
	})

	request, err := http.NewRequest(http.MethodGet, "/movie/TestName", nil)
	require.NoError(t, err)

	request = mux.SetURLVars(request, map[string]string{"name": "TestName"})
	responseRecorder := httptest.NewRecorder()

	GetMovieHandler(responseRecorder, request)

	assert.Equalf(t, http.StatusOK, responseRecorder.Code,
		"Expected 200 OK, received %v. Reason %s", responseRecorder.Code, responseRecorder.Body.String())

	var movie domain.Movie
	decoder := json.NewDecoder(responseRecorder.Body)
	err = decoder.Decode(&movie)
	require.NoError(t, err)

	assert.Equal(t, "TestName", movie.Name)
	assert.Equal(t, "TestDescription", movie.Description)
}
