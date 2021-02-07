package models

import (
	"github.com/beego/beego/v2/client/orm"
)

//Content 将Content json字符串转换为结构体，便于html界面中range
type Content struct {
	Name    string
	Field   string
	Type    string
	Content string
	Option  string
	Form    string
}

// Setting struct
type Setting struct {
	Id             int    `orm:"column(id);auto;size(11)" description:"表ID" json:"id"`
	SettingGroupId int    `orm:"column(setting_group_id);size(10);default(0)" description:"所属分组" json:"setting_group_id"`
	Name           string `orm:"column(name);size(50)" description:"名称" json:"name"`
	Description    string `orm:"column(description);size(100)" description:"描述" json:"description"`
	Code           string `orm:"column(code);size(50)" description:"代码" json:"code"`
	Content        string `orm:"column(content);type(text)" description:"设置配置及内容" json:"content"`
	//将Content json字符串转换为结构体
	ContentStrut []*Content `orm:"-"`
	SortNumber   int        `orm:"column(sort_number);size(10);default(1000)" description:"排序" json:"sort_number"`
	CreateTime   int        `orm:"column(create_time);size(10);default(0)" description:"操作时间"`
	UpdateTime   int        `orm:"column(update_time);size(10);default(0)" description:"更新时间"`
	DeleteTime   int        `orm:"column(delete_time);;size(10);default(0)" description:"删除时间" json:"delete_time"`
}

// TableName 自定义table 名称
func (*Setting) TableName() string {
	return "setting"
}

// SearchField 定义模型的可搜索字段
func (*Setting) SearchField() []string {
	return []string{"name", "description", "code"}
}

// NoDeletionId 禁止删除的数据id
func (*Setting) NoDeletionId() []int {
	return []int{1, 2, 3, 4, 5}
}

// WhereField 定义模型可作为条件的字段
func (*Setting) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*Setting) TimeField() []string {
	return []string{}
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(Setting))
}
