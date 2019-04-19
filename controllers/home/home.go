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
	this.Model = &posts.Posts{
		Id:       1,
		Admin_id: 2,
	}
}

func (this *HomeController) Index() {
	//this.Ctx.WriteString("Hello Index()!")
	// this.Data["json"] = &posts.Posts{
	// 	Id:       1,
	// 	Admin_id: 2,
	// }
	// this.ServeJSON()
	posts := &posts.Posts{}
	posts.GetPosts()
}
