package middleware

import (
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/models"
	"beego-admin/services"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strconv"
	"strings"
)

//中间件
func AuthMiddle()  {

	//不需要验证的url
	authExcept := map[string]interface{}{
		"admin/auth/login":0,
		"admin/auth/check_login":1,
		"admin/auth/logout":2,
		"admin/auth/captcha":3,
		"admin/editor/server":4,
		"admin/auth/refresh_captcha":5,
	}

	//登录认证中间件过滤器
	var filterLogin = func(ctx *context.Context) {
		url := strings.TrimLeft(ctx.Input.URL(),"/")
		fmt.Println("访问url",url)

		//需要进行登录验证
		if !isAuthExceptUrl(strings.ToLower(url),authExcept){
			//验证是否登录
			loginUser,isLogin := isLogin(ctx)
			if !isLogin{
				response.ErrorWithMessageAndUrl("未登录","/admin/auth/login",ctx)
				return
			}

			//验证，是否有权限访问
			var adminUserService services.AdminUserService
			if loginUser.Id != 1 && !adminUserService.AuthCheck(url,authExcept,loginUser){
				errorBackUrl := global.URL_CURRENT
				if ctx.Request.Method == "GET"{
					errorBackUrl = ""
				}
				response.ErrorWithMessageAndUrl("无权限", errorBackUrl,ctx)
				return
			}
		}

		checkAuth,_ := strconv.Atoi(ctx.Input.Param("check_auth"))
		if checkAuth == 1{
			response.Success(ctx)
			return
		}

	}

	beego.InsertFilter("/admin/*",beego.BeforeRouter,filterLogin)
}

//判断是否是不需要验证登录的url,只针对admin模块路由的判断
func isAuthExceptUrl(url string,m map[string]interface{}) bool {
	urlArr := strings.Split(url,"/")
	if len(urlArr) > 3{
		url = fmt.Sprintf("%s/%s/%s",urlArr[0],urlArr[1],urlArr[2])
	}
	_,ok := m[url]
	if ok {
		return true
	}else{
		return false
	}
}

//是否登录
func isLogin(ctx *context.Context) (*models.AdminUser,bool) {
	loginUser,ok := ctx.Input.Session(global.LOGIN_USER).(models.AdminUser)
	if !ok{
		loginUserIdStr := ctx.GetCookie(global.LOGIN_USER_ID)
		loginUserIdSign := ctx.GetCookie(global.LOGIN_USER_ID_SIGN)

		if loginUserIdStr != "" && loginUserIdSign != ""{
			loginUserId,_ := strconv.Atoi(loginUserIdStr)
			var adminUserService services.AdminUserService
			loginUserPointer := adminUserService.GetAdminUserById(loginUserId)

			if loginUserPointer != nil && loginUserPointer.GetSignStrByAdminUser(ctx) == loginUserIdSign{
				ctx.Output.Session(global.LOGIN_USER,*loginUserPointer)
				return loginUserPointer,true
			}
		}
		return nil,false
	}

	return &loginUser,true
}
