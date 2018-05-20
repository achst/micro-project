package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/hopehook/micro-project/common/go_lib"
	"github.com/hopehook/micro-project/common/go_util"
	"github.com/hopehook/micro-project/proto/go/service_order"
	"github.com/hopehook/micro-project/service_order/g"
	"github.com/hopehook/micro-project/service_order/handler"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
)

const DefaultConf = "my.cnf"
const MicroName = "go.micro.srv.order"
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
}

func main() {
	// implements the server.HandlerWrapper
	logWrapper := func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			log.Printf("[%v] server request: %s\n", time.Now(), req.Method())
			return fn(ctx, req, rsp)
		}
	}
	// registry
	consulAddr := g.Conf.Get("Registry", "addr")
	reg := registry.NewRegistry(
		registry.Addrs(consulAddr),
	)
	// Create a new service. Optionally include some options here.
	service := grpc.NewService(
		micro.Name(MicroName),
		micro.Version(MicroVersion),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		// wrap the handler
		micro.WrapHandler(logWrapper),
		micro.Registry(reg),
	)

	// Init will parse the command line flags.
	// optionally setup command line usage
	service.Server().Init(
		server.Wait(true),
	)
	//service.Init()

	// Register handler
	service_order.RegisterOrderServiceHandler(service.Server(), new(handler.OrderHandler))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
