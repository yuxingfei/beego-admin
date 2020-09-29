package services

import (
	"beego-admin/models"
	"github.com/astaxie/beego/orm"
)

type AdminMenuService struct {

}

//根据url获取admin_menu数据
func (*AdminMenuService)GetAdminMenuByUrl(url string) *models.AdminMenu {
	var adminMenu models.AdminMenu
	err := orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("url",url).One(&adminMenu)
	if err == nil{
		return &adminMenu
	}else{
		return nil
	}
}

//获取admin_menu 总数
func (*AdminMenuService)GetCount() int {
	count,err := orm.NewOrm().QueryTable(new(models.AdminMenu)).Count()
	if err != nil{
		return 0
	}
	return int(count)
}
