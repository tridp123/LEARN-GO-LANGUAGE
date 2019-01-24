package main

import (
	"apis/demo2api"
	"apis/mobileapi"
	"fmt"
	"middlewares/basicauth"
	"middlewares/log2middleware"
	"middlewares/logmiddleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// router.HandleFunc("/api/demo/demo1", demo1api.Demo1).Methods("GET")
	// router.HandleFunc("/api/demo/demo2", demo1api.Demo2).Methods("GET")
	// router.HandleFunc("/api/demo/demo3/{fullName}", demo1api.Demo3).Methods("GET")
	// router.HandleFunc("/api/demo/sum/{num1}/{num2}", demo1api.Sum).Methods("GET")
	// router.HandleFunc("/api/demo/demo4", demo1api.Demo4).Methods("POST")
	// router.HandleFunc("/api/demo/demo5", demo1api.Demo5).Methods("PUT")
	// router.HandleFunc("/api/demo/demo6/{id}", demo1api.Demo6).Methods("DELETE")

	// router.HandleFunc("/api/account/find", accountapi.Find).Methods("GET")
	// router.HandleFunc("/api/account/findall", accountapi.FindAll).Methods("GET")

	router.HandleFunc("/api/mobile/findall", mobileapi.FindAll).Methods("GET")
	router.HandleFunc("/api/mobile/find/{id}", mobileapi.Find).Methods("GET")
	router.HandleFunc("/api/mobile/search1/{keyword}", mobileapi.Search1).Methods("GET")
	// router.HandleFunc("/api/mobile/search2/{min}/{max}", mobileapi.Search2).Methods("GET")
	router.HandleFunc("/api/mobile/create", mobileapi.Create).Methods("POST")
	router.HandleFunc("/api/mobile/delete/{id}", mobileapi.Delete).Methods("DELETE")
	router.HandleFunc("/api/mobile/update", mobileapi.Update).Methods("PUT")

	// router.Handle("/api/demo2/hello", logmiddleware.LogMiddleWare(http.HandlerFunc(demo2api.Hello))).Methods("GET")
	router.Handle("/api/demo2/hi", log2middleware.Log2MiddleWare(
		logmiddleware.LogMiddleWare(http.HandlerFunc(demo2api.Hi)))).Methods("GET")

	router.Handle("/api/demo2/hello", basicauth.BasicAuth(http.HandlerFunc(demo2api.Hello))).Methods("GET")

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		fmt.Println(err)
	}

}
