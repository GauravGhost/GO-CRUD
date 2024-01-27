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
			Username: "user2",
			Email:    "user2@example.com",
		},
	}
	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUserById)
	

	r.Run(":8081")
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
	return
}

func GetUserById(c *gin.Context){
	fmt.Println(c.Param("id") +  "f33a")
	id := c.Param("id")
	for _, user:= range users {
		if strconv.Itoa(user.ID) == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not Found!", "status": 202})
}