package dao

import (
	"log"
	"watIwant/config"
	"watIwant/models"
)

type ItemDAO struct {
	database *config.Database
}

func NewItemDAO() ItemDAO {
	var db = config.GetDatabase()

	if !db.HasTable(models.Item{}) {
		db.CreateTable(models.Item{})
	}
	db.AutoMigrate(models.Item{})
	return ItemDAO{database: db}
}

func (itemDAO ItemDAO) ReadOne(itemId string) (models.Item, error) {
	var item models.Item
	err := itemDAO.database.Find(&item).Where(&models.Item{Name: itemId}).Error
	return item, err
}

func (itemDAO ItemDAO) ReadAll() ([]models.Item, error) {
	var collection []models.Item
	err := itemDAO.database.Find(&collection).Error
	return collection, err
}

func (itemDAO ItemDAO) Insert(item models.Item) (string, error) {
	if result := itemDAO.database.NewRecord(item); result {
		log.Println(result)
	}
	itemDAO.database.Create(&item)
	itemDAO.database.NewRecord(item)

	return item.UUID, nil
}
