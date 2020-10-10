package controllers

import (
	"beego-admin/models"
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

//添加菜单界面
func (this *AdminMenuController) Add() {

	var adminTreeService services.AdminTreeService
	parentId, _ := this.GetInt("parent_id", 0)
	parents := adminTreeService.Menu(parentId, 0)

	this.Data["parents"] = parents
	this.Data["log_method"] = new(models.AdminMenu).GetLogMethod()

	this.Layout = "public/base.html"
	this.TplName = "admin_menu/add.html"
}
