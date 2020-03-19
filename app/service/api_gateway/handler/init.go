package handler

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	proto1 "test.lee/common/proto/hello_world"
	proto2 "test.lee/common/proto/user"
)

var (
	helloWorldClient proto1.HelloWorldService
	userClient       proto2.UserAndNewsService
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
	userClient = proto2.NewUserAndNewsService("micro_pure_user_news", service.Client())

}

func registryOptions(ops *registry.Options) {
	//etcdCfg := config.GetEtcdConfig()
	//ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", "127.0.0.1", 2379)} //todo 配置代码

}
