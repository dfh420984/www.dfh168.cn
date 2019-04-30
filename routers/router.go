package routers

import (
	"blog/controllers/crawl"
	"blog/controllers/home"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &home.HomeController{}, "get:Index")
	beego.Router("/home/index", &home.HomeController{}, "get:Index")
	beego.Router("/home/total", &home.HomeController{}, "get:Total")
	beego.Router("/home/new", &home.HomeController{}, "get:New")
	beego.Router("/home/archive", &home.HomeController{}, "get:Archive")
	beego.Router("/home/category", &home.HomeController{}, "get:Category")
	beego.Router("/home/tag", &home.HomeController{}, "get:Tag")
	beego.Router("/crawl", &crawl.CrawlController{}, "get:Crawl")
}
