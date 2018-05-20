package subscriber

import (
	"context"
	"log"

	subscriber "github.com/hopehook/micro-project/proto/go/service_subscriber"
	"github.com/hopehook/micro-project/service_subscriber/router"
	"github.com/micro/go-micro/metadata"
)

// All methods of Sub will be executed when
// a message is received
type DefaultSub struct{}

// Method can be of any name
func (s *DefaultSub) Process(ctx context.Context, event *subscriber.Event) error {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Worker panic: %v", err)
		}
	}()
	fn := router.Router[event.Url]
	return fn(ctx, event)
}

// Log sub will get the same event too
func (s *DefaultSub) Log(ctx context.Context, event *subscriber.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Printf("[pubsub] Received event %+v with metadata %+v\n", event, md)
	return nil
}
