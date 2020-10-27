package controllers

import (
	"beego-admin/services"
)

type SettingController struct {
	baseController
}

//设置
func (this *SettingController) Admin() {
	this.show(1)
}

//展示单个配置信息
func (this *SettingController) show(id int) {
	var settingService services.SettingService
	data := settingService.Show(id)

	this.Data["data_config"] = data
	this.Layout = "public/base.html"
	this.TplName = "setting/show.html"
}
