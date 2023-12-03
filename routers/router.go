package routers

import (
	"beego-admin/controllers"
	"beego-admin/middleware"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dchest/captcha"
	"net/http"
)

func InitRouter() {
	//授权登录中间件
	middleware.AuthMiddle()

	web.Get("/", func(ctx *context.Context) {
		ctx.Redirect(http.StatusFound, "/admin/index/index")
	})

	//admin模块路由
	admin := web.NewNamespace("/admin",
		//UEditor控制器
		web.NSRouter("/editor/server", &controllers.EditorController{}, "get,post:Server"),

		//登录页
		web.NSRouter("/auth/login", &controllers.AuthController{}, "get:Login"),
		//退出登录
		web.NSRouter("/auth/logout", &controllers.AuthController{}, "get:Logout"),
		//二维码图片输出
		web.NSHandler("/auth/captcha/*.png", captcha.Server(240, 80)),
		//登录认证
		web.NSRouter("/auth/check_login", &controllers.AuthController{}, "post:CheckLogin"),
		//刷新验证码
		web.NSRouter("/auth/refresh_captcha", &controllers.AuthController{}, "post:RefreshCaptcha"),

		//首页
		web.NSRouter("/index/index", &controllers.IndexController{}, "get:Index"),

		//用户管理
		web.NSRouter("/admin_user/index", &controllers.AdminUserController{}, "get:Index"),

		//操作日志
		web.NSRouter("/admin_log/index", &controllers.AdminLogController{}, "get:Index"),

		//菜单管理
		web.NSRouter("/admin_menu/index", &controllers.AdminMenuController{}, "get:Index"),
		//菜单管理-添加菜单-界面
		web.NSRouter("/admin_menu/add", &controllers.AdminMenuController{}, "get:Add"),
		//菜单管理-添加菜单-创建
		web.NSRouter("/admin_menu/create", &controllers.AdminMenuController{}, "post:Create"),
		//菜单管理-修改菜单-界面
		web.NSRouter("/admin_menu/edit", &controllers.AdminMenuController{}, "get:Edit"),
		//菜单管理-更新菜单
		web.NSRouter("/admin_menu/update", &controllers.AdminMenuController{}, "post:Update"),
		//菜单管理-删除菜单
		web.NSRouter("/admin_menu/del", &controllers.AdminMenuController{}, "post:Del"),

		//系统管理-个人资料
		web.NSRouter("/admin_user/profile", &controllers.AdminUserController{}, "get:Profile"),
		//系统管理-个人资料-修改昵称
		web.NSRouter("/admin_user/update_nickname", &controllers.AdminUserController{}, "post:UpdateNickName"),
		//系统管理-个人资料-修改密码
		web.NSRouter("/admin_user/update_password", &controllers.AdminUserController{}, "post:UpdatePassword"),
		//系统管理-个人资料-修改头像
		web.NSRouter("/admin_user/update_avatar", &controllers.AdminUserController{}, "post:UpdateAvatar"),
		//系统管理-用户管理-添加界面
		web.NSRouter("/admin_user/add", &controllers.AdminUserController{}, "get:Add"),
		//系统管理-用户管理-添加
		web.NSRouter("/admin_user/create", &controllers.AdminUserController{}, "post:Create"),
		//系统管理-用户管理-修改界面
		web.NSRouter("/admin_user/edit", &controllers.AdminUserController{}, "get:Edit"),
		//系统管理-用户管理-修改
		web.NSRouter("/admin_user/update", &controllers.AdminUserController{}, "post:Update"),
		//系统管理-用户管理-启用
		web.NSRouter("/admin_user/enable", &controllers.AdminUserController{}, "post:Enable"),
		//系统管理-用户管理-禁用
		web.NSRouter("/admin_user/disable", &controllers.AdminUserController{}, "post:Disable"),
		//系统管理-用户管理-删除
		web.NSRouter("/admin_user/del", &controllers.AdminUserController{}, "post:Del"),

		//系统管理-角色管理
		web.NSRouter("/admin_role/index", &controllers.AdminRoleController{}, "get:Index"),
		//系统管理-角色管理-添加界面
		web.NSRouter("/admin_role/add", &controllers.AdminRoleController{}, "get:Add"),
		//系统管理-角色管理-添加
		web.NSRouter("/admin_role/create", &controllers.AdminRoleController{}, "post:Create"),
		//菜单管理-角色管理-修改界面
		web.NSRouter("/admin_role/edit", &controllers.AdminRoleController{}, "get:Edit"),
		//菜单管理-角色管理-修改
		web.NSRouter("/admin_role/update", &controllers.AdminRoleController{}, "post:Update"),
		//菜单管理-角色管理-删除
		web.NSRouter("/admin_role/del", &controllers.AdminRoleController{}, "post:Del"),
		//菜单管理-角色管理-启用角色
		web.NSRouter("/admin_role/enable", &controllers.AdminRoleController{}, "post:Enable"),
		//菜单管理-角色管理-禁用角色
		web.NSRouter("/admin_role/disable", &controllers.AdminRoleController{}, "post:Disable"),
		//菜单管理-角色管理-角色授权界面
		web.NSRouter("/admin_role/access", &controllers.AdminRoleController{}, "get:Access"),
		//菜单管理-角色管理-角色授权
		web.NSRouter("/admin_role/access_operate", &controllers.AdminRoleController{}, "post:AccessOperate"),

		//设置中心-后台设置
		web.NSRouter("/setting/admin", &controllers.SettingController{}, "get:Admin"),
		//设置中心-更新设置
		web.NSRouter("/setting/update", &controllers.SettingController{}, "post:Update"),

		//系统管理-开发管理-数据维护
		web.NSRouter("/database/table", &controllers.DatabaseController{}, "get:Table"),
		//系统管理-开发管理-数据维护-优化表
		web.NSRouter("/database/optimize", &controllers.DatabaseController{}, "post:Optimize"),
		//系统管理-开发管理-数据维护-修复表
		web.NSRouter("/database/repair", &controllers.DatabaseController{}, "post:Repair"),
		//系统管理-开发管理-数据维护-查看详情
		web.NSRouter("/database/view", &controllers.DatabaseController{}, "get,post:View"),

		//用户等级管理
		web.NSRouter("/user_level/index", &controllers.UserLevelController{}, "get:Index"),
		//用户等级管理-添加界面
		web.NSRouter("/user_level/add", &controllers.UserLevelController{}, "get:Add"),
		//用户等级管理-添加
		web.NSRouter("/user_level/create", &controllers.UserLevelController{}, "post:Create"),
		//用户等级管理-修改界面
		web.NSRouter("/user_level/edit", &controllers.UserLevelController{}, "get:Edit"),
		//用户等级管理-修改
		web.NSRouter("/user_level/update", &controllers.UserLevelController{}, "post:Update"),
		//用户等级管理-启用
		web.NSRouter("/user_level/enable", &controllers.UserLevelController{}, "post:Enable"),
		//用户等级管理-禁用
		web.NSRouter("/user_level/disable", &controllers.UserLevelController{}, "post:Disable"),
		//用户等级管理-删除
		web.NSRouter("/user_level/del", &controllers.UserLevelController{}, "post:Del"),
		//用户等级管理-导出
		web.NSRouter("/user_level/export", &controllers.UserLevelController{}, "get:Export"),

		//用户管理
		web.NSRouter("/user/index", &controllers.UserController{}, "get:Index"),
		//用户管理-添加界面
		web.NSRouter("/user/add", &controllers.UserController{}, "get:Add"),
		//用户管理-添加
		web.NSRouter("/user/create", &controllers.UserController{}, "post:Create"),
		//用户管理-修改界面
		web.NSRouter("/user/edit", &controllers.UserController{}, "get:Edit"),
		//用户管理-修改
		web.NSRouter("/user/update", &controllers.UserController{}, "post:Update"),
		//用户管理-启用
		web.NSRouter("/user/enable", &controllers.UserController{}, "post:Enable"),
		//用户管理-禁用
		web.NSRouter("/user/disable", &controllers.UserController{}, "post:Disable"),
		//用户管理-删除
		web.NSRouter("/user/del", &controllers.UserController{}, "post:Del"),
		//用户管理-导出
		web.NSRouter("/user/export", &controllers.UserController{}, "get:Export"),
	)

	web.AddNamespace(admin)
}
