package main

import (
	_ "tally/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Info("开始启动")
	// beego.Config.
	beego.Run("0.0.0.0:80")
}
