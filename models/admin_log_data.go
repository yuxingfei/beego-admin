package models

import (
	"github.com/astaxie/beego/orm"
)

type AdminLogData struct {
	Id           int       `orm:"column(id);auto;size(11)" description:"表ID"`
	AdminLogId   int       `orm:"column(admin_log_id);size(11)" description:"日志ID"`
	Data         string    `orm:"column(data);type(text)" description:"日志内容"`
}

//自定义table 名称
func (adminLogData *AdminLogData) TableName() string {
	return "admin_log_data"
}

//在init中注册定义的model
func init()  {
	orm.RegisterModel(new(AdminLogData))
}


