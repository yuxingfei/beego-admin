package controllers

import (
	"beego-admin/form_validate"
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/models"
	"beego-admin/services"
	"beego-admin/utils"
	excel_office "beego-admin/utils/excel-office"
	"beego-admin/utils/template"
	"github.com/adam-hanna/arrayOperations"
	"github.com/gookit/validate"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	baseController
}

//用户等级 列表页
func (this *UserController) Index() {
	var userService services.UserService
	var userLevelService services.UserLevelService

	//获取用户等级
	userLevel := userLevelService.GetUserLevel()
	userLevelMap := make(map[int]string)
	for _, item := range userLevel {
		userLevelMap[item.Id] = item.Name
	}

	data, pagination := userService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	this.Data["data"] = data
	this.Data["paginate"] = pagination
	this.Data["user_level_map"] = userLevelMap

	this.Layout = "public/base.html"
	this.TplName = "user/index.html"
}

func (this *UserController) Export() {
	exportData := this.GetString("export_data")
	if exportData == "1" {
		var userService services.UserService
		var userLevelService services.UserLevelService
		userLevel := userLevelService.GetUserLevel()
		userLevelMap := make(map[int]string)
		for _,item := range userLevel{
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
			userLevelName,ok := userLevelMap[item.UserLevelId]
			if ok {
				record = append(record,userLevelName)
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
		this.Ctx.ResponseWriter.Header().Set("a", "b")
		excel_office.ExportData(header, body, "user-"+time.Now().Format("2006-01-02-15-04-05"), "", "", this.Ctx.ResponseWriter)
	}

	response.Error(this.Ctx)
}

//用户-添加界面
func (this *UserController) Add() {
	var userLevelService services.UserLevelService

	//获取用户等级
	userLevel := userLevelService.GetUserLevel()

	this.Data["user_level_list"] = userLevel
	this.Layout = "public/base.html"
	this.TplName = "user/add.html"
}

//添加用户
func (this *UserController) Create() {
	var userForm form_validate.UserForm
	if err := this.ParseForm(&userForm); err != nil {
		response.ErrorWithMessage(err.Error(), this.Ctx)
	}

	v := validate.Struct(userForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), this.Ctx)
	}

	//处理图片上传
	_, _, err := this.GetFile("avatar")
	if err == nil {
		var attachmentService services.AttachmentService
		attachmentInfo, err := attachmentService.Upload(this.Ctx, "avatar", loginUser.Id, 0)
		if err != nil || attachmentInfo == nil {
			response.ErrorWithMessage(err.Error(), this.Ctx)
		} else {
			userForm.Avatar = attachmentInfo.Url
		}
	}

	var userService services.UserService
	insertId := userService.Create(&userForm)

	url := global.URL_BACK
	if userForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	if insertId > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//用户-修改界面
func (this *UserController) Edit() {
	id, _ := this.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", this.Ctx)
	}

	var userService services.UserService

	user := userService.GetUserById(id)
	if user == nil {
		response.ErrorWithMessage("Not Found Info By Id.", this.Ctx)
	}

	//获取用户等级
	var userLevelService services.UserLevelService
	userLevel := userLevelService.GetUserLevel()

	this.Data["user_level_list"] = userLevel
	this.Data["data"] = user

	this.Layout = "public/base.html"
	this.TplName = "user/edit.html"
}

//用户-修改
func (this *UserController) Update() {
	var userForm form_validate.UserForm
	if err := this.ParseForm(&userForm); err != nil {
		response.ErrorWithMessage(err.Error(), this.Ctx)
	}

	if userForm.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", this.Ctx)
	}

	v := validate.Struct(userForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), this.Ctx)
	}

	_, _, err := this.GetFile("avatar")
	if err == nil {
		//处理图片上传
		var attachmentService services.AttachmentService
		attachmentInfo, err := attachmentService.Upload(this.Ctx, "avatar", loginUser.Id, 0)
		if err != nil || attachmentInfo == nil {
			response.ErrorWithMessage(err.Error(), this.Ctx)
		} else {
			userForm.Avatar = attachmentInfo.Url
		}
	}

	var userService services.UserService
	num := userService.Update(&userForm)

	if num > 0 {
		response.Success(this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//启用
func (this *UserController) Enable() {
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
		response.ErrorWithMessage("请选择用户等级.", this.Ctx)
	}

	var userService services.UserService
	num := userService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//禁用
func (this *UserController) Disable() {
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

	var userService services.UserService
	num := userService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//删除
func (this *UserController) Del() {
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

	noDeletionId := new(models.User).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionId, idArr)

	if len(noDeletionId) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID为"+strings.Join(utils.IntArrToStringArr(noDeletionId), ",")+"的数据无法删除!", this.Ctx)
	}

	var userService services.UserService
	count := userService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}