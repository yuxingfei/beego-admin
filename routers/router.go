package routers

import (
	"beego-admin/controllers"
	"beego-admin/middleware"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dchest/captcha"
	"net/http"
)

func init() {
	//授权登录中间件
	middleware.AuthMiddle()

	beego.Get("/", func(ctx *context.Context) {
		ctx.Redirect(http.StatusFound, "/admin/index/index")
	})

	//admin模块路由
	admin := beego.NewNamespace("/admin",
		//登录页
		beego.NSRouter("/auth/login", &controllers.AuthController{}, "get:Login"),
		//退出登录
		beego.NSRouter("/auth/logout", &controllers.AuthController{}, "get:Logout"),
		//二维码图片输出
		beego.NSHandler("/auth/captcha/*.png", captcha.Server(240, 80)),
		//登录认证
		beego.NSRouter("/auth/check_login", &controllers.AuthController{}, "post:CheckLogin"),
		//刷新验证码
		beego.NSRouter("/auth/refresh_captcha", &controllers.AuthController{}, "post:RefreshCaptcha"),

		//首页
		beego.NSRouter("/index/index", &controllers.IndexController{}, "get:Index"),

		//用户管理
		beego.NSRouter("/admin_user/index", &controllers.AdminUserController{}, "get:Index"),

		//操作日志
		beego.NSRouter("/admin_log/index", &controllers.AdminLogController{}, "get:Index"),

		//菜单管理
		beego.NSRouter("/admin_menu/index", &controllers.AdminMenuController{}, "get:Index"),
		//菜单管理-添加菜单-界面
		beego.NSRouter("/admin_menu/add", &controllers.AdminMenuController{}, "get:Add"),
		//菜单管理-添加菜单-创建
		beego.NSRouter("/admin_menu/create", &controllers.AdminMenuController{}, "post:Create"),
		//菜单管理-修改菜单-界面
		beego.NSRouter("/admin_menu/edit", &controllers.AdminMenuController{}, "get:Edit"),
		//菜单管理-更新菜单
		beego.NSRouter("/admin_menu/update", &controllers.AdminMenuController{}, "post:Update"),
		//菜单管理-删除菜单
		beego.NSRouter("/admin_menu/del", &controllers.AdminMenuController{}, "post:Del"),

		//系统管理-个人资料
		beego.NSRouter("/admin_user/profile", &controllers.AdminUserController{}, "get:Profile"),
		//系统管理-个人资料-修改昵称
		beego.NSRouter("/admin_user/update_nickname", &controllers.AdminUserController{}, "post:UpdateNickName"),
		//系统管理-个人资料-修改密码
		beego.NSRouter("/admin_user/update_password", &controllers.AdminUserController{}, "post:UpdatePassword"),
		//系统管理-个人资料-修改头像
		beego.NSRouter("/admin_user/update_avatar", &controllers.AdminUserController{}, "post:UpdateAvatar"),
	)

	beego.AddNamespace(admin)
}
