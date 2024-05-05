package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"watIwant/dao"
	"watIwant/models"
)

type AccountController struct {
	UserDAO dao.UserDAO
}

func NewAccountController() *AccountController {
	var controller AccountController
	controller.UserDAO = dao.NewUserDAO()
	return &controller
}

func (controller AccountController) Register(context *gin.Context) {
	var userLogin models.UserLogin

	context.BindJSON(&userLogin)

	result, message := controller.UserDAO.CreateUser(userLogin)

	if result {
		context.JSON(http.StatusCreated, nil)
	} else if message == "yet exists" {
		context.JSON(http.StatusConflict, "An account already exists")
	} else {
		context.JSON(http.StatusInternalServerError, nil)
	}
}
