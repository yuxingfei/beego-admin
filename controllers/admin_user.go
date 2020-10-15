package controllers

import (
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/services"
	"beego-admin/utils"
	"encoding/base64"
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
	data, pagination := adminUserService.GetPaginateData(admin["per_page"].(int), queryParams)
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
