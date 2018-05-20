package g

import (
	"github.com/hopehook/micro-project/common/go_lib"

	"github.com/astaxie/beego/logs"
)

// Path is global path
var Path string

// Conf is to read demo.conf
var Conf *go_lib.Config

// Logger is global log fd
var Logger *logs.BeeLogger
