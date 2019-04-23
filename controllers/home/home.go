package home

import (
	"blog/models/posts"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
	Model *posts.Posts
}

//获取参数
func (this *HomeController) Prepare() {

}

func (this *HomeController) search_where() map[string]interface{} {
	search_where := make(map[string]interface{})
	page, _ := this.GetInt("page", 1)
	search_where["page"] = page
	cat_id, _ := this.GetInt("cat_id", 0)
	if cat_id > 0 {
		search_where["cat_id"] = cat_id
	}
	slug := this.GetString("slug")
	if slug != "" {
		search_where["slug"] = slug
	}
	keyword := this.GetString("keyword")
	if keyword != "" {
		search_where["keyword"] = keyword
	}
	time_create := this.GetString("time_create")
	if time_create != "" {
		search_where["time_create"] = time_create
	}
	return search_where
}

//获取列表页
func (this *HomeController) Index() {
	//this.Ctx.WriteString("Hello Index()!")
	//获取查询条件
	search_where := this.search_where()
	//获取帖子列表
	posts := &posts.Posts{}
	res := posts.GetPosts(search_where)
	//获取帖子评论数量
	if res.Code == 0 {
		com := posts.GetPostsComment()
		if com.Code == 0 {
			for i, post := range res.Data {
				res.Data[i]["com_num"] = 0
				for _, com := range com.Data {
					if post["id"] == com["posts_id"] {
						res.Data[i]["com_num"] = com["num"]
					}
				}
			}
		}
	}
	this.Data["json"] = &res
	this.ServeJSON()
}

//获取帖子数量
func (this *HomeController) Total() {
	posts := &posts.Posts{}
	//获取帖子数量
	res := posts.GetPostsTotal()
	this.Data["json"] = &res
	this.ServeJSON()
}
