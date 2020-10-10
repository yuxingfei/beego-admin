package controllers

import (
	"beego-admin/services"
)

type AdminLogController struct {
	baseController
}

//控制器，初始化函数，基础自父控制器
func (this *AdminLogController) NestPrepare() {
	//fmt.Println("AdminLogController NestPrepare")
}

func (this *AdminLogController) Index() {

	var (
		adminLogService  services.AdminLogService
		adminUserService services.AdminUserService
	)
	data, pagination := adminLogService.GetAllData(admin["per_page"].(int), queryParams)

	this.Data["admin_user_list"] = adminUserService.GetAllAdminUser()

	this.Data["data"] = data
	this.Data["paginate"] = pagination
	this.Layout = "public/base.html"
	this.TplName = "admin_log/index.html"
}
