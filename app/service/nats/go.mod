module test.lee/nats-demo-sub

go 1.13

require (
	github.com/golang/protobuf v1.3.5
	github.com/nats-io/nats-streaming-server v0.17.0 // indirect
	github.com/nats-io/nats.go v1.9.1
	github.com/nats-io/stan.go v0.6.0
	test.lee/common v0.0.0

)

replace test.lee/common => ../../common
