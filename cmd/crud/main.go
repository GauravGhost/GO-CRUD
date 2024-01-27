package main

import (
	v1 "user-crud/pkg/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost", "127.0.0.1"})
	v1.GlobalRouter(r.Group("api"))

	r.Run(":8082")
}
