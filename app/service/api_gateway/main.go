package main

import (
	"os"
	"test.lee/api_gateway/handler"
	"test.lee/api_gateway/router"
)

func main() {
	handler.Init()

	r := router.InitRouter()

	// 默认 3001 端口
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	r.Run(":" + port)
}
