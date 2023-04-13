package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name   string  `json:"name" gorm:"text;not null;"`
	Price  int     `json:"price"`
	Rating float32 `json:"rating"`
}

// constructor
func NewItem(name string, price int, rating float32) *Item {
	return &Item{
		Name:   name,
		Price:  price,
		Rating: rating,
	}
}
