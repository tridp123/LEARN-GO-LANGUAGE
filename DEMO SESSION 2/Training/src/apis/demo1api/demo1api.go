package demo1api

import (
	"encoding/json"
	"entities"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Demo1(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "hello world")
}

func Demo2(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(response, "<b><i><u>hello tri</u></i></b>")

}

func Demo3(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fullName := vars["fullName"]
	fmt.Fprintln(response, "Hello "+fullName)

}

func Sum(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	num1 := vars["num1"]
	num2 := vars["num2"]
	i1, err1 := strconv.ParseInt(num1, 10, 64)
	if err1 != nil {
		fmt.Sprintln(err1)
	}
	i2, _ := strconv.ParseInt(num2, 10, 64)
	sum := i1 + i2
	fmt.Fprintln(response, "sum: "+fmt.Sprintf("%d", sum))

}

func Demo4(response http.ResponseWriter, request *http.Request) {
	var account entities.Account
	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Account info: ")
		fmt.Println(account.ToString())
	}
}
func Demo5(response http.ResponseWriter, request *http.Request) {
	var account entities.Account
	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Account info: ")
		fmt.Println(account.ToString())
	}
}
func Demo6(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	fmt.Println("id: " + id)

}
