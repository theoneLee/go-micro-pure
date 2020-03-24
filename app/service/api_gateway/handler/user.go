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
				"user": resp.GetUser(), //todo 这里是直接使用proto定义的model的json tag。而这个tag是无法直接在proto文件处理的。因此，要自定义返回给前端的json内容，可以自己在这里打包,更甚的，可以在定义个wrapper，然后手动将proto定义的打包成wrapper
				//"boolVal":resp.GetUser().Id,
			},
		})
	} else {
		c.JSON(200, gin.H{
			"code": resp.Error.Code,
			"msg":  resp.Error.Detail,
		})
	}
}
