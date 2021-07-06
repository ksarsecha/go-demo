package main

import (
	"github.com/ksarsecha/movie_rental/handler"
	http2 "github.com/ksarsecha/movie_rental/http"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	request, err := http.NewRequest("GET", "/hello", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(handler.HelloWorld)
	handlerFunc.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf(
			"helloWorld returned invalid response code: received %d expected %d",
			status,
			http.StatusOK,
		)
	}

	expected := "Welcome to Go!"
	actual := recorder.Body.String()

	if actual != expected {
		t.Errorf("helloWorld returned invalid response data: received %s expected %s",
			actual,
			expected,
		)
	}
}

func TestRouter(t *testing.T) {
	mockServer := httptest.NewServer(http2.Router())
	response, err := http.Get(mockServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("invalid response code. Expected 200 OK received %d", response.StatusCode)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(response.Body)

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	actual := string(data)
	expected := "Welcome to Go!"

	if actual != expected {
		t.Errorf("Invalid response data. Expected %s received %s", expected, actual)
	}
}

func TestShouldReturnNotFoundForInvalidRoute(t *testing.T) {
	mockServer := httptest.NewServer(http2.Router())
	response, err := http.Get(mockServer.URL + "/invalid-route")

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Should return 404 not found, received %d", response.StatusCode)
	}
}

func TestShouldReturnMethodNotAllowedForInvalidMethod(t *testing.T) {
	mockServer := httptest.NewServer(http2.Router())

	response, err := http.Post(mockServer.URL+"/hello", "application/json", nil)

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected 405 method not allowed, received %d", response.StatusCode)
	}
}

func TestStaticFileHandler(t *testing.T) {
	mockServer := httptest.NewServer(http2.Router())

	response, err := http.Get(mockServer.URL + "/assets/")

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 ok, received %d", response.StatusCode)
	}

	actualContentType := response.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if actualContentType != expectedContentType {
		t.Errorf("Expected %s, Received %s", expectedContentType, actualContentType)
	}
}
