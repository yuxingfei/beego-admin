package controllers

import (
	"beego-admin/services"
	"net/url"
	"strings"
)

// EditorController struct
type EditorController struct {
	baseController
}

// Prepare 预准备
func (ec *EditorController) Prepare() {
	//取消_xsrf验证
	ec.EnableXSRF = false
}

// Server 方法
func (ec *EditorController) Server() {
	var result map[string]interface{}
	var ueditorService services.UeditorService
	action := ec.GetString("action")
	switch action {
	case "config":
		result = ueditorService.GetConfig()
	case "uploadimage":
		//上传图片
		result = ueditorService.UploadImage(ec.Ctx)
	case "uploadscrawl":
		//上传涂鸦
		//带+号的值如果不处理，会被转为空格
		ec.Ctx.Request.URL.RawQuery = strings.ReplaceAll(ec.Ctx.Request.URL.RawQuery, "+", "%2b")
		values, _ := url.ParseQuery(ec.Ctx.Request.URL.RawQuery)
		result = ueditorService.UploadScrawl(values)
	case "uploadvideo":
		//上传视频
		result = ueditorService.UploadVideo(ec.Ctx)
	case "uploadfile":
		//上传文件
		result = ueditorService.UploadFile(ec.Ctx)
	case "listimage":
		//列出图片
		result = ueditorService.ListImage(ec.Input())
	case "listfile":
		//列出文件
		result = ueditorService.ListFile(ec.Input())
	case "catchimage":
		//抓取远程文件
		result = ueditorService.CatchImage(ec.Ctx)
	default:
		result = map[string]interface{}{
			"state": "请求地址出错",
		}
	}
	ec.Data["json"] = result
	ec.ServeJSON()

}
