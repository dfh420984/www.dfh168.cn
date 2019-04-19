package posts

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	page     int = 1
	pageSize int = 20
)

type Posts struct {
	Id          int    `json:"id"`
	Admin_id    int    `json:"admin_id"`
	Cat_id      int    `json:"cat_id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Image       string `json:"image"`
	Content     string `json:"content"`
	Status      int    `json:"status"`
	Time_create string `json:"time_create"`
	Time_update string `json:"time_update"`
}

func (this *Posts) GetPosts() {
	o := orm.NewOrm()
	var maps []orm.Params
	offset := (page - 1) * pageSize
	num, err := o.Raw("SELECT * FROM posts limit ?,?", offset, pageSize).Values(&maps)
	if err == nil {
		fmt.Printf("num=%v,posts=%v \n", num, maps)
	}
}
