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

//根据设置id，获取对应的setting info
func (*SettingService) GetSettingInfoById(id int) *models.Setting {
	setting := models.Setting{Id: id}
	orm.NewOrm().Read(&setting)
	return &setting
}

//根据id修改content的内容
func (*SettingService) UpdateSettingInfoToContent(id int, content string) int {
	affectRow, err := orm.NewOrm().QueryTable(new(models.Setting)).Filter("id", id).Update(orm.Params{
		"content": content,
	})
	if err == nil {
		return int(affectRow)
	} else {
		return 0
	}
}
