package main

import (
	_ "blog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//初始化数据库
func initDB() {
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

func main() {
	orm.Debug = true
	beego.Run()
}
