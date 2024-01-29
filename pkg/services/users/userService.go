package users

import (
	"strconv"
	"user-crud/internal/utils"
	"user-crud/pkg/models"
	"user-crud/pkg/repositories"
)

type User = models.User
type User1 = models.User1

var users = utils.UserData

func GetUsersService() []User {
	return users
}

func GetUserByIdService(id string) User {
	for _, user := range users {
		if strconv.Itoa(user.ID) == id {
			return user
		}
	}
	return User{}
}

func CreateUserService(user User1) User1 {

	// user.ID = len(users) + 1

	// users = append(users, user)
	repo, err := repositories.UserRepository()
	if err != nil {
		return User1{}
	}
	repo.CreateUser(&user)
	return user
}

func UpdateUserService(id string, updatedUser User) User {
	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			updatedUser.ID = user.ID
			users[i] = updatedUser
			return updatedUser
		}
	}
	return User{}
}

func DeleteUserService(id string) bool {

	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}
