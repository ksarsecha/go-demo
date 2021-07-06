package http

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	rf := func(method, path string) *http.Request {
		request, err := http.NewRequest(method, path, nil)
		require.NoError(t, err)
		return request
	}

	testRouter(t, rf(http.MethodGet, "/hello"))
	testRouter(t, rf(http.MethodPost, "/movie"))
	testRouter(t, rf(http.MethodGet, "/assets/index.html"))
}

func testRouter(t *testing.T, request *http.Request) {
	router := Router()

	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)

	assert.NotEqual(t, http.StatusNotFound, responseRecorder.Code, "For route: %s", request.URL)
	assert.NotEqual(t, http.StatusMethodNotAllowed, responseRecorder.Code, "For route: %s", request.URL)
}
