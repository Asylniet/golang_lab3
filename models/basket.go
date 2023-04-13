package models

import "gorm.io/gorm"

type Basket struct {
	gorm.Model
	UserId int
	ItemId int
}

// constructor
func NewBasket(userId, itemId int) *Basket {
	return &Basket{
		UserId: userId,
		ItemId: itemId,
	}
}
