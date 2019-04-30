package crawl

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type CrawlController struct {
	beego.Controller
}

var (
	crawl_url string = "https://www.cnblogs.com/"
)

func (this *CrawlController) Crawl() {
	req := httplib.Get(crawl_url)
	str, err := req.String()
	if err != nil {
		panic(err)
	}
	this.Ctx.WriteString(str)
}

//捕获异常
func (this *CrawlController) CatchError() {
	defer func() {
		if err := recover(); err != nil {
			this.Ctx.WriteString(err.(error).Error())
		}
	}()
}
