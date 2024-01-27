package users

import (
	"strconv"
	"user-crud/internal/utils"
	"user-crud/pkg/models"
)

type User = models.User

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

func CreateUserService(user User) User {

	user.ID = len(users) + 1
	users = append(users, user)
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
