package models

import (
	"github.com/astaxie/beego/orm"
)

type AdminMenu struct {
	Id          int       `orm:"column(id);auto;size(11)" description:"表ID"`
	ParentId    int       `orm:"column(parent_id);size(11);default(0)" description:"父级菜单"`
	Name        string    `orm:"column(name);size(30)" description:"名称"`
	Url         string    `orm:"column(url);size(100);index" description:"url"`
	RouteName   string    `orm:"column(route_name);size(100)" description:"路由名或者路由标识"`
	Icon        string    `orm:"column(icon);size(30);default(fa-list)" description:"图标"`
	IsShow      int8      `orm:"column(is_show);size(1);default(1)" description:"等级"`
	SortId      int       `orm:"column(sort_id);size(10);default(1000)" description:"排序"`
	LogMethod   string    `orm:"column(log_method);size(8);default(不记录)" description:"记录日志方法"`
}

//自定义table 名称
func (adminUser *AdminMenu) TableName() string {
	return "admin_menu"
}

//在init中注册定义的model
func init()  {
	orm.RegisterModel(new(AdminMenu))
}


