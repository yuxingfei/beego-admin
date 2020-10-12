package controllers

import (
	"beego-admin/form_validate"
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/models"
	"beego-admin/services"
	"fmt"
	"github.com/gookit/validate"
	"strings"
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
	var adminMenuService services.AdminMenuService
	adminMenuForm := form_validate.AdminMenuForm{}

	if err := this.ParseForm(&adminMenuForm); err != nil{
		response.ErrorWithMessage(err.Error(),this.Ctx)
		return
	}

	//去除Url前后两侧的空格
	if adminMenuForm.Url != ""{
		adminMenuForm.Url = strings.TrimSpace(adminMenuForm.Url)
	}

	fmt.Println("adminMenuForm = ",adminMenuForm)

	//数据校验
	v := validate.Struct(adminMenuForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(),this.Ctx)
		return
	}

	//添加之前url验重
	if adminMenuService.IsUrlUnique(adminMenuForm.Url){
		response.ErrorWithMessage("url已经存在.",this.Ctx)
		return
	}

	//创建
	_,err := adminMenuService.Create(&adminMenuForm)
	if err != nil{
		response.Error(this.Ctx)
		return
	}

	url := global.URL_BACK
	if adminMenuForm.IsCreate == 1{
		url = global.URL_RELOAD
	}

	response.SuccessWithMessageAndUrl("添加成功",url,this.Ctx)
	return
}