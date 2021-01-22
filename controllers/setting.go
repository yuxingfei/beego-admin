package controllers

import (
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/services"
	"encoding/json"
	"strconv"
)

// SettingController struct
type SettingController struct {
	baseController
}

// Admin 设置
func (sc *SettingController) Admin() {
	sc.show(1)
}

// show 展示单个配置信息
func (sc *SettingController) show(id int) {
	var settingService services.SettingService
	data := settingService.Show(id)

	sc.Data["data_config"] = data
	sc.Layout = "public/base.html"
	sc.TplName = "setting/show.html"
}

// Update 设置中心-更新设置
func (sc *SettingController) Update() {
	id := sc.Input().Get("id")

	if id == "" {
		response.ErrorWithMessage("参数错误.", sc.Ctx)
	}

	var settingService services.SettingService
	idInt, _ := strconv.Atoi(id)
	setting := settingService.GetSettingInfoById(idInt)

	if setting == nil {
		response.ErrorWithMessage("无法更新配置信息", sc.Ctx)
	}

	err := json.Unmarshal([]byte(setting.Content), &setting.ContentStrut)
	if err != nil {
		response.ErrorWithMessage("无法更新配置信息", sc.Ctx)
	}

	for key, value := range setting.ContentStrut {
		switch value.Type {
		case "image", "file":
			//单个文件上传
			var attachmentService services.AttachmentService
			attachmentInfo, err := attachmentService.Upload(sc.Ctx, value.Field, loginUser.Id, 0)
			if err == nil && attachmentInfo != nil {
				//图片上传成功
				setting.ContentStrut[key].Content = attachmentInfo.Url
			}
		case "multi_file", "multi_image":
			//多个文件上传
			var attachmentService services.AttachmentService
			attachments, err := attachmentService.UploadMulti(sc.Ctx, value.Field, loginUser.Id, 0)
			if err == nil && attachments != nil {
				var urls []string
				for _, atta := range attachments {
					urls = append(urls, atta.Url)
				}
				if len(urls) > 0 {
					urlsByte, err := json.Marshal(&urls)
					if err == nil {
						setting.ContentStrut[key].Content = string(urlsByte)
					}
				}
			}
		default:
			setting.ContentStrut[key].Content = sc.Input().Get(value.Field)
		}
	}

	//修改内容
	contentStrutByte, err := json.Marshal(&setting.ContentStrut)
	if err == nil {
		affectRow := settingService.UpdateSettingInfoToContent(idInt, string(contentStrutByte))
		if affectRow > 0 {
			//更新全局配置
			settingService.LoadOrUpdateGlobalBaseConfig(setting)
			response.SuccessWithMessageAndUrl("修改成功", global.URL_RELOAD, sc.Ctx)
		} else {
			response.ErrorWithMessage("没有可更新的信息", sc.Ctx)
		}
	} else {
		response.ErrorWithMessage("修改失败 err:"+err.Error(), sc.Ctx)
	}

}