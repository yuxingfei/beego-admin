package controllers

import (
	"beego-admin/form_validate"
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/services"
	"beego-admin/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/dchest/captcha"
	"github.com/gookit/validate"
	"gopkg.in/ini.v1"
	"net/http"
)

var adminLogService services.AdminLogService

type AuthController struct {
	baseController
}

//控制器，初始化函数，基础自父控制器
func (this *AuthController) NestPrepare() {
	//fmt.Println("AuthController NestPrepare")
}

//登录界面
func (this *AuthController) Login() {

	//获取登录配置信息,不适用beego的获取配置方式
	//beego获取配置会写入缓存，只有程序重启后，才会重新加载配置信息
	cfg, err := ini.Load("conf/admin.conf")

	var loginConfig struct {
		Token      string
		Captcha    string
		Background string
	}

	if err != nil {
		fmt.Printf("Fail to read file conf/admin.conf: %v", err)
		loginConfig.Token = "1"
		loginConfig.Captcha = "1"
		loginConfig.Background = "/static/admin/images/default_background.jpeg"
	} else {
		loginConfig.Token = cfg.Section("login").Key("token").String()
		loginConfig.Captcha = cfg.Section("login").Key("captcha").String()
		loginConfig.Background = cfg.Section("login").Key("background").String()
	}

	this.Data["login_config"] = loginConfig

	//登录验证码
	this.Data["captcha"] = utils.GetCaptcha()

	this.TplName = "auth/login.html"
}

//退出登录
func (this *AuthController) Logout() {
	this.DelSession(global.LOGIN_USER)
	this.Ctx.SetCookie(global.LOGIN_USER_ID, "", -1)
	this.Ctx.SetCookie(global.LOGIN_USER_ID_SIGN, "", -1)
	this.Redirect("/admin/auth/login", http.StatusFound)
}

//登录认证
func (this *AuthController) CheckLogin() {
	//数据校验
	valid := validation.Validation{}
	loginForm := form_validate.LoginForm{}

	if err := this.ParseForm(&loginForm); err != nil {
		response.ErrorWithMessage(err.Error(), this.Ctx)
	}

	v := validate.Struct(loginForm)

	//看是否需要校验验证码
	isCaptcha, _ := beego.AppConfig.Int("login::captcha")
	if isCaptcha > 0 {
		valid.Required(loginForm.Captcha, "captcha").Message("请输入验证码.")
		if ok := captcha.VerifyString(loginForm.CaptchaId, loginForm.Captcha); !ok {
			response.ErrorWithMessage("验证码错误.", this.Ctx)
		}
	}

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), this.Ctx)
	}

	//基础验证通过后，进行用户验证
	var adminUserService services.AdminUserService
	loginUser, err := adminUserService.CheckLogin(loginForm, this.Ctx)
	if err != nil {
		response.ErrorWithMessage(err.Error(), this.Ctx)
	}

	//登录日志记录
	adminLogService.LoginLog(loginUser.Id, this.Ctx)

	redirect, _ := this.GetSession("redirect").(string)
	if redirect != "" {
		response.SuccessWithMessageAndUrl("登录成功", redirect, this.Ctx)
	} else {
		response.SuccessWithMessageAndUrl("登录成功", "/admin/index/index", this.Ctx)
	}
}

//刷新验证码
func (this *AuthController) RefreshCaptcha() {
	captchaId := this.GetString("captchaId")
	res := map[string]interface{}{
		"isNew": false,
	}
	if captchaId == "" {
		res["msg"] = "参数错误."
	}

	isReload := captcha.Reload(captchaId)
	if isReload {
		res["captchaId"] = captchaId
	} else {
		res["isNew"] = true
		res["captcha"] = utils.GetCaptcha()
	}

	this.Data["json"] = res

	this.ServeJSON()
}
