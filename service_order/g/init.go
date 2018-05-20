package g

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hopehook/micro-project/common/go_lib"
)

// InitGlobal init glabal varibles
func InitGlobal() {
	// handler Logs
	{
		section := "Logs"
		path := fmt.Sprintf("%s/log/app.%d.log", Path, os.Getpid())
		fmt.Println(path)
		separate := Conf.Get(section, "separate")
		separates := strings.Split(separate, ",")
		level, _ := strconv.Atoi(Conf.Get(section, "level"))
		maxdays, _ := strconv.Atoi(Conf.Get(section, "maxdays"))
		Logger = go_lib.InitLogger(path, separates, level, maxdays)
	}
}
