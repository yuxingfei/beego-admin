package controllers

import (
	"beego-admin/services"
)

// AdminLogController struct.
type AdminLogController struct {
	baseController
}

// Index index.
func (alc *AdminLogController) Index() {
	var (
		adminLogService  services.AdminLogService
		adminUserService services.AdminUserService
	)
	data, pagination := adminLogService.GetPaginateData(alc.Option["per_page"].(int), alc.QueryParams)

	alc.Data["admin_user_list"] = adminUserService.GetAllAdminUser()

	alc.Data["data"] = data
	alc.Data["paginate"] = pagination
	alc.Layout = "public/base.html"
	alc.TplName = "admin_log/index.html"
}
