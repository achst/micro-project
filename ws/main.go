package main

import (
	"log"
	"time"

	"strings"

	"github.com/hopehook/micro-project/common/go_client"
	"github.com/hopehook/micro-project/common/go_lib"
	"github.com/hopehook/micro-project/common/go_util"
	"github.com/hopehook/micro-project/ws/g"
	"github.com/hopehook/micro-project/ws/router"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-web"
)

const DefaultConf = "my.cnf"
const MicroName = "go.micro.web.ws"
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
		web.Registry(reg),
	)
	service.Init()                     // init service
	service.Handle("/", router.Router) // Register Handler
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
