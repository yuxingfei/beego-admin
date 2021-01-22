package formvalidate

import "github.com/gookit/validate"

// AdminRoleForm 菜单管理-添加角色表单
type AdminRoleForm struct {
	Id          int    `form:"id"`
	Name        string `form:"name" validate:"required"`
	Description string `form:"description" validate:"required"`
	Status      int8   `form:"status" validate:"int"`
	IsCreate    int    `form:"_create"`
}

// Messages 自定义验证返回消息
func (f AdminRoleForm) Messages() map[string]string {
	return validate.MS{
		"Name.required":        "名称不能为空.",
		"Description.required": "介绍不能为空.",
		"Status.int":           "请选择启用状态.",
	}
}
