package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	proto1 "test.lee/common/proto/hello_world"
)

//// Error 错误结构体
//type Error struct {
//	Code   string `json:"code"`
//	Detail string `json:"detail"`
//}

func Echo(c *gin.Context) {

	userName := c.Query("userName")
	content := c.Query("content")

	//go micro client rpc call
	resp, err := helloWorldClient.Echo(context.TODO(), &proto1.EchoRequest{UserName: userName, Content: content})
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": gin.H{
			"user": resp.GetUser(),
		},
	})
}

func Print(c *gin.Context) {
	//go micro client rpc call
	resp, err := helloWorldClient.Print(context.TODO(), &proto1.PrintRequest{Id: 11})
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": gin.H{
			"print": resp.GetPrintContent(),
		},
	})
}
