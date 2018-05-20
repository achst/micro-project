package router

import (
	"context"

	subscriber "github.com/hopehook/micro-project/proto/go/service_subscriber"
	"github.com/hopehook/micro-project/service_subscriber/handler"
)

var Router = map[string]func(context.Context, *subscriber.Event) error{}

func Init() {
	Router["/test"] = handler.Test
}
