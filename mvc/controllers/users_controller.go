package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirtaylor88/golang-microservices/mvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		// Just return the Bad request to the client
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("user_id must be a number!"))
		return
	}
	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		// Handle the error and return to the client
		return
	}
	// return user to the client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
