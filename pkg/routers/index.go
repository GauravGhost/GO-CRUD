package v1

import (
	v1Router "user-crud/pkg/routers/v1"

	"github.com/gin-gonic/gin"
)

func GlobalRouter(router *gin.RouterGroup) {
	v1 := router.Group("v1")

	v1Router.UserRouter(v1.Group("users"))
}
