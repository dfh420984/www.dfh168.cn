package posts

import (
	"fmt"
	"strconv"

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

type Result struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    []orm.Params `json:"data"`
}

//获取帖子列表
func (this *Posts) GetPosts(search_where map[string]interface{}) (res Result) {
	o := orm.NewOrm()
	var maps []orm.Params
	page = search_where["page"].(int)
	offset := (page - 1) * pageSize
	sql := `SELECT p.*, a.email,a.mobile,a.nick_name,a.alias,
		c.content as cat_content,c.slug as cat_slug 
		FROM posts AS p 
		LEFT JOIN admin AS a ON p.admin_id = a.id 
		LEFT JOIN category AS c ON p.cat_id = c.id 
		WHERE %s 
		ORDER BY p.view_num desc LIMIT %s;`
	where := "p.status =1 "
	if _, ok := search_where["id"]; ok {
		info := fmt.Sprintf("AND p.id = %d ", search_where["id"])
		where += info
	}
	if _, ok := search_where["keyword"]; ok {
		info := fmt.Sprintf("AND (p.content LIKE '%%%s%%' OR p.title LIKE '%%%s%%' OR p.slug LIKE '%%%s%%') ", search_where["keyword"], search_where["keyword"], search_where["keyword"])
		where += info
	}
	if _, ok := search_where["slug"]; ok {
		info := fmt.Sprintf("AND  p.slug LIKE '%%%s%%' ", search_where["slug"])
		where += info
	}
	if _, ok := search_where["cat_id"]; ok {
		info := fmt.Sprintf("AND p.cat_id = %d ", search_where["cat_id"])
		where += info
	}
	if _, ok := search_where["time_create"]; ok {
		info := fmt.Sprintf("AND DATE_FORMAT(p.time_create,'%%Y%%m') = %s ", search_where["time_create"])
		where += info
	}
	limit := strconv.Itoa(offset) + "," + strconv.Itoa(pageSize)
	sql = fmt.Sprintf(sql, where, limit)
	num, err := o.Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		//info := fmt.Sprintf("num=%v,posts=%v \n", num, maps)
		res.Code = 0
		res.Message = "ok"
		res.Data = maps
	} else if err == nil && num == 0 {
		res.Code = 0
		res.Message = "帖子数量为0"
		res.Data = nil
	} else {
		res.Code = 1
		res.Message = "获取帖子列表失败:"
		res.Data = maps
		fmt.Printf("res=%v\n", res)
	}
	this.CatchError(res)
	return res
}

//获取最新文章
func (this *Posts) GetNewPosts(search_where map[string]interface{}) (res Result) {
	o := orm.NewOrm()
	var maps []orm.Params
	sql := `SELECT id,title FROM posts WHERE %s ORDER BY time_create DESC LIMIT 0,10;`
	where := "status =1 "
	if _, ok := search_where["keyword"]; ok {
		info := fmt.Sprintf("AND (content LIKE '%%%s%%' OR title LIKE '%%%s%%' OR slug LIKE '%%%s%%') ", search_where["keyword"], search_where["keyword"], search_where["keyword"])
		where += info
	}
	if _, ok := search_where["slug"]; ok {
		info := fmt.Sprintf("AND  slug LIKE '%%%s%%' ", search_where["slug"])
		where += info
	}
	if _, ok := search_where["cat_id"]; ok {
		info := fmt.Sprintf("AND cat_id = %d ", search_where["cat_id"])
		where += info
	}
	if _, ok := search_where["time_create"]; ok {
		info := fmt.Sprintf("AND DATE_FORMAT(time_create,'%%Y%%m') = %s ", search_where["time_create"])
		where += info
	}
	sql = fmt.Sprintf(sql, where)
	num, err := o.Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		res.Code = 0
		res.Message = "ok"
		res.Data = maps
	} else if err == nil && num == 0 {
		res.Code = 0
		res.Message = "最新文章数量为0"
		res.Data = nil
	} else {
		res.Code = 1
		res.Message = "获取最新文章失败:" + err.Error()
		res.Data = nil
	}
	this.CatchError(res)
	return res
}

