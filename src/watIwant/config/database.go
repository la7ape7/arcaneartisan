package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"time"
)

type IDatabase interface {
}

type Database struct {
	*gorm.DB
}

var database *Database

func InitializeDatabase() *Database {

	databaseConfig := GetConfiguration().Database

	for database == nil {
		db, err := gorm.Open("sqlite3", databaseConfig.SqliteFilePath)
		if err != nil {
			log.Println("DB : Connection failed at " + databaseConfig.SqliteFilePath + ", waiting 5 secondes before retrying")
			time.Sleep(5 * time.Second)
		} else {
			log.Println("DB : Connexion succeed !")
			database = &Database{db}
			db.LogMode(true)
		}
	}
	return database
}

func GetDatabase() *Database {
	if database == nil {
		InitializeDatabase()
	}
	return database
}
