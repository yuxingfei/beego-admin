package main

import (
	"beego-admin/initialize/conf"
	"beego-admin/initialize/mysql"
	"beego-admin/initialize/session"
	"beego-admin/routers"
	"beego-admin/utils/template"
	beego "github.com/beego/beego/v2/adapter"
)

func main() {
	//初始化配置文件
	conf.InitConfig()
	//初始化数据库
	mysql.InitDb()
	//初始化session
	session.InitSession()
	//初始化router
	routers.InitRouter()
	//初始化自定义模版
	template.InitTtempate()
	//输出文件名和行号
	beego.SetLogFuncCall(true)

	//启动beego
	beego.Run()
}
