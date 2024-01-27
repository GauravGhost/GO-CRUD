package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost", "127.0.0.1"})
	users = []User{
		{
			ID:       1,
			Username: "Gyanendra Kumar",
			Email:    "gyanenra@gmail.com",
		},
		{
			ID:       2,
			Username: "Gaurav Kumar",
			Email:    "gaurav@gmail.com",
		},
	}
	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUserById)
	r.POST("/users", CreateUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)

	r.Run(":8082")
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
	return
}

func GetUserById(c *gin.Context) {
	fmt.Println(c.Param("id") + "f33a")
	id := c.Param("id")
	for _, user := range users {
		if strconv.Itoa(user.ID) == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not Found!", "status": 202})
}

func CreateUser(c *gin.Context) {
	var user User
	fmt.Println("error-------------------")
	if err := c.BindJSON(&user); err != nil {
		fmt.Println("error----------------")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = len(users) + 1
	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			var UpdatedUser User
			if err := c.BindJSON(&UpdatedUser); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			UpdatedUser.ID = user.ID
			users[i] = UpdatedUser

			c.JSON(http.StatusOK, UpdatedUser)
			return
		}

	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User Not Found"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User Deleted Successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
