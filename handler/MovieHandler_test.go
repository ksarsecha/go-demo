package handler

import (
	"bytes"
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

	response := httptest.NewRecorder()

	CreateMovieHandler(response, request)

	assert.Equalf(t, http.StatusCreated, response.Code,
		"Expected 201 Created, received %v. Reason: %s", response.Code, response.Body.String())
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
