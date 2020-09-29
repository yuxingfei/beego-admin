package services

import (
	"beego-admin/models"
	"github.com/astaxie/beego/orm"
)

type AdminRoleService struct {

}

//获取admin_role 总数
func (*AdminRoleService)GetCount() int {
	count,err := orm.NewOrm().QueryTable(new(models.AdminRole)).Count()
	if err != nil{
		return 0
	}
	return int(count)
}
