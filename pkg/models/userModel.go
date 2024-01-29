package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type User1 struct {
	gorm.Model
	username string
	email    string
}

func (User1) TableName() string {
	return "users"
}