package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/util/log"
	c "test.lee/common/cfg"
	s "test.lee/common/proto/user"
	"test.lee/user/handler"
	"test.lee/user/model/util"
)

func main() {
	// todo grpc service
	c.InitConfig("./cfg/config.yml")

	util.InitMysql()

	micReg := etcd.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name("micro_pure_user_news"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	s.RegisterUserAndNewsHandler(service.Server(), new(handler.Service))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
func registryOptions(ops *registry.Options) {
	//etcdCfg := config.GetEtcdConfig()
	//ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", "127.0.0.1", 2379)}
}
