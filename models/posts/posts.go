package posts

import (
	"github.com/astaxie/beego"
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

//初始化数据库
func initDB() {
	//开启调试模式
	orm.Debug = true
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

func (this *Posts) GetPosts() (res Result) {
	o := orm.NewOrm()
	var maps []orm.Params
	offset := (page - 1) * pageSize
	num, err := o.Raw("SELECT * FROM posts limit ?,?", offset, pageSize).Values(&maps)
	if err == nil && num > 0 {
		//info := fmt.Sprintf("num=%v,posts=%v \n", num, maps)
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