//获取帖子总数
func (this *Posts) GetPostsTotal(search_where map[string]interface{}) (res Result) {
	o := orm.NewOrm()
	var maps []orm.Params
	sql := `SELECT count(*) as num
		FROM posts
		WHERE %s;`
	where := "status =1 "
	if _, ok := search_where["keyword"]; ok {
		info := fmt.Sprintf("AND (content LIKE '%%%s%%' OR title LIKE '%%%s%%' OR slug LIKE '%%%s%%') ", search_where["keyword"], search_where["keyword"], search_where["keyword"])
		where += info
	}
	if _, ok := search_where["slug"]; ok {
		info := fmt.Sprintf("AND  slug LIKE '%%%s%%' ", search_where["slug"])
		where += info
	}
	if _, ok := search_where["cat_id"]; ok {
		info := fmt.Sprintf("AND cat_id = %d ", search_where["cat_id"])
		where += info
	}
	if _, ok := search_where["time_create"]; ok {
		info := fmt.Sprintf("AND DATE_FORMAT(time_create,'%%Y%%m') = %s ", search_where["time_create"])
		where += info
	}
	sql = fmt.Sprintf(sql, where)
	num, err := o.Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		for i, v := range maps {
			num, ok := v["num"].(string)
			total := 0
			all := 0
			if ok {
				total, _ = strconv.Atoi(num)
			}
			if total%pageSize == 0 {
				all = total / pageSize
			} else {
				all = (total / pageSize) + 1
			}
			maps[i]["num"] = all
		}
		res.Code = 0
		res.Message = "ok"
		res.Data = maps
	} else if err == nil && num == 0 {
		res.Code = 0
		res.Message = "帖子数量为0"
		res.Data = nil
	} else {
		res.Code = 1
		res.Message = "获取帖子总数失败:" + err.Error()
		res.Data = nil
	}
	this.CatchError(res)
	return res
}

//获取帖子评论数量
func (this *Posts) GetPostsComment() (res Result) {
	o := orm.NewOrm()
	var maps []orm.Params
	//offset := (page - 1) * pageSize
	sql := `SELECT posts_id,COUNT(*) as num FROM comments
		WHERE %s`
	where := "status =1 GROUP BY  posts_id "
	sql = fmt.Sprintf(sql, where)
	num, err := o.Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		//info := fmt.Sprintf("num=%v,posts=%v \n", num, maps)
		res.Code = 0
		res.Message = "ok"
		res.Data = maps
	} else if err == nil && num == 0 {
		res.Code = 0
		res.Message = "帖子评论数量为0"
		res.Data = nil
	} else {
		res.Code = 1
		res.Message = "获取帖子评论数量失败:" + err.Error()
		res.Data = nil
	}
	this.CatchError(res)
	return res
}

//获取归档
func (this *Posts) GetArchive() (res Result) {
	o := orm.NewOrm()
	var maps []orm.Params
	sql := `SELECT DATE_FORMAT(time_create,'%Y%m') as time_create FROM posts GROUP BY time_create`
	num, err := o.Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		res.Code = 0
		res.Message = "ok"
		res.Data = maps
	} else if err == nil && num == 0 {
		res.Code = 0
		res.Message = "最新文章数量为0"
		res.Data = nil
	} else {
		res.Code = 1
		res.Message = "获取最新文章失败:" + err.Error()
		res.Data = nil
	}
	this.CatchError(res)
	return res
}

//获取分类
func (this *Posts) GetCategory() (res Result) {
	o := orm.NewOrm()
	var maps []orm.Params
	sql := `SELECT c.id AS c_id,c.content AS c_content,COUNT(p.cat_id) AS num FROM category AS c LEFT JOIN posts AS p ON p.cat_id = c.id GROUP BY c_id`
	num, err := o.Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		res.Code = 0
		res.Message = "ok"
		res.Data = maps
	} else if err == nil && num == 0 {
		res.Code = 0
		res.Message = "最新分类数量为0"
		res.Data = nil
	} else {
		res.Code = 1
		res.Message = "获取分类失败:" + err.Error()
		res.Data = nil
	}
	this.CatchError(res)
	return res
}

//获取标签
func (this *Posts) GetTag() (res Result) {
	o := orm.NewOrm()
	var maps []orm.Params
	sql := `SELECT DISTINCT(slug) as slug FROM posts WHERE status = 1`
	num, err := o.Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		res.Code = 0
		res.Message = "ok"
		res.Data = maps
	} else if err == nil && num == 0 {
		res.Code = 0
		res.Message = "最新标签数量为0"
		res.Data = nil
	} else {
		res.Code = 1
		res.Message = "获取标签失败:" + err.Error()
		res.Data = nil
	}
	this.CatchError(res)
	return res
}

//捕获异常
func (this *Posts) CatchError(res Result) {
	defer func() {
		if err := recover(); err != nil {
			res.Code = 1
			res.Message = "获取帖子列表失败:" + err.(error).Error()
			res.Data = nil
		}
	}()
}
