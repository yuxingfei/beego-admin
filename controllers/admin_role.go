package controllers

import (
	"beego-admin/formvalidate"
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/models"
	"beego-admin/services"
	"beego-admin/utils"
	"github.com/adam-hanna/arrayOperations"
	"github.com/astaxie/beego/orm"
	"github.com/gookit/validate"
	"strconv"
	"strings"
)

// AdminRoleController struct.
type AdminRoleController struct {
	baseController
}

// Index 角色管理首页.
func (arc *AdminRoleController) Index() {
	var adminRoleService services.AdminRoleService
	data, pagination := adminRoleService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	arc.Data["data"] = data
	arc.Data["paginate"] = pagination
	arc.Layout = "public/base.html"
	arc.TplName = "admin_role/index.html"
}

// Add 角色管理-添加界面.
func (arc *AdminRoleController) Add() {
	arc.Layout = "public/base.html"
	arc.TplName = "admin_role/add.html"
}

// Create 角色管理-添加角色.
func (arc *AdminRoleController) Create() {
	var adminRoleForm formvalidate.AdminRoleForm
	if err := arc.ParseForm(&adminRoleForm); err != nil {
		response.ErrorWithMessage(err.Error(), arc.Ctx)
	}

	v := validate.Struct(adminRoleForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), arc.Ctx)
	}

	var adminRoleService services.AdminRoleService

	//名称验重
	if adminRoleService.IsExistName(adminRoleForm.Name, 0) {
		response.ErrorWithMessage("名称已存在.", arc.Ctx)
	}

	url := global.URL_BACK
	if adminRoleForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	//添加
	insertID := adminRoleService.Create(&adminRoleForm)
	if insertID > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, arc.Ctx)
	} else {
		response.Error(arc.Ctx)
	}
}

// Edit 菜单管理-角色管理-修改界面.
func (arc *AdminRoleController) Edit() {
	id, _ := arc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", arc.Ctx)
	}

	var (
		adminRoleService services.AdminRoleService
	)

	adminRole := adminRoleService.GetAdminRoleById(id)
	if adminRole == nil {
		response.ErrorWithMessage("Not Found Info By Id.", arc.Ctx)
	}

	arc.Data["data"] = adminRole

	arc.Layout = "public/base.html"
	arc.TplName = "admin_role/edit.html"
}

// Update 菜单管理-角色管理-修改.
func (arc *AdminRoleController) Update() {
	var adminRoleForm formvalidate.AdminRoleForm
	if err := arc.ParseForm(&adminRoleForm); err != nil {
		response.ErrorWithMessage(err.Error(), arc.Ctx)
	}

	//id验证
	if adminRoleForm.Id == 0 {
		response.ErrorWithMessage("id错误.", arc.Ctx)
	}

	//字段验证
	v := validate.Struct(adminRoleForm)
	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), arc.Ctx)
	}

	var adminRoleService services.AdminRoleService

	//名称验重
	if adminRoleService.IsExistName(adminRoleForm.Name, adminRoleForm.Id) {
		response.ErrorWithMessage("名称已存在.", arc.Ctx)
	}

	//修改
	num := adminRoleService.Update(&adminRoleForm)
	if num > 0 {
		response.Success(arc.Ctx)
	} else {
		response.Error(arc.Ctx)
	}
}

// Del 删除.
func (arc *AdminRoleController) Del() {
	idStr := arc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		arc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("参数id错误.", arc.Ctx)
	}

	var adminRoleService services.AdminRoleService

	noDeletionID := new(models.AdminRole).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID为"+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"的数据无法删除!", arc.Ctx)
	}

	count := adminRoleService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, arc.Ctx)
	} else {
		response.Error(arc.Ctx)
	}
}

// Enable 启用.
func (arc *AdminRoleController) Enable() {
	idStr := arc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		arc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("请选择启用的角色.", arc.Ctx)
	}

	var adminRoleService services.AdminRoleService
	num := adminRoleService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, arc.Ctx)
	} else {
		response.Error(arc.Ctx)
	}
}

// Disable 禁用.
func (arc *AdminRoleController) Disable() {
	idStr := arc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		arc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("请选择禁用的角色.", arc.Ctx)
	}

	var adminRoleService services.AdminRoleService
	num := adminRoleService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, arc.Ctx)
	} else {
		response.Error(arc.Ctx)
	}
}

// Access 菜单管理-角色管理-角色授权界面.
func (arc *AdminRoleController) Access() {
	id, _ := arc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", arc.Ctx)
	}

	var (
		adminRoleService services.AdminRoleService
		adminMenuService services.AdminMenuService
		adminTreeService services.AdminTreeService
	)

	data := adminRoleService.GetAdminRoleById(id)
	menu := adminMenuService.AllMenu()

	menuMap := make(map[int]orm.Params)

	for _, adminMenu := range menu {
		id := adminMenu.Id
		if menuMap[id] == nil {
			menuMap[id] = make(orm.Params)
		}
		menuMap[id]["Id"] = id
		menuMap[id]["ParentId"] = adminMenu.ParentId
		menuMap[id]["Name"] = adminMenu.Name
		menuMap[id]["Url"] = adminMenu.Url
		menuMap[id]["Icon"] = adminMenu.Icon
		menuMap[id]["IsShow"] = adminMenu.IsShow
		menuMap[id]["SortId"] = adminMenu.SortId
		menuMap[id]["LogMethod"] = adminMenu.LogMethod
	}

	html := adminTreeService.AuthorizeHtml(menuMap, strings.Split(data.Url, ","))

	arc.Data["data"] = data
	arc.Data["html"] = html

	arc.Layout = "public/base.html"
	arc.TplName = "admin_role/access.html"
}

// AccessOperate 菜单管理-角色管理-角色授权.
func (arc *AdminRoleController) AccessOperate() {
	id, _ := arc.GetInt("id", -1)
	if id < 0 {
		response.ErrorWithMessage("Params is Error.", arc.Ctx)
	}

	url := make([]string, 0)
	arc.Ctx.Input.Bind(&url, "url")

	if len(url) == 0 {
		response.ErrorWithMessage("请选择授权的菜单", arc.Ctx)
	}

	if !utils.InArrayForString(url, "1") {
		response.ErrorWithMessage("首页权限必选", arc.Ctx)
	}

	var adminRoleService services.AdminRoleService
	num := adminRoleService.StoreAccess(id, url)
	if num > 0 {
		response.Success(arc.Ctx)
	} else {
		response.Error(arc.Ctx)
	}

}
