package controllers

import (
	"beego-admin/global/response"
	"beego-admin/services"
)

// DatabaseController struct
type DatabaseController struct {
	baseController
}

// Table 显示数据表
func (dc *DatabaseController) Table() {
	var databaseService services.DatabaseService
	data, affectRows := databaseService.GetTableStatus()

	dc.Data["data"] = data
	dc.Data["total"] = affectRows

	dc.Layout = "public/base.html"
	dc.TplName = "database/table.html"
}

// Optimize 优化表
func (dc *DatabaseController) Optimize() {
	name := dc.GetString("name")

	if name == "" {
		response.ErrorWithMessage("请指定要优化的表", dc.Ctx)
	}
	var databaseService services.DatabaseService
	ok := databaseService.OptimizeTable(name)
	if ok {
		response.SuccessWithMessage("数据表"+name+"优化成功", dc.Ctx)
	} else {
		response.ErrorWithMessage("数据表"+name+"优化失败", dc.Ctx)
	}
}

// Repair 修复数据表
func (dc *DatabaseController) Repair() {
	name := dc.GetString("name")

	if name == "" {
		response.ErrorWithMessage("请指定要修复的表", dc.Ctx)
	}
	var databaseService services.DatabaseService
	ok := databaseService.OptimizeTable(name)
	if ok {
		response.SuccessWithMessage("数据表"+name+"修复成功", dc.Ctx)
	} else {
		response.ErrorWithMessage("数据表"+name+"修复失败", dc.Ctx)
	}
}

// View 查看数据表
func (dc *DatabaseController) View() {
	name := dc.GetString("name")

	if name == "" {
		response.ErrorWithMessage("请指定要查看的表", dc.Ctx)
	}

	var databaseService services.DatabaseService
	data := databaseService.GetFullColumnsFromTable(name)

	dc.Data["data"] = data

	dc.TplName = "database/view.html"
}
