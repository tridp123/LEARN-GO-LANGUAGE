package accountapi

import (
	"encoding/json"
	"entities"
	"net/http"
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

func Find(response http.ResponseWriter, request *http.Request) {
	account := entities.Account{
		"acc2", "abc", 22, true,
	}
	responseWithJSON(response, http.StatusOK, account)
}

func FindAll(response http.ResponseWriter, request *http.Request) {
	accounts := []entities.Account{
		entities.Account{
			"acc2", "abc2", 22, true,
		},
		entities.Account{
			"acc3", "abc3", 23, true,
		}, entities.Account{
			"acc4", "abc3", 24, false,
		},
	}
	responseWithJSON(response, http.StatusOK, accounts)
}
