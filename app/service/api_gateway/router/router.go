package router

import (
	"github.com/gin-gonic/gin"
	"test.lee/api_gateway/handler"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/print", handler.Print)
	r.GET("/echo", handler.Echo)

	r.GET("/SaveUser", handler.SaveUser)

	return r
}
