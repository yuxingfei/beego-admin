package controllers

import (
	"beego-admin/global"
	"beego-admin/services"
	"beego-admin/utils"
	"bufio"
	"encoding/base64"
	beego "github.com/beego/beego/v2/adapter"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

// IndexController struct
type IndexController struct {
	baseController
}

// PackageLib 加载的包信息
type PackageLib struct {
	Name    string
	Version string
}

// Index 首页
func (ic *IndexController) Index() {
	ic.Data["login_user"] = ic.User

	//默认密码修改检测
	ic.Data["password_danger"] = 0

	//是否首页显示提示信息
	ic.Data["show_notice"] = global.BA_CONFIG.Base.ShowNotice
	//提示内容
	ic.Data["notice_content"] = global.BA_CONFIG.Base.NoticeContent

	//默认密码修改检测
	loginUserPassword, _ := base64.StdEncoding.DecodeString(ic.User.Password)
	if global.BA_CONFIG.Base.PasswordWarning == "1" && utils.PasswordVerify("123456", string(loginUserPassword)) {
		ic.Data["password_danger"] = 1
	}

	//后台用户数量
	var adminUserService services.AdminUserService
	ic.Data["admin_user_count"] = adminUserService.GetCount()
	//后台角色数量
	var adminRoleService services.AdminRoleService
	ic.Data["admin_role_count"] = adminRoleService.GetCount()
	//后台菜单数量
	var adminMenuService services.AdminMenuService
	ic.Data["admin_menu_count"] = adminMenuService.GetCount()
	//后台日志数量
	var adminLogService services.AdminLogService
	ic.Data["admin_log_count"] = adminLogService.GetCount()
	//系统信息
	ic.Data["system_info"] = ic.getSystemInfo()

	ic.Layout = "public/base.html"

	ic.TplName = "index/index.html"
}

// getSystemInfo 获取系统信息
func (ic *IndexController) getSystemInfo() map[string]interface{} {
	systemInfo := make(map[string]interface{})
	//服务器系统
	systemInfo["server_os"] = runtime.GOOS
	//Go版本
	systemInfo["go_version"] = runtime.Version()
	//文件上传默认内存缓存大小
	systemInfo["upload_file_max_memory"] = int(beego.BConfig.MaxMemory / 1024 / 1024)
	//beego版本
	systemInfo["beego_version"] = beego.VERSION
	//当前后台版本
	systemInfo["admin_version"] = global.BA_CONFIG.Base.Version
	//mysql版本
	var databaseService services.DatabaseService
	systemInfo["db_version"] = databaseService.GetMysqlVersion()
	//go时区
	systemInfo["timezone"] = time.UTC
	//当前时间
	systemInfo["date_time"] = time.Now().Format("2006-01-02 15:04:05")
	//用户IP
	systemInfo["user_ip"] = ic.Ctx.Input.IP()

	userAgent := ic.Ctx.Input.Header("user-agent")

	userOs := "Other"
	if strings.Contains(userAgent, "win") {
		userOs = "Windows"
	} else if strings.Contains(userAgent, "mac") {
		userOs = "MAC"
	} else if strings.Contains(userAgent, "linux") {
		userOs = "Linux"
	} else if strings.Contains(userAgent, "unix") {
		userOs = "Unix"
	} else if strings.Contains(userAgent, "bsd") {
		userOs = "BSD"
	} else if strings.Contains(userAgent, "iPad") || strings.Contains(userAgent, "iPhone") {
		userOs = "IOS"
	} else if strings.Contains(userAgent, "android") {
		userOs = "Android"
	}

	userBrowser := "Other"
	if strings.Contains(userAgent, "MSIE") {
		userBrowser = "MSIE"
	} else if strings.Contains(userAgent, "Firefox") {
		userBrowser = "Firefox"
	} else if strings.Contains(userAgent, "Chrome") {
		userBrowser = "Chrome"
	} else if strings.Contains(userAgent, "Safari") {
		userBrowser = "Safari"
	} else if strings.Contains(userAgent, "Opera") {
		userBrowser = "Opera"
	}

	//用户系统
	systemInfo["user_os"] = userOs
	//用户浏览器
	systemInfo["user_browser"] = userBrowser

	//读取go.mod文件
	var requireList []*PackageLib
	srcFile, err := os.Open("go.mod")
	if err != nil {
		beego.Error(err)
	} else {
		defer srcFile.Close()
		reader := bufio.NewReader(srcFile)
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				continue
			}
			if string(line) != "" {
				strArr := strings.Split(strings.TrimSpace(string(line)), " ")
				lenStrArr := len(strArr)
				//常规require方式
				if strArr[0] == "require" && lenStrArr >= 3 {
					packageLib := PackageLib{
						Name:    strArr[1],
						Version: strArr[2],
					}
					requireList = append(requireList, &packageLib)
				} else {
					//require多个时候
					if lenStrArr >= 2 && strings.Contains(strArr[0], "/") {
						packageLib := PackageLib{
							Name:    strArr[0],
							Version: strArr[1],
						}
						requireList = append(requireList, &packageLib)
					}
				}
			}
		}
	}

	systemInfo["require_list"] = requireList

	return systemInfo
}
