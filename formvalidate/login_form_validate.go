package formvalidate

import "github.com/gookit/validate"

// LoginForm login 表单
type LoginForm struct {
	Username  string `form:"username" validate:"required"`
	Password  string `form:"password" validate:"required"`
	Captcha   string `form:"captcha"`
	CaptchaId string `form:"captchaId"`
	Remember  string `form:"remember"`
}

// Messages 自定义验证返回消息
func (f LoginForm) Messages() map[string]string {
	return validate.MS{
		"Username.required": "用户名不能为空.",
		"Password.required": "密码不能为空.",
	}
}
