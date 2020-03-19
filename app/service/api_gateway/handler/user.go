package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	proto2 "test.lee/common/proto/user"
)

func SaveUser(c *gin.Context) {
	userName := c.Query("userName")

	u := proto2.User{UserName: userName}
	resp, err := userClient.SaveUser(context.Background(), &proto2.UserReq{User: &u})
	if err != nil {
		log.Info(err)
	}
	if resp.Success {
		c.JSON(200, gin.H{
			"code": 0,
			"data": gin.H{
				"user": resp.GetUser(),
			},
		})
	} else {
		c.JSON(200, gin.H{
			"code": resp.Error.Code,
			"msg":  resp.Error.Detail,
		})
	}
}
