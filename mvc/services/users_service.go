package services

import (
	"github.com/sirtaylor88/golang-microservices/mvc/domain"
	"github.com/sirtaylor88/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
