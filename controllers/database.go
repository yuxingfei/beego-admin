package controllers

import (
	"beego-admin/global/response"
	"beego-admin/services"
)

type DatabaseController struct {
	baseController
}

//显示数据表
func (this *DatabaseController) Table() {
	var databaseService services.DatabaseService
	data, affectRows := databaseService.GetTableStatus()

	this.Data["data"] = data
	this.Data["total"] = affectRows

	this.Layout = "public/base.html"
	this.TplName = "database/table.html"
}

//优化表
func (this *DatabaseController) Optimize() {
	name := this.GetString("name")

	if name == "" {
		response.ErrorWithMessage("请指定要优化的表", this.Ctx)
	}
	var databaseService services.DatabaseService
	ok := databaseService.OptimizeTable(name)
	if ok {
		response.SuccessWithMessage("数据表"+name+"优化成功", this.Ctx)
	} else {
		response.ErrorWithMessage("数据表"+name+"优化失败", this.Ctx)
	}
}

//修复数据表
func (this *DatabaseController) Repair() {
	name := this.GetString("name")

	if name == "" {
		response.ErrorWithMessage("请指定要修复的表", this.Ctx)
	}
	var databaseService services.DatabaseService
	ok := databaseService.OptimizeTable(name)
	if ok {
		response.SuccessWithMessage("数据表"+name+"修复成功", this.Ctx)
	} else {
		response.ErrorWithMessage("数据表"+name+"修复失败", this.Ctx)
	}
}

//查看数据表
func (this *DatabaseController) View() {
	name := this.GetString("name")

	if name == "" {
		response.ErrorWithMessage("请指定要查看的表", this.Ctx)
	}

	var databaseService services.DatabaseService
	data := databaseService.GetFullColumnsFromTable(name)

	this.Data["data"] = data

	this.TplName = "database/view.html"
}
