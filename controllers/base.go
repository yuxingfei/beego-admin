package controllers

import (
	"beego-admin/global"
	"beego-admin/models"
	"beego-admin/services/admin_log_service"
	"beego-admin/services/admin_menu_service"
	"fmt"
	"github.com/astaxie/beego"
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

//父控制器初始化
func (this *baseController) Prepare() {
	//访问url
	url := strings.ToLower(strings.TrimLeft(this.Ctx.Input.URL(),"/"))
	//登录用户
	loginUser,_ := this.GetSession(global.LOGIN_USER).(models.AdminUser)

	//基础变量
	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev"{
		this.Data["debug"] = true
	}else{
		this.Data["debug"] = false
	}
	this.Data["cookie_prefix"] = ""
	fmt.Println("base")

	//admin基础配置
	adminConfig,err := beego.AppConfig.GetSection("base")
	if err != nil{
		beego.Error("get base config fail. error:",err)
	}
	//每页预览的数量
	perPage := this.Ctx.GetCookie("admin_per_page")
	if perPage == ""{
		perPage = "10"
	}
	perPageInt,_ := strconv.Atoi(perPage)
	if perPageInt >= 100{
		perPage = "100"
	}


	//记录日志
	adminMenu := admin_menu_service.GetAdminMenuByUrl(url)
	title := "Title"
	if adminMenu != nil{
		title = adminMenu.Name
		if strings.ToLower(adminMenu.LogMethod) == strings.ToLower(this.Ctx.Input.Method()){
			admin_log_service.CreateAdminLog(&loginUser,adminMenu,url,this.Ctx)
		}
	}

	this.Data["admin"] = map[string]interface{}{
		"pjax":this.Ctx.Input.Header("X-PJAX") == "true",
		"user":loginUser,
		"menu":1,
		"name":adminConfig["name"],
		"author":adminConfig["author"],
		"version":adminConfig["version"],
		"short_name":adminConfig["short_name"],
		"per_page":perPage,
		"per_page_config":[]int{10,20,30,50,100},
		"title":title,
	}

	fmt.Println("baseController Prepare")
	fmt.Println("loginUser",loginUser)

	//ajax头部统一设置_xsrf
	this.Data["xsrf_token"] = this.XSRFToken()

	//判断子控制器是否事项了初始化方法
	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
