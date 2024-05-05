package dao

import (
	"golang.org/x/crypto/bcrypt"
	"watIwant/config"
	"watIwant/models"
)

type AuthDAO struct {
	database *config.Database
}

func NewAuthDAO() AuthDAO {
	var db = config.GetDatabase()
	return AuthDAO{database: db}
}

func (auth AuthDAO) Login(login string, password string) bool {
	var user models.User
	err := auth.database.Where(models.UserLogin{Username: login}).Find(&user).Limit(1).Error
	if err != nil {
		return false
	}

	compareError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if compareError != nil {
		return false
	} else {
		return true
	}
}
