package controllers

import (
	"beego-admin/global"
	"beego-admin/models"
	"beego-admin/services"
	"fmt"
	"github.com/astaxie/beego"
	"net/url"
	"strconv"
	"strings"
)

//定义子控制器初始化方法
type NestPreparer interface {
	NestPrepare()
}

type baseController struct {
	beego.Controller
}

var (
	//后台变量
	admin map[string]interface{}
	//当前用户
	loginUser models.AdminUser
	//参数
	queryParams url.Values
)

//父控制器初始化
func (this *baseController) Prepare() {
	//访问url
	url := strings.ToLower(strings.TrimLeft(this.Ctx.Input.URL(), "/"))

	//query参数
	queryParams = this.Input()
	queryParams.Set("url", this.Ctx.Input.URL())
	fmt.Println("queryParams = ", queryParams)
	if len(queryParams) > 0 {
		for k, val := range queryParams {
			v, ok := strconv.Atoi(val[0])
			if ok == nil {
				this.Data[k] = v
			} else {
				this.Data[k] = val[0]
			}
		}
	}

	//登录用户
	var isOk bool
	loginUser, isOk = this.GetSession(global.LOGIN_USER).(models.AdminUser)

	//基础变量
	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		this.Data["debug"] = true
	} else {
		this.Data["debug"] = false
	}
	this.Data["cookie_prefix"] = ""

	//admin基础配置
	adminConfig, err := beego.AppConfig.GetSection("base")
	if err != nil {
		beego.Error("get base config fail. error:", err)
	}
	//每页预览的数量
	perPageStr := this.Ctx.GetCookie("admin_per_page")
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
	adminMenu := adminMenuService.GetAdminMenuByUrl(url)
	title := ""
	if adminMenu != nil {
		title = adminMenu.Name
		if strings.ToLower(adminMenu.LogMethod) == strings.ToLower(this.Ctx.Input.Method()) {
			var adminLogService services.AdminLogService
			adminLogService.CreateAdminLog(&loginUser, adminMenu, url, this.Ctx)
		}
	}

	//左侧菜单
	menu := ""
	if "admin/auth/login" != url && !(this.Ctx.Input.Header("X-PJAX") == "true") && isOk {
		var adminTreeService services.AdminTreeService
		menu = adminTreeService.GetLeftMenu(url, loginUser)
	}

	admin = map[string]interface{}{
		"pjax":            this.Ctx.Input.Header("X-PJAX") == "true",
		"user":            &loginUser,
		"menu":            menu,
		"name":            adminConfig["name"],
		"author":          adminConfig["author"],
		"version":         adminConfig["version"],
		"short_name":      adminConfig["short_name"],
		"link":            adminConfig["link"],
		"per_page":        perPage,
		"per_page_config": []int{10, 20, 30, 50, 100},
		"title":           title,
	}
	this.Data["admin"] = admin

	//ajax头部统一设置_xsrf
	this.Data["xsrf_token"] = this.XSRFToken()

	//判断子控制器是否事项了初始化方法
	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
