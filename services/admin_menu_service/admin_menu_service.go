package admin_menu_service

import (
	"beego-admin/models"
	"github.com/astaxie/beego/orm"
)

//根据url获取admin_menu数据
func GetAdminMenuByUrl(url string) *models.AdminMenu {
	var adminMenu models.AdminMenu
	err := orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("url",url).One(&adminMenu)
	if err == nil{
		return &adminMenu
	}else{
		return nil
	}
}
