package go_client

import (
	"context"
	"log"
	"time"

	subscriber "github.com/hopehook/micro-project/proto/go/service_subscriber"
	"github.com/micro/go-micro"
	"github.com/pborman/uuid"
)

var Pub *PubClient

type PubClient struct {
	publisher micro.Publisher
}

// SendEv send events using the publisher
func (client *PubClient) Publish(ctx context.Context, url string, data string) {
	// create new event
	ev := &subscriber.Event{
		Id:        uuid.NewUUID().String(),
		Timestamp: time.Now().Unix(),
		Message:   data,
		Url:       url,
	}

	// publish an event
	if err := client.publisher.Publish(ctx, ev); err != nil {
		log.Printf("error publishing %v\n", err)
	}
}
