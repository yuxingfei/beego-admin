package controllers

import (
	"beego-admin/services"
	"fmt"
)

type AdminLogController struct {
	baseController
}


//控制器，初始化函数，基础自父控制器
func (this *AdminLogController) NestPrepare() {
	//fmt.Println("AdminLogController NestPrepare")
}

func (this *AdminLogController)Index()  {

	fmt.Println("this.Input() = ",this.Input())
	fmt.Println("this.Ctx.Request.RequestURI = ",this.Ctx.Request.RequestURI)
	fmt.Println("this.Ctx.Request.URL = ",this.Ctx.Request.URL)
	fmt.Println("this.Ctx.Input.Data() = ",this.Ctx.Input.Data())
	fmt.Println("this.Ctx.Input.Params() = ",this.Ctx.Input.Params())
	var adminLogService services.AdminLogService
	data , pagination := adminLogService.GetAllData()
	this.Data["data"] = data
	this.Data["paginate"] =pagination

	this.Layout = "public/base.html"
	this.TplName = "admin_log/index.html"
}
