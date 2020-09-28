package controllers

import "beego-admin/services/admin_user_service"

type AdminUserController struct {
	baseController
}

//控制器，初始化函数，基础自父控制器
func (this *AdminUserController) NestPrepare() {
	//fmt.Println("AdminUserController NestPrepare")
}

func (this *AdminUserController)Index()  {

	this.Data["data"] = admin_user_service.GetAllData()

	this.Layout = "public/base.html"
	this.TplName = "admin_user/index.html"
}
