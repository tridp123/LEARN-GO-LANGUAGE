package demo2api

import (
	"fmt"
	"net/http"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "hello world")
}
func Hi(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "hi cc")
}
