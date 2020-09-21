package controllers

import (
	"beego-admin/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type AuthController struct {
	baseController
}

//控制器，初始化函数，基础自父控制器
func (this *AuthController) NestPrepare() {
	fmt.Println("AuthController NestPrepare")
}

//登录界面
func (this *AuthController) Login()  {
	//获取登录配置信息
	loginConfig := struct {
		Token string
		Captcha string
		Background string
	}{
		Token:beego.AppConfig.DefaultString("login::token","1"),
		Captcha:beego.AppConfig.DefaultString("login::captcha","1"),
		Background:beego.AppConfig.DefaultString("login::background","/static/admin/images/default_background.jpeg"),
	}
	this.Data["login_config"] = loginConfig

	//登录验证码
	this.Data["captcha_id"] = utils.CaptchaId()

	this.TplName = "auth/login.tpl"
}