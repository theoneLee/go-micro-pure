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
	"test.lee/user/model"
	"test.lee/user/model/util"
)

func main() {
	// todo grpc service
	c.InitConfig("./cfg/config.yml")

	util.InitMysql()
	//自动迁移 只会 创建表、缺失的列、缺失的索引， 不会 更改现有列的类型或删除未使用的列，以此来保护您的数据。
	util.GetDb().AutoMigrate(&model.User{}, &model.News{})

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
