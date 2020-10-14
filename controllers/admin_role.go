package controllers

import (
	"beego-admin/form_validate"
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/services"
	"github.com/gookit/validate"
)

type AdminRoleController struct {
	baseController
}

//控制器，初始化函数，基础自父控制器
func (this *AdminRoleController) NestPrepare() {
	//fmt.Println("AdminRoleController NestPrepare")
}

//角色管理首页
func (this *AdminRoleController) Index() {
	var adminRoleService services.AdminRoleService
	data, pagination := adminRoleService.GetAllData(admin["per_page"].(int), queryParams)

	this.Data["data"] = data
	this.Data["paginate"] = pagination
	this.Layout = "public/base.html"
	this.TplName = "admin_role/index.html"
}

//角色管理-添加界面
func (this *AdminRoleController) Add() {
	this.Layout = "public/base.html"
	this.TplName = "admin_role/add.html"
}

//角色管理-添加角色
func (this *AdminRoleController) Create() {
	var adminRoleForm form_validate.AdminRoleForm
	if err := this.ParseForm(&adminRoleForm); err != nil{
		response.ErrorWithMessage(err.Error(),this.Ctx)
		return
	}

	v := validate.Struct(adminRoleForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(),this.Ctx)
		return
	}

	var adminRoleService services.AdminRoleService

	//名称验重
	if adminRoleService.IsExistName(adminRoleForm.Name,0) {
		response.ErrorWithMessage("名称已存在.",this.Ctx)
		return
	}

	url := global.URL_BACK
	if adminRoleForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	//添加
	insertId := adminRoleService.Create(&adminRoleForm)
	if insertId > 0{
		response.SuccessWithMessageAndUrl("添加成功",url,this.Ctx)
	}else {
		response.Error(this.Ctx)
	}
	return
}
