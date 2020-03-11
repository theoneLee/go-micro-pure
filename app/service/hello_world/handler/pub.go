package handler

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/pborman/uuid"
	proto "test.lee/common/proto/pub_sub"
	"time"
)

var Pub micro.Event

// send events using the publisher
func SendEv(topic string, p micro.Publisher) {
	t := time.NewTicker(time.Second)

	for _ = range t.C {
		// create new event
		ev := &proto.Event{
			Id:        uuid.NewUUID().String(),
			Timestamp: time.Now().Unix(),
			Message:   fmt.Sprintf("Messaging you all day on %s", topic),
		}

		log.Logf("publishing %+v\n", ev)

		// publish an event
		if err := p.Publish(context.Background(), ev); err != nil {
			log.Logf("error publishing: %v", err)
		}
	}
}

func SendEchoEvent(topic string, p micro.Publisher) {

	// create new event
	ev := &proto.Event{
		Id:        uuid.NewUUID().String(),
		Timestamp: time.Now().Unix(),
		Message:   fmt.Sprintf("echo===================== Messaging you all day on %s", topic),
	}
	// publish an event
	if err := p.Publish(context.Background(), ev); err != nil {
		log.Logf("error publishing: %v", err)
	}
}
