package crawl

import (
	"blog/models/crawl"
	"blog/models/posts"
	_ "strconv"
	"time"

	"github.com/astaxie/beego/httplib"

	"github.com/astaxie/beego"
)

type CrawlController struct {
	beego.Controller
}

var (
	crawl_url  string = "https://www.cnblogs.com/dfh168/p/10720787.html"
	postsInfo  posts.Posts
	crawlModel *crawl.CrawlModel
)

func (this *CrawlController) Crawl() {
	//1.链接redis
	this.ConnectRedis()
	crawl.PutinQueue(crawl_url)
	this.Ctx.WriteString("<br>" + "start of crawl" + "</br>")
	for {
		length := crawl.GetQueueLength()
		if length == 0 {
			break
		}
		crawl_url = crawl.PopfromQueue()
		//判断sUrl是否被访问过
		if crawl.IsVisit(crawl_url) {
			continue
		}
		req := httplib.Get(crawl_url)
		str, err := req.String()
		if err != nil {
			panic(err)
		}
		//提取文章相关信息
		postsInfo.Title = crawlModel.GetContent(str, `<a\s*.*\s*class="postTitle2"\s*.*>(.*)</a>`)
		if postsInfo.Title != "" {
			postsInfo.Id = 0
			postsInfo.Content = crawlModel.GetContent(str, `<div.*class="blogpost-body">([\s|\S]+?)</div>`)
			timeStr := time.Now().Format("2006-01-02 15:04:05")
			postsInfo.Time_create = timeStr
			postsInfo.Time_update = timeStr
			//插入文章
			crawlModel.AddPosts(&postsInfo)
		}
		//	提取该页面的索引连接
		urls := crawlModel.GetUrls(str, `<a.*?href="(https://www.cnblogs.com/.+?)".*>.+</a>`)

		for _, url := range urls {
			crawl.PutinQueue(url)
			this.Ctx.WriteString("<br>" + url + "</br>")
		}
		//crawl_url 应当记录到访问队列set中
		crawl.AddToSet(crawl_url)

		time.Sleep(time.Second * 2)
	}
	this.Ctx.WriteString("<br>" + "end of crawl" + "</br>")
}

//链接redis
func (this *CrawlController) ConnectRedis() {
	//连接redis
	redisHost := beego.AppConfig.String("redishost")
	redisPort := beego.AppConfig.String("redisport")
	redisStr := redisHost + ":" + redisPort
	crawl.ConnectRedis(redisStr)
}

//捕获异常
func (this *CrawlController) CatchError() {
	defer func() {
		if err := recover(); err != nil {
			this.Ctx.WriteString(err.(error).Error())
		}
	}()
}
