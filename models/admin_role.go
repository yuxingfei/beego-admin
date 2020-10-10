package models

import (
	"github.com/astaxie/beego/orm"
)

type AdminRole struct {
	Id          int    `orm:"column(id);auto;size(11)" description:"表ID"`
	Name        string `orm:"column(name);size(50)" description:"名称"`
	Description string `orm:"column(description);size(100)" description:"简介"`
	Url         string `orm:"column(url);size(1000)" description:"权限"`
	Status      int8   `orm:"column(status);size(1);default(1)" description:"是否启用 0：否 1：是"`
}

//自定义table 名称
func (adminUser *AdminRole) TableName() string {
	return "admin_role"
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(AdminRole))
}
