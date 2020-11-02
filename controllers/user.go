package controllers

import (
	"beego-admin/services"
)

type UserController struct {
	baseController
}

//用户等级 列表页
func (this *UserController) Index() {
	var userService services.UserService
	data, pagination := userService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	this.Data["data"] = data
	this.Data["paginate"] = pagination

	this.Layout = "public/base.html"
	this.TplName = "user/index.html"
}
