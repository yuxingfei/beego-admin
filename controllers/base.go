package controllers

import (
	"beego-admin/global"
	"beego-admin/models"
	"beego-admin/services"
	beego "github.com/beego/beego/v2/adapter"
	"net/url"
	"strconv"
	"strings"
)

// NestPreparer 定义子控制器初始化方法
type NestPreparer interface {
	NestPrepare()
}

// baseController struct
type baseController struct {
	beego.Controller
}

var (
	//后台变量
	admin map[string]interface{}
	//当前用户
	loginUser models.AdminUser
	//参数
	gQueryParams url.Values
)

// Prepare 父控制器初始化
func (bc *baseController) Prepare() {
	//访问url
	requestURL := strings.ToLower(strings.TrimLeft(bc.Ctx.Input.URL(), "/"))

	//query参数
	//只有分页首页列表时才会使用
	if bc.Ctx.Input.IsGet() {
		gQueryParams, _ = url.ParseQuery(bc.Ctx.Request.URL.RawQuery)
		gQueryParams.Set("queryParamUrl", bc.Ctx.Input.URL())
		if len(gQueryParams) > 0 {
			for k, val := range gQueryParams {
				v, ok := strconv.Atoi(val[0])
				if ok == nil {
					bc.Data[k] = v
				} else {
					bc.Data[k] = val[0]
				}
			}
		}
	}

	//登录用户
	var isOk bool
	loginUser, isOk = bc.GetSession(global.LOGIN_USER).(models.AdminUser)

	//基础变量
	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		bc.Data["debug"] = true
	} else {
		bc.Data["debug"] = false
	}
	bc.Data["cookie_prefix"] = ""

	//每页预览的数量
	perPageStr := bc.Ctx.GetCookie("admin_per_page")
	var perPage int
	if perPageStr == "" {
		perPage = 10
	} else {
		perPage, _ = strconv.Atoi(perPageStr)
	}
	if perPage >= 100 {
		perPage = 100
	}

	//记录日志
	var adminMenuService services.AdminMenuService
	adminMenu := adminMenuService.GetAdminMenuByUrl(requestURL)
	title := ""
	if adminMenu != nil {
		title = adminMenu.Name
		if strings.ToLower(adminMenu.LogMethod) == strings.ToLower(bc.Ctx.Input.Method()) {
			var adminLogService services.AdminLogService
			adminLogService.CreateAdminLog(&loginUser, adminMenu, requestURL, bc.Ctx)
		}
	}

	//左侧菜单
	menu := ""
	if "admin/auth/login" != requestURL && !(bc.Ctx.Input.Header("X-PJAX") == "true") && isOk {
		var adminTreeService services.AdminTreeService
		menu = adminTreeService.GetLeftMenu(requestURL, loginUser)
	}

	admin = map[string]interface{}{
		"pjax":            bc.Ctx.Input.Header("X-PJAX") == "true",
		"user":            &loginUser,
		"menu":            menu,
		"name":            global.BA_CONFIG.Base.Name,
		"author":          global.BA_CONFIG.Base.Author,
		"version":         global.BA_CONFIG.Base.Version,
		"short_name":      global.BA_CONFIG.Base.ShortName,
		"link":            global.BA_CONFIG.Base.Link,
		"per_page":        perPage,
		"per_page_config": []int{10, 20, 30, 50, 100},
		"title":           title,
	}
	bc.Data["admin"] = admin

	//ajax头部统一设置_xsrf
	bc.Data["xsrf_token"] = bc.XSRFToken()

	//判断子控制器是否事项了初始化方法
	if app, ok := bc.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
