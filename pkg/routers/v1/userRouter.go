package v1

import (
	"user-crud/pkg/controllers/users"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	r.GET("/", users.GetUsers)
	r.GET("/:id", users.GetUserById)
	r.POST("/", users.CreateUser)
	r.PUT("/:id", users.UpdateUser)
	r.DELETE("/:id", users.DeleteUser)
}
