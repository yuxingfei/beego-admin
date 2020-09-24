package admin_role_service

import (
	"beego-admin/models"
	"github.com/astaxie/beego/orm"
)

//获取admin_role 总数
func GetCount() int {
	count,err := orm.NewOrm().QueryTable(new(models.AdminRole)).Count()
	if err != nil{
		return 0
	}
	return int(count)
}
