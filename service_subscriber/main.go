package main

import (
	"log"
	"strings"

	"github.com/hopehook/micro-project/service_subscriber/subscriber"
	"github.com/micro/go-micro"

	"github.com/hopehook/micro-project/common/go_lib"
	"github.com/hopehook/micro-project/common/go_util"
	"github.com/hopehook/micro-project/service_subscriber/g"
	"github.com/hopehook/micro-project/service_subscriber/router"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/broker/redis"
)

const DefaultConf = "my.cnf"
const MicroName = "go.micro.srv.service_subscriber"
const MicroVersion = "v1"

func init() {
	confPath := go_util.GetConfPath(DefaultConf)
	// 初始化配置文件
	g.Conf = go_lib.InitConfig(confPath)
	// 初始化基准路径
	g.Path = confPath[:strings.LastIndex(confPath, "/")]
	// 初始化基础性日志
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// 初始化全局变量
	g.InitGlobal()
	// 初始化路由
	router.Init()
}

func main() {
	// registry
	consulAddr := g.Conf.Get("Registry", "addr")
	myRegistry := registry.NewRegistry(
		registry.Addrs(consulAddr),
	)
	// broker
	brokerAddr := g.Conf.Get("Broker", "addr")
	myBroker := redis.NewBroker(
		broker.Addrs(brokerAddr),
	)
	// create a service
	service := micro.NewService(
		micro.Name(MicroName),
		micro.Version(MicroVersion),
		micro.Registry(myRegistry),
		micro.Broker(myBroker),
	)
	if err := broker.Connect(); err != nil {
		log.Fatal(err.Error())
	}
	// parse command line
	service.Server().Init(
		server.Wait(true),
	)

	// register service-service_subscriber with queue, each message is delivered to a unique service-service_subscriber
	micro.RegisterSubscriber("default", service.Server(), new(subscriber.DefaultSub), server.SubscriberQueue("queue.default"))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
