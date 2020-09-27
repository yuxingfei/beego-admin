package main

import (
	_ "beego-admin/initialize/mysql"
	_ "beego-admin/initialize/session"
	_ "beego-admin/routers"
	"github.com/astaxie/beego"
)

func main() {
	//输出文件名和行号
	beego.SetLogFuncCall(true)

	//启动beego
	beego.Run()
}

