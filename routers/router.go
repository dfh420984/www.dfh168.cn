package routers

import (
	"blog/controllers/home"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &home.HomeController{}, "get:Index")
	beego.Router("/home/index", &home.HomeController{}, "get:Index")
	beego.Router("/home/total", &home.HomeController{}, "get:Total")
}
