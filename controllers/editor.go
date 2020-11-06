package controllers

import (
	"beego-admin/services"
	"net/url"
	"strings"
)

type EditorController struct {
	baseController
}

func (this *EditorController) Prepare() {
	//取消_xsrf验证
	this.EnableXSRF = false
}

func (this *EditorController) Server() {
	result := make(map[string]interface{})

	var ueditorService services.UeditorService
	action := this.GetString("action")
	switch action {
	case "config":
		result = ueditorService.GetConfig()
	case "uploadimage":
		//上传图片
		result = ueditorService.UploadImage(this.Ctx)
	case "uploadscrawl":
		//上传涂鸦
		//带+号的值如果不处理，会被转为空格
		this.Ctx.Request.URL.RawQuery = strings.ReplaceAll(this.Ctx.Request.URL.RawQuery, "+", "%2b")
		values, _ := url.ParseQuery(this.Ctx.Request.URL.RawQuery)
		result = ueditorService.UploadScrawl(values)
	case "uploadvideo":
		//上传视频
		result = ueditorService.UploadVideo(this.Ctx)
	case "uploadfile":
		//上传文件
		result = ueditorService.UploadFile(this.Ctx)
	case "listimage":
		//列出图片
		result = ueditorService.ListImage(this.Input())
	case "listfile":
		//列出文件
		result = ueditorService.ListFile(this.Input())
	case "catchimage":
		//抓取远程文件
		result = ueditorService.CatchImage(this.Ctx)
	default:
		this.Data["json"] = map[string]string{
			"state": "请求地址出错",
		}
	}
	this.Data["json"] = result
	this.ServeJSON()

}
