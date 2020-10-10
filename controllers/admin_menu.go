package controllers

import (
	"beego-admin/services"
)

type AdminMenuController struct {
	baseController
}

//控制器，初始化函数，基础自父控制器
func (this *AdminMenuController) NestPrepare() {
	//fmt.Println("AdminMenuController NestPrepare")
}

//菜单首页
func (this *AdminMenuController) Index() {

	var adminMenuService services.AdminMenuService
	this.Data["data"] = adminMenuService.AdminMenuTree()

	this.Layout = "public/base.html"
	this.TplName = "admin_menu/index.html"
}
