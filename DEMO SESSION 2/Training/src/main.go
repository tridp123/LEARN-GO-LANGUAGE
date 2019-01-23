package main

import (
	"apis/accountapi"
	"apis/demo1api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/demo/demo1", demo1api.Demo1).Methods("GET")
	router.HandleFunc("/api/demo/demo2", demo1api.Demo2).Methods("GET")
	router.HandleFunc("/api/demo/demo3/{fullName}", demo1api.Demo3).Methods("GET")
	router.HandleFunc("/api/demo/sum/{num1}/{num2}", demo1api.Sum).Methods("GET")
	router.HandleFunc("/api/demo/demo4", demo1api.Demo4).Methods("POST")
	router.HandleFunc("/api/demo/demo5", demo1api.Demo5).Methods("PUT")
	router.HandleFunc("/api/demo/demo6/{id}", demo1api.Demo6).Methods("DELETE")

	router.HandleFunc("/api/account/find", accountapi.Find).Methods("GET")
	router.HandleFunc("/api/account/findall", accountapi.FindAll).Methods("GET")

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		fmt.Println(err)
	}

}
