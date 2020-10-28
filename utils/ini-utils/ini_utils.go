package ini_utils

import (
	"beego-admin/models"
	"fmt"
	"gopkg.in/ini.v1"
)

//更新conf/admin.conf
func UpdateAdminConf(setting *models.Setting) bool {

	if setting == nil {
		return false
	}

	cfg, err := ini.Load("conf/admin.conf")
	if err != nil {
		fmt.Printf("Fail to read file conf/admin.conf: %v", err)
		return false
	}

	if setting.Code == "base" {
		for _, content := range setting.ContentStrut {
			switch content.Field {
			case "name":
				cfg.Section("base").Key("name").SetValue(content.Content)
			case "short_name":
				cfg.Section("base").Key("short_name").SetValue(content.Content)
			case "author":
				cfg.Section("base").Key("author").SetValue(content.Content)
			case "version":
				cfg.Section("base").Key("version").SetValue(content.Content)
			case "link":
				cfg.Section("base").Key("link").SetValue(content.Content)
			}
		}
	} else if setting.Code == "login" {
		for _, content := range setting.ContentStrut {
			switch content.Field {
			case "token":
				cfg.Section("login").Key("token").SetValue(content.Content)
			case "captcha":
				cfg.Section("login").Key("captcha").SetValue(content.Content)
			case "background":
				cfg.Section("login").Key("background").SetValue(content.Content)
			}
		}
	} else if setting.Code == "index" {
		for _, content := range setting.ContentStrut {
			switch content.Field {
			case "password_warning":
				cfg.Section("index").Key("password_warning").SetValue(content.Content)
			case "show_notice":
				cfg.Section("index").Key("show_notice").SetValue(content.Content)
			case "notice_content":
				cfg.Section("index").Key("notice_content").SetValue(content.Content)
			}
		}
	} else {
		return false
	}

	cfg.SaveTo("conf/admin.conf")

	return true
}
