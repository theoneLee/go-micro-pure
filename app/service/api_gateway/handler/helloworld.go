package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	proto1 "test.lee/hello_world/proto"
)

var (
	helloWorldClient proto1.HelloWorldService
)

//// Error 错误结构体
//type Error struct {
//	Code   string `json:"code"`
//	Detail string `json:"detail"`
//}

func Init() {
	//使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)
	service := micro.NewService(micro.Name("helloworld.client"), micro.Registry(micReg), micro.Version("latest"))

	service.Init()
	helloWorldClient = proto1.NewHelloWorldService("micro_pure_helloworld", service.Client())

	//helloWorldClient = proto1.NewHelloWorldService("micro_pure_helloworld",client.DefaultClient)//todo 无法通过etcd发现到对应服务

}

func registryOptions(ops *registry.Options) {
	//etcdCfg := config.GetEtcdConfig()
	//ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", "127.0.0.1", 2379)} //todo 配置代码

}

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
