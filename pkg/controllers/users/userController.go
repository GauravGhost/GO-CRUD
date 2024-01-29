package users

import (
	"net/http"
	"user-crud/pkg/models"
	"user-crud/pkg/services/users"

	"github.com/gin-gonic/gin"
)

type User = models.User
type User1 = models.User1

// var users = utils.UserData

func GetUsers(c *gin.Context) {
	users := users.GetUsersService()
	c.JSON(http.StatusOK, users)
	return
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user := users.GetUserByIdService(id)

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user User1
	if err := c.BindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message":"Bad Body Request"})
	}
	user = users.CreateUserService(user)
	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	if err := c.BindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message":"Bad Body Request"})
	}
	user = users.UpdateUserService(id, user)

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	users.DeleteUserService(id)
	c.JSON(http.StatusNotFound, gin.H{"message": "Delete Successfully"})
}
