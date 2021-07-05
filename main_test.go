package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	request, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(handler)
	handlerFunc.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned invalid response code: received %d expected %d",
			status,
			http.StatusOK,
		)
	}

	expected := "Welcome to Go!"
	actual := recorder.Body.String()

	if actual != expected {
		t.Errorf("handler returned invalid response data: received %s expected %s",
			actual,
			expected,
		)
	}
}
