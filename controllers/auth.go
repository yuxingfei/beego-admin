package controllers

import (
	"beego-admin/global/response"
	"beego-admin/utils"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
)

type AuthController struct {
	baseController
}

//login 表单
type LoginForm struct {
	Username  string `form:"username"`
	Password  string `form:"password"`
	Captcha   string `form:"captcha"`
	CaptchaId string `form:"captchaId"`
	Remember  string `form:"remember"`
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
	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())

	this.TplName = "auth/login.tpl"
}

//登录认证
func (this *AuthController)CheckLogin()  {
	//数据校验
	loginForm := LoginForm{}
	if err := this.ParseForm(&loginForm); err != nil{
		response.ErrorWithMessage(err.Error(),this.Ctx)
		return
	}

	fmt.Println(loginForm)
	this.Data["json"] = loginForm
	this.ServeJSON()
}