package repositories

import (
	"log"
	"user-crud/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User = models.User1
type UserTable struct {
	userDb string
	db     *gorm.DB
}

func UserRepository() (*UserTable, error) {
	userDb := "user-crud/database/user.sqlite"
	db, err := gorm.Open(sqlite.Open(userDb), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &UserTable{
		userDb: userDb,
		db:     db,
	}, nil
}

func (model *UserTable) CreateUser(data *User) {
	response := model.db.Create(&data)

	if response.Error != nil {
		log.Fatalf("failed to create the user")
	}
}
