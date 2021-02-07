package controllers

import (
	"beego-admin/formvalidate"
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/services"
	"beego-admin/utils"
	"github.com/beego/beego/v2/adapter/validation"
	"github.com/dchest/captcha"
	"github.com/gookit/validate"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var adminLogService services.AdminLogService

// AuthController struct
type AuthController struct {
	baseController
}

// Login 登录界面
func (ac *AuthController) Login() {
	//加载登录配置信息
	var settingService services.SettingService
	data := settingService.Show(1)
	for _,setting := range data {
		settingService.LoadOrUpdateGlobalBaseConfig(setting)
	}

	//获取登录配置信息
	loginConfig := struct {
		Token      string
		Captcha    string
		Background string
	}{
		Token:      global.BA_CONFIG.Login.Token,
		Captcha:    global.BA_CONFIG.Login.Captcha,
	}
	//登录背景图片
	if _,err := os.Stat(strings.TrimLeft(global.BA_CONFIG.Login.Background,"/")); err != nil{
		global.BA_CONFIG.Login.Background = "/static/admin/images/login-default-bg.jpg"
	}
	loginConfig.Background = global.BA_CONFIG.Login.Background

	//login界面只需要name字段
	admin := map[string]interface{}{
		"name":            global.BA_CONFIG.Base.Name,
		"title":           "登录",
	}

	ac.Data["login_config"] = loginConfig
	//登录验证码
	ac.Data["captcha"] = utils.GetCaptcha()
	ac.Data["admin"] = admin

	ac.TplName = "auth/login.html"
}

// Logout 退出登录
func (ac *AuthController) Logout() {
	ac.DelSession(global.LOGIN_USER)
	ac.Ctx.SetCookie(global.LOGIN_USER_ID, "", -1)
	ac.Ctx.SetCookie(global.LOGIN_USER_ID_SIGN, "", -1)
	ac.Redirect("/admin/auth/login", http.StatusFound)
}

// CheckLogin 登录认证
func (ac *AuthController) CheckLogin() {
	//数据校验
	valid := validation.Validation{}
	loginForm := formvalidate.LoginForm{}

	if err := ac.ParseForm(&loginForm); err != nil {
		response.ErrorWithMessage(err.Error(), ac.Ctx)
	}

	v := validate.Struct(loginForm)

	//看是否需要校验验证码
	isCaptcha, _ := strconv.Atoi(global.BA_CONFIG.Login.Captcha)
	if isCaptcha > 0 {
		valid.Required(loginForm.Captcha, "captcha").Message("请输入验证码.")
		if ok := captcha.VerifyString(loginForm.CaptchaId, loginForm.Captcha); !ok {
			response.ErrorWithMessage("验证码错误.", ac.Ctx)
		}
	}

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), ac.Ctx)
	}

	//基础验证通过后，进行用户验证
	var adminUserService services.AdminUserService
	loginUser, err := adminUserService.CheckLogin(loginForm, ac.Ctx)
	if err != nil {
		response.ErrorWithMessage(err.Error(), ac.Ctx)
	}

	//登录日志记录
	adminLogService.LoginLog(loginUser.Id, ac.Ctx)

	redirect, _ := ac.GetSession("redirect").(string)
	if redirect != "" {
		response.SuccessWithMessageAndUrl("登录成功", redirect, ac.Ctx)
	} else {
		response.SuccessWithMessageAndUrl("登录成功", "/admin/index/index", ac.Ctx)
	}
}

// RefreshCaptcha 刷新验证码
func (ac *AuthController) RefreshCaptcha() {
	captchaID := ac.GetString("captchaId")
	res := map[string]interface{}{
		"isNew": false,
	}
	if captchaID == "" {
		res["msg"] = "参数错误."
	}

	isReload := captcha.Reload(captchaID)
	if isReload {
		res["captchaId"] = captchaID
	} else {
		res["isNew"] = true
		res["captcha"] = utils.GetCaptcha()
	}

	ac.Data["json"] = res

	ac.ServeJSON()
}
