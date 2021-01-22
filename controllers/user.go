package controllers

import (
	"beego-admin/formvalidate"
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/models"
	"beego-admin/services"
	"beego-admin/utils"
	"beego-admin/utils/exceloffice"
	"beego-admin/utils/template"
	"github.com/adam-hanna/arrayOperations"
	"github.com/gookit/validate"
	"strconv"
	"strings"
	"time"
)

// UserController struct
type UserController struct {
	baseController
}

// Index 用户等级 列表页
func (uc *UserController) Index() {
	var userService services.UserService
	var userLevelService services.UserLevelService

	//获取用户等级
	userLevel := userLevelService.GetUserLevel()
	userLevelMap := make(map[int]string)
	for _, item := range userLevel {
		userLevelMap[item.Id] = item.Name
	}

	data, pagination := userService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	uc.Data["data"] = data
	uc.Data["paginate"] = pagination
	uc.Data["user_level_map"] = userLevelMap

	uc.Layout = "public/base.html"
	uc.TplName = "user/index.html"
}

// Export 导出
func (uc *UserController) Export() {
	exportData := uc.GetString("export_data")
	if exportData == "1" {
		var userService services.UserService
		var userLevelService services.UserLevelService
		userLevel := userLevelService.GetUserLevel()
		userLevelMap := make(map[int]string)
		for _, item := range userLevel {
			userLevelMap[item.Id] = item.Name
		}

		data := userService.GetExportData(gQueryParams)
		header := []string{"ID", "头像", "用户等级", "用户名", "手机号", "昵称", "是否启用", "创建时间"}
		body := [][]string{}
		for _, item := range data {
			record := []string{
				strconv.Itoa(item.Id),
				item.Avatar,
			}
			userLevelName, ok := userLevelMap[item.UserLevelId]
			if ok {
				record = append(record, userLevelName)
			}
			record = append(record, item.Username)
			record = append(record, item.Mobile)
			record = append(record, item.Nickname)

			if item.Status == 1 {
				record = append(record, "是")
			} else {
				record = append(record, "否")
			}
			record = append(record, template.UnixTimeForFormat(item.CreateTime))
			body = append(body, record)
		}
		uc.Ctx.ResponseWriter.Header().Set("a", "b")
		exceloffice.ExportData(header, body, "user-"+time.Now().Format("2006-01-02-15-04-05"), "", "", uc.Ctx.ResponseWriter)
	}

	response.Error(uc.Ctx)
}

// Add 用户-添加界面
func (uc *UserController) Add() {
	var userLevelService services.UserLevelService

	//获取用户等级
	userLevel := userLevelService.GetUserLevel()

	uc.Data["user_level_list"] = userLevel
	uc.Layout = "public/base.html"
	uc.TplName = "user/add.html"
}

// Create 添加用户
func (uc *UserController) Create() {
	var userForm formvalidate.UserForm
	if err := uc.ParseForm(&userForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	v := validate.Struct(userForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	//处理图片上传
	_, _, err := uc.GetFile("avatar")
	if err == nil {
		var attachmentService services.AttachmentService
		attachmentInfo, err := attachmentService.Upload(uc.Ctx, "avatar", loginUser.Id, 0)
		if err != nil || attachmentInfo == nil {
			response.ErrorWithMessage(err.Error(), uc.Ctx)
		} else {
			userForm.Avatar = attachmentInfo.Url
		}
	}

	var userService services.UserService
	insertID := userService.Create(&userForm)

	url := global.URL_BACK
	if userForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	if insertID > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Edit 用户-修改界面
func (uc *UserController) Edit() {
	id, _ := uc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", uc.Ctx)
	}

	var userService services.UserService

	user := userService.GetUserById(id)
	if user == nil {
		response.ErrorWithMessage("Not Found Info By Id.", uc.Ctx)
	}

	//获取用户等级
	var userLevelService services.UserLevelService
	userLevel := userLevelService.GetUserLevel()

	uc.Data["user_level_list"] = userLevel
	uc.Data["data"] = user

	uc.Layout = "public/base.html"
	uc.TplName = "user/edit.html"
}

// Update 用户-修改
func (uc *UserController) Update() {
	var userForm formvalidate.UserForm
	if err := uc.ParseForm(&userForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	if userForm.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", uc.Ctx)
	}

	v := validate.Struct(userForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	_, _, err := uc.GetFile("avatar")
	if err == nil {
		//处理图片上传
		var attachmentService services.AttachmentService
		attachmentInfo, err := attachmentService.Upload(uc.Ctx, "avatar", loginUser.Id, 0)
		if err != nil || attachmentInfo == nil {
			response.ErrorWithMessage(err.Error(), uc.Ctx)
		} else {
			userForm.Avatar = attachmentInfo.Url
		}
	}

	var userService services.UserService
	num := userService.Update(&userForm)

	if num > 0 {
		response.Success(uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Enable 启用
func (uc *UserController) Enable() {
	idStr := uc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		uc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("请选择用户等级.", uc.Ctx)
	}

	var userService services.UserService
	num := userService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Disable 禁用
func (uc *UserController) Disable() {
	idStr := uc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		uc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("请选择禁用的用户.", uc.Ctx)
	}

	var userService services.UserService
	num := userService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Del 删除
func (uc *UserController) Del() {
	idStr := uc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		uc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("参数id错误.", uc.Ctx)
	}

	noDeletionID := new(models.User).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID为"+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"的数据无法删除!", uc.Ctx)
	}

	var userService services.UserService
	count := userService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}
