package formvalidate

import "github.com/gookit/validate"

// AdminMenuForm 菜单管理-添加菜单表单
type AdminMenuForm struct {
	Id        int    `form:"id"`
	ParentId  int    `form:"parent_id" validate:"min:0"`
	Name      string `form:"name" validate:"required"`
	Url       string `form:"url" validate:"required"`
	Icon      string `form:"icon" validate:"required"`
	IsShow    int8   `form:"is_show"`
	SortId    int    `form:"sort_id" validate:"required"`
	LogMethod string `form:"log_method" validate:"required"`
	IsCreate  int    `form:"_create"`
}

// Messages 自定义验证返回消息
func (f AdminMenuForm) Messages() map[string]string {
	return validate.MS{
		"ParentId.min":       "parent_id必须大于等于0.",
		"Name.required":      "菜单名称不能为空.",
		"Url.required":       "url不能为空.",
		"Icon.required":      "图标不能为空.",
		"SortId.required":    "排序不能为空.",
		"LogMethod.required": "记录类型不能为空.",
	}
}
