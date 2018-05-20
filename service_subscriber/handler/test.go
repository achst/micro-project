package handler

import (
	"context"

	subscriber "github.com/hopehook/micro-project/proto/go/service_subscriber"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/metadata"
)

func Test(ctx context.Context, event *subscriber.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.1111111] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}
