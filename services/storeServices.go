package services

import (
	"github.com/Asylniet/golang_lab3/database"
	"github.com/Asylniet/golang_lab3/models"
	"gorm.io/gorm"
)

func GetItems() ([]models.Item, *gorm.DB) {
	var items []models.Item
	result := database.DB.Db.Find(&items)
	return items, result
}

func GetItem(id string) (models.Item, *gorm.DB) {
	var item models.Item
	result := database.DB.Db.First(&item, id)
	return item, result
}

func AddItem(name string, price int) *gorm.DB {
	item := models.NewItem(name, price, 0)
	result := database.DB.Db.Create(&item)
	return result
}

func RemoveItem(id string) *gorm.DB {
	result := database.DB.Db.Delete(&models.Item{}, id)
	return result
}

func UpdateItem(name, id string, price int) (*gorm.DB, string) {
	var item models.Item
	result := database.DB.Db.First(&item, id)

	if item == (models.Item{}) {
		return result, "No item found"
	}

	item.Name = name
	item.Price = price
	result = database.DB.Db.Save(&item)

	return result, ""
}

func SearchItems(name, filter, sort string) ([]models.Item, *gorm.DB) {
	var items []models.Item
	result := database.DB.Db.Where("LOWER(name) LIKE ?", "%"+name+"%").Order(sort).Find(&items)
	//items = models.FilterItems(items, filter)
	return items, result
}
