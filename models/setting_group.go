package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// SettingGroup struct
type SettingGroup struct {
	Id             int    `orm:"column(id);auto;size(11)" description:"表ID" json:"id"`
	Module         string `orm:"column(module);size(30)" description:"作用模块" json:"module"`
	Name           string `orm:"column(name);size(50)" description:"名称" json:"name"`
	Description    string `orm:"column(description);size(100)" description:"描述" json:"description"`
	Code           string `orm:"column(code);size(50)" description:"代码" json:"code"`
	SortNumber     int    `orm:"column(sort_number);size(10);default(1000)" description:"排序" json:"sort_number"`
	AutoCreateMenu int    `orm:"column(auto_create_menu);size(1);default(0)" description:"自动生成菜单" json:"auto_create_menu"`
	AutoCreateFile int    `orm:"column(auto_create_file);size(1);default(0)" description:"自动生成配置文件" json:"auto_create_file"`
	Icon           string `orm:"column(icon);size(30);default(fa-list)" description:"图标" json:"icon"`
	CreateTime     int    `orm:"column(create_time);size(10);default(0)" description:"操作时间"`
	UpdateTime     int    `orm:"column(update_time);size(10);default(0)" description:"更新时间"`
	DeleteTime     int    `orm:"column(delete_time);;size(10);default(0)" description:"删除时间" json:"delete_time"`
}

// TableName 自定义table 名称
func (*SettingGroup) TableName() string {
	return "setting_group"
}

// SearchField 定义模型的可搜索字段
func (*SettingGroup) SearchField() []string {
	return []string{"name", "description", "code"}
}

// NoDeletionId 禁止删除的数据id
func (*SettingGroup) NoDeletionId() []int {
	return []int{1, 2, 3, 4, 5}
}

// WhereField 定义模型可作为条件的字段
func (*SettingGroup) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*SettingGroup) TimeField() []string {
	return []string{}
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(SettingGroup))
}
