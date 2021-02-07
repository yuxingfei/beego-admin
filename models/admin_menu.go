package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// AdminMenu struct
type AdminMenu struct {
	Id        int    `orm:"column(id);auto;size(11)" description:"表ID"`
	ParentId  int    `orm:"column(parent_id);size(11);default(0)" description:"父级菜单"`
	Name      string `orm:"column(name);size(30)" description:"名称"`
	Url       string `orm:"column(url);size(100);index" description:"url"`
	Icon      string `orm:"column(icon);size(30);default(fa-list)" description:"图标"`
	IsShow    int8   `orm:"column(is_show);size(1);default(1)" description:"等级"`
	SortId    int    `orm:"column(sort_id);size(10);default(1000)" description:"排序"`
	LogMethod string `orm:"column(log_method);size(8);default(不记录)" description:"记录日志方法"`
}

// TableName 自定义table 名称
func (*AdminMenu) TableName() string {
	return "admin_menu"
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(AdminMenu))
}

// GetLogMethod 自定义方法
func (*AdminMenu) GetLogMethod() []string {
	return []string{"不记录", "GET", "POST", "PUT", "DELETE"}
}

// SearchField 定义模型的可搜索字段
func (*AdminMenu) SearchField() []string {
	return []string{}
}

// WhereField 定义模型可作为条件的字段
func (*AdminMenu) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*AdminMenu) TimeField() []string {
	return []string{}
}

// NoDeletionId 不允许删除的id
func (*AdminMenu) NoDeletionId() []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
}
