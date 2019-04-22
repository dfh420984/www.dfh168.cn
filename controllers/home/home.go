package home

import (
	"blog/models/posts"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
	Model *posts.Posts
}

func (this *HomeController) Prepare() {

}

func (this *HomeController) Index() {
	//this.Ctx.WriteString("Hello Index()!")
	posts := &posts.Posts{}
	res := posts.GetPosts()
	//获取帖子数量
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
