package handler

import (
	"fmt"
	"net/http"
)

func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to Go!")
}
