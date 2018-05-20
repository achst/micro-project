package router

import (
	"github.com/hopehook/micro-project/api/handler"
	"github.com/julienschmidt/httprouter"
)

var Router = httprouter.New()

// Init init routers
// Url prefix must keep consistent with MicroName's last name,
// then micro api will proxy http requests to this server.
// eg: url: "/api/*"  <->  "go.micro.api.api"
//     url: "/auth/*" <->  "go.micro.api.auth"
func Init() {
	// normal api
	Router.GET("/api/order/list", handler.Raw(handler.GetOrderList))
	// async api
	Router.POST("/api/order/update", handler.Raw(handler.UpdateOrder))
}
