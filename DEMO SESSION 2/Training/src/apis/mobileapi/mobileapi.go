package mobileapi

import (
	"config"
	"encoding/json"
	"entities"
	"models"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

//chu thuong => private
func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{"error": msg})
}
func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		product, err2 := mobileModel.FindAll()
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responseWithJSON(response, http.StatusOK, product)
		}
	}
}

func Find(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		vars := mux.Vars(request)
		id := vars["id"]
		product, err2 := mobileModel.Find(id)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responseWithJSON(response, http.StatusOK, product)
		}
	}
}

func Search1(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		vars := mux.Vars(request)
		keyword := vars["keyword"]
		product, err2 := mobileModel.Search1(keyword)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responseWithJSON(response, http.StatusOK, product)
		}
	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		var mobile entities.Mobile
		mobile.Id = bson.NewObjectId()
		err2 := json.NewDecoder(request.Body).Decode(&mobile)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := mobileModel.Create(&mobile)
			if err3 != nil {
				responseWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				responseWithJSON(response, http.StatusOK, mobile)
			}
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		vars := mux.Vars(request)
		id := vars["id"]
		err3 := mobileModel.Delete(id)
		if err3 != nil {
			responseWithError(response, http.StatusBadRequest, err3.Error())
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		var mobile entities.Mobile
		err2 := json.NewDecoder(request.Body).Decode(&mobile)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := mobileModel.Update(&mobile)
			if err3 != nil {
				responseWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				responseWithJSON(response, http.StatusOK, mobile)
			}
		}
	}
}
