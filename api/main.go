package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"strings"

	"github.com/hopehook/micro-project/api/g"
	"github.com/hopehook/micro-project/api/router"
	"github.com/hopehook/micro-project/common/go_client"
	"github.com/hopehook/micro-project/common/go_lib"
	"github.com/hopehook/micro-project/common/go_util"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-web"
)

const DefaultConf = "my.cnf"
const MicroName = "go.micro.api.api"
const Version = "v1"

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
	// 初始化路由注册
	router.Init()
	// 初始化 rpc 客户端
	go_client.InitRPC(g.Conf)
	// 初始化 pub 客户端
	go_client.InitPub(g.Conf)
}

func main() {
	// http server
	srv := &http.Server{
		ReadTimeout:    15 * time.Second, // 读取 http 请求超时时间(header 和 body)
		WriteTimeout:   15 * time.Second, // 响应 http 返回超时时间
		MaxHeaderBytes: 1 << 20,          // 1 MB. 限制头部最大字节数
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gracefulShutdown := func() error {
		srv.Shutdown(ctx)
		return errors.New("server gracefully stopped")
	}

	// registry
	consulAddr := g.Conf.Get("Registry", "addr")
	reg := registry.NewRegistry(
		registry.Addrs(consulAddr),
	)
	// create service
	service := web.NewService(
		web.Name(MicroName),
		web.Version(Version),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*10),
		web.Server(srv),                  // set custom server
		web.BeforeStop(gracefulShutdown), // graceful shutdown
		web.Registry(reg),
	)
	service.Init()                     // init service
	service.Handle("/", router.Router) // Register Handler
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
