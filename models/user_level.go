package models

import (
	"github.com/astaxie/beego/orm"
)

type UserLevel struct {
	Id          int    `orm:"column(id);auto;size(11)" description:"表ID" json:"id"`
	Name        string `orm:"column(name);size(20)" description:"名称" json:"name"`
	Description string `orm:"column(description);size(50)" description:"简介" json:"description"`
	Img         string `orm:"column(img);size(255)" description:"图片" json:"img"`
	Status      int8   `orm:"column(status);size(1);default(1)" description:"是否启用 0：否 1：是" json:"status"`
	CreateTime  int    `orm:"column(create_time);size(10);default(0)" description:"操作时间" json:"create_time"`
	UpdateTime  int    `orm:"column(update_time);size(10);default(0)" description:"更新时间" json:"update_time"`
	DeleteTime  int    `orm:"column(delete_time);;size(10);default(0)" description:"删除时间" json:"delete_time"`
}

//自定义table 名称
func (*UserLevel) TableName() string {
	return "user_level"
}

//定义模型的可搜索字段
func (*UserLevel) SearchField() []string {
	return []string{"name", "description"}
}

//禁止删除的数据id
func (*UserLevel) NoDeletionId() []int {
	return []int{}
}

//定义模型可作为条件的字段
func (*UserLevel) WhereField() []string {
	return []string{}
}

//定义可做为时间范围查询的字段
func (*UserLevel) TimeField() []string {
	return []string{}
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(UserLevel))
}
