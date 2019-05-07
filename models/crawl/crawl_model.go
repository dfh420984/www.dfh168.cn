package crawl

import (
	"blog/models/posts"
	"regexp"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type CrawlModel struct {
}

var (
	db orm.Ormer
)

//初始化数据库
func (this *CrawlModel) initDB() {
	db = orm.NewOrm()
	db.Using("default") // 默认使用 default，你可以指定为其他数据库
}

func (this *CrawlModel) init() {
	this.initDB()
}

func (this *CrawlModel) AddPosts(postsInfo *posts.Posts) (int64, error) {
	this.init()
	id, err := db.Insert(postsInfo)
	return id, err
}

func (this *CrawlModel) GetContent(content string, rule string) string {
	if content == "" {
		return ""
	}
	reg := regexp.MustCompile(rule)
	result := reg.FindAllStringSubmatch(content, -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

func (this *CrawlModel) GetUrls(content string, rule string) []string {
	reg := regexp.MustCompile(rule)
	result := reg.FindAllStringSubmatch(content, -1)

	var postSets []string
	for _, v := range result {
		postSets = append(postSets, v[1])
	}
	return postSets
}
