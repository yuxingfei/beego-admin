package main

import (
	"beego-admin/initialize/mysql"
	_ "beego-admin/routers"
	"github.com/astaxie/beego"
)

func main() {
	//输出文件名和行号
	beego.SetLogFuncCall(true)

	//mysql数据库注册
	mysql.Init()

	//启动beego
	beego.Run()
}

