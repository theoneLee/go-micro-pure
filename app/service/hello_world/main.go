package main

import (
	"fmt"
	p "github.com/golang/protobuf/proto"
	"github.com/nats-io/stan.go"
	"github.com/pborman/uuid"
	proto "test.lee/common/proto/pub_sub"
	"time"

	//"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	s "test.lee/common/proto/hello_world"
	"test.lee/hello_world/handler"
)

func main() {

	//Run pubsub client(pub)
	Init()

	//自引入的nats消息队列
	InitNats()

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

	//Init() //todo 上面service.Run()会进入阻塞。导致不会调用这里的Init()
}

func InitNats() {
	clusterID := "test-cluster"                  //需要和sub方一致
	clientID := "stan-sub-demo2"                 //需要每个client都唯一
	sc, err := stan.Connect(clusterID, clientID) //todo 会向clusterID一致的sub方的subject为foo，投递消息
	if err != nil {
		panic(err)
	}
	subject := "foo"
	sc.Publish(subject, []byte("Hello World in rpc")) // does not return until an ack has been received from NATS Streaming
	ev := &proto.Event{
		Id:        uuid.NewUUID().String(),
		Timestamp: time.Now().Unix(),
		Message:   fmt.Sprintf("Messaging you all day on %s", subject),
	}
	bytes, _ := p.Marshal(ev) //自己序列化结构体，再用nats发布,nats的sub接受后，再使用proto反序列化
	sc.Publish("model", bytes)

	// Close connection
	sc.Close()
}

func Init() {
	micReg := etcd.NewRegistry(registryOptions)
	// create a service
	pubService := micro.NewService(
		micro.Name("go.micro.cli.pubsub"),
		micro.Registry(micReg),
	)
	// parse command line
	pubService.Init()

	// create publisher
	handler.Pub = micro.NewPublisher("example.topic.pubsub.1", pubService.Client()) // todo 改为使用Init函数。且获取Pub使用getter
	pub1 := micro.NewPublisher("example.topic.pubsub.1", pubService.Client())
	pub2 := micro.NewPublisher("example.topic.pubsub.2", pubService.Client())

	// pub to topic 1
	go handler.SendEv("example.topic.pubsub.1", pub1)
	go handler.SendEv("example.topic.pubsub.1", pub1)
	// pub to topic 2
	go handler.SendEv("example.topic.pubsub.2", pub2)
	go handler.SendEv("example.topic.pubsub.2", pub2)
	go handler.SendEv("example.topic.pubsub.2", pub2)

	//// block forever
	//select {}
}

func registryOptions(ops *registry.Options) {
	//etcdCfg := config.GetEtcdConfig()
	//ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", "127.0.0.1", 2379)}
}
