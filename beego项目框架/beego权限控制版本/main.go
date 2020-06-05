package main

import (
	_ "./models"
	_ "./routers"

	"github.com/astaxie/beego"
	"github.com/wonderivan/logger"
)

func main() {

	// 处理日志等级
	logger.SetLogger("./conf/log.json")

	beego.SetStaticPath("admin/pages/", "./pages")

	beego.Run()
}
