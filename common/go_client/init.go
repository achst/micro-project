package go_client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hopehook/micro-project/common/go_lib"
	"github.com/hopehook/micro-project/proto/go/service_order"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/broker/redis"
	"github.com/micro/go-plugins/client/grpc"
	"github.com/micro/go-plugins/wrapper/select/roundrobin"
)

// log wrapper logs every time a request is made
type logWrapper struct {
	client.Client
}

// Implements client.Wrapper as logWrapper
func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	log.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Method())
	return l.Client.Call(ctx, req, rsp)
}

// trace wrapper attaches a unique trace ID - timestamp
type traceWrapper struct {
	client.Client
}

// Implements client.Wrapper as traceWrapper
func traceWrap(c client.Client) client.Client {
	return &traceWrapper{c}
}

func (t *traceWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = map[string]string{}
	}
	md["X-Trace-Id"] = fmt.Sprintf("%d", time.Now().Unix())
	ctx = metadata.NewContext(ctx, md)
	return t.Client.Call(ctx, req, rsp)
}

// TODO
func metricsWrap(cf client.CallFunc) client.CallFunc {
	return func(ctx context.Context, addr string, req client.Request, rsp interface{}, opts client.CallOptions) error {
		t := time.Now()
		err := cf(ctx, addr, req, rsp, opts)
		log.Printf("[Metrics Wrapper] called: %s %s.%s duration: %v\n", addr, req.Service(), req.Method(), time.Since(t))
		return err
	}
}

// InitRPC init micro rpc clients
func InitRPC(conf *go_lib.Config) {
	roundRobinWrap := roundrobin.NewClientWrapper()
	// registry
	consulAddr := conf.Get("Registry", "addr")
	myRegistry := registry.NewRegistry(
		registry.Addrs(consulAddr),
	)
	// client
	myClient := grpc.NewClient(
		client.RequestTimeout(time.Second*5),
		client.DialTimeout(time.Second*5),
		client.Wrap(roundRobinWrap), // using a round robin client wrapper
		client.Wrap(traceWrap),
		client.Wrap(logWrap),
		client.WrapCall(metricsWrap),
		client.Registry(myRegistry),
	)
	// use the generated client stub
	OrderServiceClient = service_order.NewOrderServiceClient("go.micro.srv.order", myClient)
}

// InitPub init publish clients
func InitPub(conf *go_lib.Config) {
	// broker
	brokerAddr := conf.Get("Broker", "addr")
	myBroker := redis.NewBroker(
		broker.Addrs(brokerAddr),
	)
	myClient := client.NewClient(
		client.Broker(myBroker),
	)
	if err := broker.Connect(); err != nil {
		log.Fatal(err.Error())
	}
	// use the generated client stub
	Pub = &PubClient{
		publisher: micro.NewPublisher("default", myClient),
	}
}
