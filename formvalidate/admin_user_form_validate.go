package formvalidate

import "github.com/gookit/validate"

// AdminUserForm admin_user 表单
type AdminUserForm struct {
	Id       int    `form:"id"`
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
	Nickname string `form:"nickname" validate:"required"`
	Avatar   string `form:"avatar"`
	Role     string `form:"role" validate:"required"`
	Status   int    `form:"status"`
	IsCreate int    `form:"_create"`
}

// Messages 自定义验证返回消息
func (f AdminUserForm) Messages() map[string]string {
	return validate.MS{
		"Username.required": "请填写账号.",
		"Password.required": "请填写密码.",
		"Nickname.required": "请填写昵称.",
		"Role.required":     "请选择角色.",
	}
}
