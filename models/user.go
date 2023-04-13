package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

// constructor
func NewUser(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}
