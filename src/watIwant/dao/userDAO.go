package dao

import (
	"log"
	"watIwant/config"
	"watIwant/models"
	"watIwant/utils"
)

type UserDAO struct {
	database *config.Database
}

func NewUserDAO() UserDAO {
	var db = config.GetDatabase()

	if !db.HasTable(models.User{}) {
		db.CreateTable(models.User{})
	}
	db.AutoMigrate(models.User{})
	return UserDAO{database: db}

}

func (userDAO UserDAO) GetPublicUser(username string) *models.UserPublic {
	var user models.UserPublic
	err := userDAO.database.Where(models.UserLogin{Username: username}).Find(&models.User{}).Error
	if err != nil {
		return nil
	}
	return &user
}

func (userDAO UserDAO) CreateUser(loginUser models.UserLogin) (bool, string) {
	var user models.User

	hashedPassword, hashError := utils.HashString(loginUser.Password)
	if hashError != nil {
		return false, "error during hash"
	}

	user.Username = loginUser.Username
	user.Password = string(hashedPassword)

	err := userDAO.database.Where(models.UserLogin{Username: loginUser.Username}).Find(&models.User{}).Error
	if err == nil {
		return false, "yet exists"
	}

	if result := userDAO.database.NewRecord(user); result {
		log.Println(result)
	}
	userDAO.database.Create(&user)
	userDAO.database.NewRecord(user)
	return true, "OK"
}
