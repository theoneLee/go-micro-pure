package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/util/log"
	proto "test.lee/common/proto/pub_sub"
)

// todo: All methods of Sub will be executed when a message is received
type Sub struct{}

// Method can be of any name
func (s *Sub) Process(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.1 process1] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

//func (s *Sub) Process2(ctx context.Context, event *proto.Event) error {
//	md, _ := metadata.FromContext(ctx)
//	log.Logf("[pubsub.1 process2] Received event %+v with metadata %+v\n", event, md)
//	// do something with event
//	return nil
//}

// Alternatively a function can be used
func subEv(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.2] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

func main() {
	micReg := etcd.NewRegistry(registryOptions)

	// create a service
	service := micro.NewService(
		micro.Name("go.micro.srv.pubsub"),
		micro.Registry(micReg),
	)
	// parse command line
	service.Init()

	// register subscriber //TODO 在开启多个pub_sub实例情况下，client pub一个消息，这多个pub_sub实例都能获得该消息。而下面的只有一个实例能拿到消息。
	micro.RegisterSubscriber("example.topic.pubsub.1", service.Server(), new(Sub))

	// register subscriber with queue, each message is delivered to a unique subscriber
	micro.RegisterSubscriber("example.topic.pubsub.2", service.Server(), subEv, server.SubscriberQueue("queue.pubsub"))

	//ss,_ :=broker.DefaultBroker.Subscribe("example.topic.pubsub.3", Sub)
	//ss.Unsubscribe()
	//broker.Broker().Subscribe()
	////sub:=service.Server().NewSubscriber("example.topic.pubsub.3", new(Sub))
	////service.Server().Subscribe(sub)
	//go func() {
	//	time.Sleep(10*time.Second)
	//
	//}()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	//etcdCfg := config.GetEtcdConfig()
	//ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", "127.0.0.1", 2379)} //todo 配置代码

}
