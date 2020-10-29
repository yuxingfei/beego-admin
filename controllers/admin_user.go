package controllers

import (
	"beego-admin/form_validate"
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/models"
	"beego-admin/services"
	"beego-admin/utils"
	"encoding/base64"
	"github.com/adam-hanna/arrayOperations"
	"github.com/gookit/validate"
	"strconv"
	"strings"
)

type AdminUserController struct {
	baseController
}

//控制器，初始化函数，基础自父控制器
func (this *AdminUserController) NestPrepare() {
	//fmt.Println("AdminUserController NestPrepare")
}

////用户管理-首页
func (this *AdminUserController) Index() {
	var adminUserService services.AdminUserService
	data, pagination := adminUserService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	this.Data["data"] = data
	this.Data["paginate"] = pagination

	this.Layout = "public/base.html"
	this.TplName = "admin_user/index.html"
}

//用户管理-添加界面
func (this *AdminUserController) Add() {
	var adminRoleService services.AdminRoleService
	roles := adminRoleService.GetAllData()

	this.Data["roles"] = roles
	this.Layout = "public/base.html"
	this.TplName = "admin_user/add.html"
}

//用户管理-添加界面
func (this *AdminUserController) Create() {
	var adminUserForm form_validate.AdminUserForm
	if err := this.ParseForm(&adminUserForm); err != nil {
		response.ErrorWithMessage(err.Error(), this.Ctx)
	}
	roles := make([]string, 0)
	this.Ctx.Input.Bind(&roles, "role")

	adminUserForm.Role = strings.Join(roles, ",")

	v := validate.Struct(adminUserForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), this.Ctx)
	}

	//账号验重
	var adminUserService services.AdminUserService
	if adminUserService.IsExistName(strings.TrimSpace(adminUserForm.Username), 0) {
		response.ErrorWithMessage("账号已经存在", this.Ctx)
	}
	//默认头像
	adminUserForm.Avatar = "/static/admin/images/avatar.png"

	insertId := adminUserService.Create(&adminUserForm)

	url := global.URL_BACK
	if adminUserForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	if insertId > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}

}

//系统管理-用户管理-修改界面
func (this *AdminUserController) Edit() {
	id, _ := this.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", this.Ctx)
	}

	var (
		adminUserService services.AdminUserService
		adminRoleService services.AdminRoleService
	)

	adminUser := adminUserService.GetAdminUserById(id)
	if adminUser == nil {
		response.ErrorWithMessage("Not Found Info By Id.", this.Ctx)
	}

	roles := adminRoleService.GetAllData()
	this.Data["roles"] = roles
	this.Data["data"] = adminUser
	this.Data["role_arr"] = strings.Split(adminUser.Role, ",")

	this.Layout = "public/base.html"
	this.TplName = "admin_user/edit.html"
}

//系统管理-用户管理-修改
func (this *AdminUserController) Update() {
	var adminUserForm form_validate.AdminUserForm
	if err := this.ParseForm(&adminUserForm); err != nil {
		response.ErrorWithMessage(err.Error(), this.Ctx)
	}

	if adminUserForm.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", this.Ctx)
	}

	roles := make([]string, 0)
	this.Ctx.Input.Bind(&roles, "role")

	adminUserForm.Role = strings.Join(roles, ",")

	v := validate.Struct(adminUserForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), this.Ctx)
	}

	//账号验重
	var adminUserService services.AdminUserService
	if adminUserService.IsExistName(strings.TrimSpace(adminUserForm.Username), adminUserForm.Id) {
		response.ErrorWithMessage("账号已经存在", this.Ctx)
	}

	num := adminUserService.Update(&adminUserForm)

	if num > 0 {
		response.Success(this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//启用
func (this *AdminUserController) Enable() {
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
		response.ErrorWithMessage("请选择启用的用户.", this.Ctx)
	}

	var adminUserService services.AdminUserService
	num := adminUserService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//禁用
func (this *AdminUserController) Disable() {
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
		response.ErrorWithMessage("请选择禁用的用户.", this.Ctx)
	}

	var adminUserService services.AdminUserService
	num := adminUserService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//删除
func (this *AdminUserController) Del() {
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

	noDeletionId := new(models.AdminUser).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionId, idArr)

	if len(noDeletionId) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID为"+strings.Join(utils.IntArrToStringArr(noDeletionId), ",")+"的数据无法删除!", this.Ctx)
	}

	var adminUserService services.AdminUserService
	count := adminUserService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//系统管理-个人资料
func (this *AdminUserController) Profile() {
	this.Layout = "public/base.html"
	this.TplName = "admin_user/profile.html"
}

//系统管理-个人资料-修改昵称
func (this *AdminUserController) UpdateNickName() {
	id, err := this.GetInt("id")
	nickname := strings.TrimSpace(this.GetString("nickname"))

	if nickname == "" || err != nil {
		response.ErrorWithMessage("参数错误", this.Ctx)
	}

	var adminUserService services.AdminUserService
	num := adminUserService.UpdateNickName(id, nickname)

	if num > 0 {
		//修改成功后，更新session的登录用户信息
		loginAdminUser := adminUserService.GetAdminUserById(id)
		this.SetSession(global.LOGIN_USER, *loginAdminUser)
		response.SuccessWithMessageAndUrl("修改成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//系统管理-个人资料-修改密码
func (this *AdminUserController) UpdatePassword() {
	id, err := this.GetInt("id")
	password := this.GetString("password")
	newPassword := this.GetString("new_password")
	reNewPassword := this.GetString("renew_password")

	if err != nil || password == "" || newPassword == "" || reNewPassword == "" {
		response.ErrorWithMessage("Bad Parameter.", this.Ctx)
	}

	if newPassword != reNewPassword {
		response.ErrorWithMessage("两次输入的密码不一致.", this.Ctx)
	}

	if password == newPassword {
		response.ErrorWithMessage("新密码与旧密码一致，无需修改", this.Ctx)
	}

	loginUserPassword, err := base64.StdEncoding.DecodeString(loginUser.Password)

	if err != nil {
		response.ErrorWithMessage("err:"+err.Error(), this.Ctx)
	}

	if !utils.PasswordVerify(password, string(loginUserPassword)) {
		response.ErrorWithMessage("当前密码不正确", this.Ctx)
	}

	var adminUserService services.AdminUserService
	num := adminUserService.UpdatePassword(id, newPassword)
	if num > 0 {
		response.SuccessWithMessageAndUrl("修改成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//系统管理-个人资料-修改头像
func (this *AdminUserController) UpdateAvatar() {
	_, _, err := this.GetFile("avatar")
	if err != nil {
		response.ErrorWithMessage("上传头像错误"+err.Error(), this.Ctx)
	}

	var (
		attachmentService services.AttachmentService
		adminUserService  services.AdminUserService
	)
	attachmentInfo, err := attachmentService.Upload(this.Ctx, "avatar", loginUser.Id, 0)
	if err != nil || attachmentInfo == nil {
		response.ErrorWithMessage(err.Error(), this.Ctx)
	} else {
		//头像上传成功，更新用户的avatar头像信息
		num := adminUserService.UpdateAvatar(loginUser.Id, attachmentInfo.Url)
		if num > 0 {
			//修改成功后，更新session的登录用户信息
			loginAdminUser := adminUserService.GetAdminUserById(loginUser.Id)
			this.SetSession(global.LOGIN_USER, *loginAdminUser)
			response.SuccessWithMessageAndUrl("修改成功", global.URL_RELOAD, this.Ctx)
		} else {
			response.Error(this.Ctx)
		}
	}

}
