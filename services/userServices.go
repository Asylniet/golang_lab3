package services

import (
	"github.com/Asylniet/golang_lab3/database"
	"github.com/Asylniet/golang_lab3/handlers"
	"github.com/Asylniet/golang_lab3/models"
	"gorm.io/gorm"
)

func GetUsers() ([]models.User, *gorm.DB) {
	var users []models.User
	result := database.DB.Db.Find(&users)

	return users, result
}

func GetUserById(id string) (models.User, *gorm.DB) {
	var user models.User
	result := database.DB.Db.First(&user, id)
	return user, result
}

func AddUser(name, password string) *gorm.DB {
	hashedPassword := handlers.HashPassword(password)
	user := models.NewUser(name, hashedPassword)
	result := database.DB.Db.Create(&user)
	return result
}

func LoginUser(username, password string) (models.User, *gorm.DB) {
	var users []models.User
	result := database.DB.Db.Where("username = ?", username).Find(&users)

	for i := 0; i < len(users); i++ {
		if handlers.CheckPasswordHash(password, users[i].Password) {
			return users[i], result
		}
	}
	return models.User{}, result
}

func RemoveUser(id string) *gorm.DB {
	result := database.DB.Db.Delete(&models.User{}, id)
	return result
}

func UpdateUser(id, name, password string) *gorm.DB {
	hashedPassword := handlers.HashPassword(password)
	var user models.User
	result := database.DB.Db.First(&user, id)
	if user == (models.User{}) {
		return result
	}
	user.Username = name
	user.Password = hashedPassword
	result = database.DB.Db.Save(&user)
	return result
}

func RateItem(id string, rating float32) (*gorm.DB, string) {
	var item models.Item
	result := database.DB.Db.First(&item, id)
	if item == (models.Item{}) {
		return result, "No item found"
	}

	item.Rating = (item.Rating + rating) / 2
	result = database.DB.Db.Save(&item)
	return result, ""
}

func SaveItem(userId, itemId int) *gorm.DB {
	basket := models.NewBasket(userId, itemId)
	result := database.DB.Db.Create(&basket)
	return result
}

func RemoveSavedItem(userId, itemId string) *gorm.DB {
	result := database.DB.Db.Where("user_id = ? and item_id = ?", userId, itemId).Delete(models.Basket{})
	return result
}

func GetSavedItems(userId string) (*gorm.DB, []models.Item) {
	var items []models.Item
	result := database.DB.Db.Model(&models.Basket{}).Select("items.id, name, price, rating").Joins("left join items on baskets.user_id = ?", userId).Where("items.id = baskets.item_id").Scan(&items)
	return result, items
}
