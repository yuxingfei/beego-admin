package admin_log_service

import (
	"beego-admin/models"
	"beego-admin/utils/encrypter"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"time"
)

//创建操作日志
func CreateAdminLog(loginUser *models.AdminUser,menu *models.AdminMenu,url string,ctx *context.Context)  {
	var adminLog models.AdminLog

	if loginUser == nil{
		adminLog.AdminUserId = 0
	}else{
		adminLog.AdminUserId = loginUser.Id
	}
	adminLog.Name = menu.Name
	adminLog.LogMethod = menu.LogMethod
	adminLog.Url = url
	adminLog.LogIp = ctx.Input.IP()
	adminLog.CreateTime = int(time.Now().Unix())
	adminLog.UpdateTime = int(time.Now().Unix())

	o := orm.NewOrm()
	//开启事务
	err := o.Begin()

	adminLogId,err := o.Insert(&adminLog)
	if err != nil{
		o.Rollback()
	}

	//adminLogData数据表添加数据
	jsonData,_ := json.Marshal(ctx.Input.Params())
	cryptData := encrypter.Encrypt(jsonData,[]byte(beego.AppConfig.String("app_key")))
	var adminLogData models.AdminLogData
	adminLogData.AdminLogId = int(adminLogId)
	adminLogData.Data = cryptData
	_,err = o.Insert(&adminLogData)
	if err != nil{
		o.Rollback()
	}

	o.Commit()
}
