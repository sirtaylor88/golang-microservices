package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirtaylor88/golang-microservices/mvc/services"
	"github.com/sirtaylor88/golang-microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		// Just return the Bad request to the client
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number!",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		// Handle the error and return to the client
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}
	// return user to the client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
