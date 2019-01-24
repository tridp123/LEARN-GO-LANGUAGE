package log2middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Log2MiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		fmt.Println("url: ", request.URL)
		today := time.Now()
		fmt.Println("date: " + today.Format("01/02/2016 15:04:05"))
		handler.ServeHTTP(response, request)
	})
}
