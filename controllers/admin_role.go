package controllers

import (
	"beego-admin/form_validate"
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
	data, pagination := adminRoleService.GetPaginateData(admin["per_page"].(int), queryParams)

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
	if err := this.ParseForm(&adminRoleForm); err != nil {
		response.ErrorWithMessage(err.Error(), this.Ctx)
	}

	v := validate.Struct(adminRoleForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), this.Ctx)
	}

	var adminRoleService services.AdminRoleService

	//名称验重
	if adminRoleService.IsExistName(adminRoleForm.Name, 0) {
		response.ErrorWithMessage("名称已存在.", this.Ctx)
	}

	url := global.URL_BACK
	if adminRoleForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	//添加
	insertId := adminRoleService.Create(&adminRoleForm)
	if insertId > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

////菜单管理-角色管理-修改界面
func (this *AdminRoleController) Edit() {
	id, _ := this.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", this.Ctx)
	}

	var (
		adminRoleService services.AdminRoleService
	)

	adminRole := adminRoleService.GetAdminRoleById(id)
	if adminRole == nil {
		response.ErrorWithMessage("Not Found Info By Id.", this.Ctx)
	}

	this.Data["data"] = adminRole

	this.Layout = "public/base.html"
	this.TplName = "admin_role/edit.html"
}

////菜单管理-角色管理-修改
func (this *AdminRoleController) Update() {
	var adminRoleForm form_validate.AdminRoleForm
	if err := this.ParseForm(&adminRoleForm); err != nil {
		response.ErrorWithMessage(err.Error(), this.Ctx)
	}

	//id验证
	if adminRoleForm.Id == 0 {
		response.ErrorWithMessage("id错误.", this.Ctx)
	}

	//字段验证
	v := validate.Struct(adminRoleForm)
	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), this.Ctx)
	}

	var adminRoleService services.AdminRoleService

	//名称验重
	if adminRoleService.IsExistName(adminRoleForm.Name, adminRoleForm.Id) {
		response.ErrorWithMessage("名称已存在.", this.Ctx)
	}

	//修改
	num := adminRoleService.Update(&adminRoleForm)
	if num > 0 {
		response.Success(this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//删除
func (this *AdminRoleController) Del() {
	idStr := this.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		this.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("参数id错误.", this.Ctx)
	}

	var adminRoleService services.AdminRoleService

	noDeletionId := new(models.AdminRole).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionId, idArr)

	if len(noDeletionId) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID为"+strings.Join(utils.IntArrToStringArr(noDeletionId), ",")+"的数据无法删除!", this.Ctx)
	}

	count := adminRoleService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//启用
func (this *AdminRoleController) Enable() {
	idStr := this.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		this.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("请选择启用的角色.", this.Ctx)
	}

	var adminRoleService services.AdminRoleService
	num := adminRoleService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//禁用
func (this *AdminRoleController) Disable() {
	idStr := this.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		this.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("请选择禁用的角色.", this.Ctx)
	}

	var adminRoleService services.AdminRoleService
	num := adminRoleService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//菜单管理-角色管理-角色授权界面
func (this *AdminRoleController) Access() {
	id, _ := this.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", this.Ctx)
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

	this.Data["data"] = data
	this.Data["html"] = html

	this.Layout = "public/base.html"
	this.TplName = "admin_role/access.html"
}

//菜单管理-角色管理-角色授权
func (this *AdminRoleController) AccessOperate()  {
	id,_ := this.GetInt("id",-1)
	if id < 0 {
		response.ErrorWithMessage("Params is Error.",this.Ctx)
	}

	url := make([]string,0)
	this.Ctx.Input.Bind(&url,"url")

	if len(url) == 0{
		response.ErrorWithMessage("请选择授权的菜单",this.Ctx)
	}

	if !utils.InArrayForString(url,"1") {
		response.ErrorWithMessage("首页权限必选",this.Ctx)
	}

	var adminRoleService services.AdminRoleService
	num := adminRoleService.StoreAccess(id,url)
	if num > 0{
		response.Success(this.Ctx)
	}else {
		response.Error(this.Ctx)
	}

}