package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

// CaptchaController struct
type CaptchaController struct {
	beego.Controller
}

// CaptchaId 获取CaptchaId
func (cc *CaptchaController) CaptchaId() {
	captchaID := captcha.NewLen(6)
	cc.Ctx.WriteString(captchaID)
}
