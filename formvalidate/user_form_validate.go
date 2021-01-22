package formvalidate

import "github.com/gookit/validate"

// UserForm user 表单
type UserForm struct {
	Id          int    `form:"id"`
	Avatar      string `form:"avatar"`
	Username    string `form:"username" validate:"required"`
	Nickname    string `form:"nickname" validate:"required"`
	Mobile      string `form:"mobile" validate:"required"`
	UserLevelId int    `form:"user_level_id" validate:"required"`
	Password    string `form:"password" validate:"required"`
	Status      int8   `form:"status"`
	Description string `form:"description"`
	CreateTime  int    `form:"create_time"`
	UpdateTime  int    `form:"update_time"`
	DeleteTime  int    `form:"delete_time"`
	IsCreate    int    `form:"_create"`
}

// Messages 自定义验证返回消息
func (f UserForm) Messages() map[string]string {
	return validate.MS{
		"UserLevelId.required": "用户等级不能为空.",
		"Username.required":    "用户名不能为空.",
		"Mobile.required":      "手机号不能为空.",
		"Nickname.required":    "昵称不能为空.",
		"Password.required":    "密码不能为空.",
	}
}
