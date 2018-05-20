package router

import (
	"github.com/hopehook/micro-project/ws/handler"
	"github.com/julienschmidt/httprouter"
)

var Router = httprouter.New()

// Init init routers
// When use micro web proxy websocket, please add url prefix "stream" to follow url.
// eg:
//    url: ""go.micro.web.stream" ""/test"  ->  "ws://localhost:8082/stream/test"
func Init() {
	Router.GET("/test", handler.Raw(handler.Stream))
}
