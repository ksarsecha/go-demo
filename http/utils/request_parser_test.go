package utils

import (
	"bytes"
	"github.com/ksarsecha/movie_rental/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestShouldParseRequestSuccessfully(t *testing.T) {
	movieBytes := bytes.NewBufferString("{\"name\":\"TestMovieName\", \"description\":\"TestMovieDescription\"}")
	request, err := http.NewRequest(http.MethodGet, "/test", movieBytes)
	require.NoError(t, err)

	var actualMovie domain.Movie
	err = ParseRequest(request, &actualMovie)
	require.NoError(t, err)

	expectedMovie := domain.Movie{
		Name:        "TestMovieName",
		Description: "TestMovieDescription",
	}

	assert.Equal(t, expectedMovie, actualMovie)
}

func TestBodyShouldNotBeNil(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/test", nil)
	require.NoError(t, err)

	var actualMovie domain.Movie
	err = ParseRequest(request, &actualMovie)
	require.EqualError(t, err, "request body can not be nil")
}

func TestRequestShouldNotBeNil(t *testing.T) {
	var actualMovie domain.Movie
	err := ParseRequest(nil, actualMovie)

	require.EqualError(t, err, "request can not be nil")
}
