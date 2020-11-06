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

type UserLevelController struct {
	baseController
}

//用户等级 列表页
func (this *UserLevelController) Index() {
	var userLevelService services.UserLevelService
	data, pagination := userLevelService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	this.Data["data"] = data
	this.Data["paginate"] = pagination

	this.Layout = "public/base.html"
	this.TplName = "user_level/index.html"
}

func (this *UserLevelController) Export() {
	exportData := this.GetString("export_data")
	if exportData == "1" {
		var userLevelService services.UserLevelService
		data := userLevelService.GetExportData(gQueryParams)
		header := []string{"ID", "名称", "简介", "是否启用", "创建时间"}
		body := [][]string{}
		for _, item := range data {
			record := []string{
				strconv.Itoa(item.Id),
				item.Name,
				item.Description,
			}
			if item.Status == 1 {
				record = append(record, "是")
			} else {
				record = append(record, "否")
			}
			record = append(record, template.UnixTimeForFormat(item.CreateTime))
			body = append(body, record)
		}
		this.Ctx.ResponseWriter.Header().Set("a", "b")
		excel_office.ExportData(header, body, "user_level-"+time.Now().Format("2006-01-02-15-04-05"), "", "", this.Ctx.ResponseWriter)
	}

	response.Error(this.Ctx)
}

//用户等级-添加界面
func (this *UserLevelController) Add() {
	this.Layout = "public/base.html"
	this.TplName = "user_level/add.html"
}

//用户等级-添加
func (this *UserLevelController) Create() {
	var userLevelForm form_validate.UserLevelForm
	if err := this.ParseForm(&userLevelForm); err != nil {
		response.ErrorWithMessage(err.Error(), this.Ctx)
	}

	v := validate.Struct(userLevelForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), this.Ctx)
	}

	//处理图片上传
	_, _, err := this.GetFile("img")
	if err == nil {
		var attachmentService services.AttachmentService
		attachmentInfo, err := attachmentService.Upload(this.Ctx, "img", loginUser.Id, 0)
		if err != nil || attachmentInfo == nil {
			response.ErrorWithMessage(err.Error(), this.Ctx)
		} else {
			userLevelForm.Img = attachmentInfo.Url
		}
	}

	var userLevelService services.UserLevelService
	insertId := userLevelService.Create(&userLevelForm)

	url := global.URL_BACK
	if userLevelForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	if insertId > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//用户等级-修改界面
func (this *UserLevelController) Edit() {
	id, _ := this.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", this.Ctx)
	}

	var userLevelService services.UserLevelService

	userLevel := userLevelService.GetUserLevelById(id)
	if userLevel == nil {
		response.ErrorWithMessage("Not Found Info By Id.", this.Ctx)
	}

	this.Data["data"] = userLevel

	this.Layout = "public/base.html"
	this.TplName = "user_level/edit.html"
}

//用户等级-修改
func (this *UserLevelController) Update() {
	var userLevelForm form_validate.UserLevelForm
	if err := this.ParseForm(&userLevelForm); err != nil {
		response.ErrorWithMessage(err.Error(), this.Ctx)
	}

	if userLevelForm.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", this.Ctx)
	}

	v := validate.Struct(userLevelForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), this.Ctx)
	}

	_, _, err := this.GetFile("img")
	if err == nil {
		//处理图片上传
		var attachmentService services.AttachmentService
		attachmentInfo, err := attachmentService.Upload(this.Ctx, "img", loginUser.Id, 0)
		if err != nil || attachmentInfo == nil {
			response.ErrorWithMessage(err.Error(), this.Ctx)
		} else {
			userLevelForm.Img = attachmentInfo.Url
		}
	}

	var userLevelService services.UserLevelService
	num := userLevelService.Update(&userLevelForm)

	if num > 0 {
		response.Success(this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//启用
func (this *UserLevelController) Enable() {
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

	var userLevelService services.UserLevelService
	num := userLevelService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//禁用
func (this *UserLevelController) Disable() {
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

	var userLevelService services.UserLevelService
	num := userLevelService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//删除
func (this *UserLevelController) Del() {
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

	noDeletionId := new(models.UserLevel).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionId, idArr)

	if len(noDeletionId) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID为"+strings.Join(utils.IntArrToStringArr(noDeletionId), ",")+"的数据无法删除!", this.Ctx)
	}

	var userLevelService services.UserLevelService
	count := userLevelService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}
