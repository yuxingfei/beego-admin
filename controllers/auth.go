package controllers

import (
	"beego-admin/form_validate"
	"beego-admin/global/response"
	"beego-admin/services/admin_log_service"
	"beego-admin/services/admin_user_service"
	"beego-admin/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/dchest/captcha"
	"github.com/gookit/validate"
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
	this.Data["captcha"] = utils.GetCaptcha()

	this.TplName = "auth/login.tpl"
}

//登录认证
func (this *AuthController)CheckLogin()  {
	//数据校验
	valid := validation.Validation{}
	loginForm := form_validate.LoginForm{}

	if err := this.ParseForm(&loginForm); err != nil{
		response.ErrorWithMessage(err.Error(),this.Ctx)
		return
	}

	v := validate.Struct(loginForm)

	//看是否需要校验验证码
	isCaptcha,_ := beego.AppConfig.Int("login::captcha")
	if isCaptcha > 0{
		valid.Required(loginForm.Captcha,"captcha").Message("请输入验证码.")
		if ok := captcha.VerifyString(loginForm.CaptchaId,loginForm.Captcha); !ok{
			response.ErrorWithMessage("验证码错误.",this.Ctx)
			return
		}
	}

	if !v.Validate(){
		response.ErrorWithMessage(v.Errors.One(),this.Ctx)
		return
	}

	//基础验证通过后，进行用户验证
	loginUser,err := admin_user_service.CheckLogin(loginForm,this.Ctx)
	if err != nil{
		response.ErrorWithMessage(err.Error(),this.Ctx)
		return
	}

	//登录日志记录
	admin_log_service.LoginLog(loginUser.Id,this.Ctx)

	redirect,_ := this.GetSession("redirect").(string)
	if redirect != ""{
		response.SuccessWithMessageAndUrl("登录成功",redirect,this.Ctx)
	}else{
		response.SuccessWithMessageAndUrl("登录成功","/admin/index/index",this.Ctx)
	}

	return
}

//刷新验证码
func (this *AuthController)RefreshCaptcha()  {
	captchaId := this.GetString("captchaId")
	res := map[string]interface{}{
		"isNew":false,
	}
	if captchaId == ""{
		res["msg"] = "参数错误."
	}

	isReload := captcha.Reload(captchaId)
	if isReload{
		res["captchaId"] = captchaId
	}else{
		res["isNew"] = true
		res["captcha"] = utils.GetCaptcha()
	}

	this.Data["json"] = res

	this.ServeJSON()
}