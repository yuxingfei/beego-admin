package services

import (
	"beego-admin/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type SettingService struct {
	BaseService
}

func (settingService *SettingService) Show(id int) []*models.Setting {
	data := settingService.getDataBySettingGroupId(id)

	var settingFormService SettingFormService

	for key, value := range data {
		//contentNew := ""
		//value.Content转为json
		var contents []*models.Content
		if value.Content == "" {
			continue
		}
		err := json.Unmarshal([]byte(value.Content), &contents)

		if err != nil {
			continue
		}

		var contentNew []*models.Content
		for _, content := range contents {
			content.Form = settingFormService.GetFieldForm(content.Type, content.Name, content.Field, content.Content, content.Option)
			contentNew = append(contentNew, content)
		}
		data[key].ContentStrut = contentNew
	}

	return data
}

//根据设置分组id获取多个设置信息
func (*SettingService) getDataBySettingGroupId(settingGroupId int) []*models.Setting {
	var settings []*models.Setting
	_, err := orm.NewOrm().QueryTable(new(models.Setting)).Filter("setting_group_id", settingGroupId).All(&settings)
	if err != nil {
		return nil
	} else {
		return settings
	}
}
