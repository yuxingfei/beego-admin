package controllers

import (
	"beego-admin/form_validate"
	"beego-admin/global/response"
	"beego-admin/models"
	"beego-admin/services"
	"fmt"
	"github.com/gookit/validate"
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

//添加菜单
func (this *AdminMenuController) Create()  {
	adminMenuForm := form_validate.AdminMenuForm{}

	if err := this.ParseForm(&adminMenuForm); err != nil{
		response.ErrorWithMessage(err.Error(),this.Ctx)
		return
	}

	fmt.Println("adminMenuForm = ",adminMenuForm)


	//数据校验
	v := validate.Struct(adminMenuForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(),this.Ctx)
		return
	}

}