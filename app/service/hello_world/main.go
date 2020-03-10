package main

import (
	"fmt"
	//"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"test.lee/hello_world/handler"
	s "test.lee/hello_world/proto"
)

func main() {

	//Run pubsub client(pub)

	micReg := etcd.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name("micro_pure_helloworld"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	s.RegisterHelloWorldHandler(service.Server(), new(handler.Service))

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
