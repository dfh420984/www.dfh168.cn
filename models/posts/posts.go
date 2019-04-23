package posts

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	page     int = 1
	pageSize int = 1
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

//初始化数据库
func initDB() {
	//开启调试模式
	//orm.Debug = true
	//1.驱动类型
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//2.数据库配置
	dbHost := beego.AppConfig.String("mysqlhost")
	dbPort := beego.AppConfig.String("mysqport")
	dbDataBase := beego.AppConfig.String("mysqldb")
	dbUserName := beego.AppConfig.String("mysqluser")
	dbPwd := beego.AppConfig.String("mysqlpass")
	//3.数据库连接
	conn := dbUserName + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDataBase + "?charset=utf8"
	//4.注册默认数据库
	orm.RegisterDataBase("default", "mysql", conn, 30, 30)
}

func init() {
	initDB()
}

//获取帖子列表
func (this *Posts) GetPosts(search_where map[string]interface{}) (res Result) {
	defer func() {
		if err := recover(); err != nil {
			res.Code = 1
			res.Message = "获取帖子列表失败:"
			res.Data = nil
		}
	}()
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
	if _, ok := search_where["keyword"]; ok {
		info := fmt.Sprintf("AND (p.content LIKE '%%%s%%' OR p.title LIKE '%%%s%%' OR p.slug LIKE '%%%s%%') ", search_where["keyword"], search_where["keyword"], search_where["keyword"])
		where += info
	}
	if _, ok := search_where["slug"]; ok {
		info := fmt.Sprintf("AND  slug LIKE '%%%s%%' ", search_where["keyword"])
		where += info
	}
	if _, ok := search_where["cat_id"]; ok {
		info := fmt.Sprintf("AND cat_id = %d ", search_where["cat_id"])
		where += info
	}
	limit := strconv.Itoa(offset) + "," + strconv.Itoa(pageSize)
	sql = fmt.Sprintf(sql, where, limit)
	num, err := o.Raw(sql).Values(&maps)
	fmt.Printf("err=%v,num=%v\n", err, num)
	if err == nil && num > 0 {
		//info := fmt.Sprintf("num=%v,posts=%v \n", num, maps)
		res.Code = 0
		res.Message = "ok"
		res.Data = maps
	} else {
		res.Code = 1
		res.Message = "获取帖子列表失败:"
		res.Data = maps
		fmt.Printf("res=%v\n", res)
	}
	return res
}

//获取最新文章

//获取帖子总数
func (this *Posts) GetPostsTotal() (res Result) {
	o := orm.NewOrm()
	var maps []orm.Params
	sql := `SELECT count(*) as num
		FROM posts
		WHERE %s;`
	where := "status =1 "
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
	} else {
		res.Code = 1
		res.Message = "获取帖子列表失败:" + err.Error()
		res.Data = nil
	}
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
	} else {
		res.Code = 1
		res.Message = "获取帖子评论数量失败:" + err.Error()
		res.Data = nil
	}
	return res
}
