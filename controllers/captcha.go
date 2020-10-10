package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

type CaptchaController struct {
	beego.Controller
}

func (this *CaptchaController) CaptchaId() {
	captchaId := captcha.NewLen(6)
	this.Ctx.WriteString(captchaId)
}
