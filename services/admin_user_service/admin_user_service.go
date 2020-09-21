package admin_user_service

import (
	"beego-admin/models"
	"beego-admin/utils"
	"github.com/astaxie/beego/orm"
)

//根据id获取一条admin_user数据
func GetAdminUserById(id int) *models.AdminUser {
	o := orm.NewOrm()
	adminUser := models.AdminUser{Id:id}
	err := o.Read(&adminUser)
	if err != nil{
		return nil
	}
	return &adminUser
}

//权限检测
func AuthCheck(url string,authExcept map[string]interface{},loginUser *models.AdminUser) bool {
	authUrl := loginUser.GetAuthUrl()
	if utils.KeyInMap(url,authExcept) || utils.KeyInMap(url,authUrl){
		return true
	}else {
		return false
	}
}

