module test.lee/hello_world

go 1.13

require (
	github.com/golang/protobuf v1.3.4
	github.com/micro/go-micro/v2 v2.2.0
	github.com/nats-io/nats-streaming-server v0.17.0 // indirect
	github.com/nats-io/stan.go v0.6.0
	github.com/pborman/uuid v1.2.0
	test.lee/common v0.0.0
)

replace test.lee/common => ../../common
