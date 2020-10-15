package models

import (
	"github.com/astaxie/beego/orm"
)

type AdminLogData struct {
	Id         int    `orm:"column(id);auto;size(11)" description:"表ID"`
	AdminLogId int    `orm:"column(admin_log_id);size(11)" description:"日志ID"`
	Data       string `orm:"column(data);type(text)" description:"日志内容"`
}

//自定义table 名称
func (*AdminLogData) TableName() string {
	return "admin_log_data"
}

//定义模型的可搜索字段
func (*AdminLogData) SearchField() []string {
	return []string{}
}

//定义模型可作为条件的字段
func (*AdminLogData) WhereField() []string {
	return []string{}
}

//定义可做为时间范围查询的字段
func (*AdminLogData) TimeField() []string {
	return []string{}
}

//禁止删除的数据id
func (*AdminLogData) NoDeletionId() []int {
	return []int{}
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(AdminLogData))
}
