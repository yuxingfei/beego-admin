package controllers

import (
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/services"
	ini_utils "beego-admin/utils/ini-utils"
	"encoding/json"
	"fmt"
	"strconv"
)

type SettingController struct {
	baseController
}

//设置
func (this *SettingController) Admin() {
	this.show(1)
}

//展示单个配置信息
func (this *SettingController) show(id int) {
	var settingService services.SettingService
	data := settingService.Show(id)

	this.Data["data_config"] = data
	this.Layout = "public/base.html"
	this.TplName = "setting/show.html"
}

//设置中心-更新设置
func (this *SettingController) Update() {
	id := this.Input().Get("id")

	if id == "" {
		response.ErrorWithMessage("参数错误.", this.Ctx)
	}

	var settingService services.SettingService
	idInt, _ := strconv.Atoi(id)
	setting := settingService.GetSettingInfoById(idInt)

	if setting == nil {
		response.ErrorWithMessage("无法更新配置信息", this.Ctx)
	}

	err := json.Unmarshal([]byte(setting.Content), &setting.ContentStrut)
	if err != nil {
		response.ErrorWithMessage("无法更新配置信息", this.Ctx)
	}

	for key, value := range setting.ContentStrut {
		switch value.Type {
		case "image", "file":
			//单个文件上传
			var attachmentService services.AttachmentService
			attachmentInfo, err := attachmentService.Upload(this.Ctx, value.Field, loginUser.Id, 0)
			fmt.Println("attachmentInfo = ", attachmentInfo)
			if err == nil && attachmentInfo != nil {
				//图片上传成功
				setting.ContentStrut[key].Content = attachmentInfo.Url
			}
		case "multi_file", "multi_image":
			//多个文件上传
			var attachmentService services.AttachmentService
			attachments, err := attachmentService.UploadMulti(this.Ctx, value.Field, loginUser.Id, 0)
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
			setting.ContentStrut[key].Content = this.Input().Get(value.Field)
		}
	}

	//修改内容
	contentStrutByte, err := json.Marshal(&setting.ContentStrut)
	fmt.Println("contentStrutByte = ", string(contentStrutByte))
	if err == nil {
		affectRow := settingService.UpdateSettingInfoToContent(idInt, string(contentStrutByte))
		if affectRow > 0 {
			//自动更新配置文件
			ini_utils.UpdateAdminConf(setting)
			response.SuccessWithMessageAndUrl("修改成功", global.URL_RELOAD, this.Ctx)
		} else {
			response.ErrorWithMessage("没有可更新的信息", this.Ctx)
		}
	} else {
		response.ErrorWithMessage("修改失败 err:"+err.Error(), this.Ctx)
	}

}
