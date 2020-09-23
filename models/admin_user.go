package models

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"strings"
)

type AdminUser struct {
	Id         int       `orm:"column(id);auto;size(11)" description:"表ID"`
	Username   string    `orm:"column(username);size(30)" description:"用户名"`
	Password   string    `orm:"column(password);size(255)" description:"密码"`
	Nickname   string    `orm:"column(nickname);size(30)" description:"昵称"`
	Avatar     string    `orm:"column(avatar);size(255)" description:"头像"`
	Role       string    `orm:"column(role);size(200)" description:"角色"`
	Status     int8      `orm:"column(status);size(1)" description:"是否启用 0：否 1：是"`
	DeleteTime int `orm:"column(delete_time);;size(10);default(0)" description:"删除时间"`
}

//自定义table 名称
func (adminUser *AdminUser) TableName() string {
	return "admin_user"
}

//在init中注册定义的model
func init()  {
	orm.RegisterModel(new(AdminUser))
}

//获取加密字符串，用在登录的时候加密处理
func (adminUser *AdminUser)GetSignStrByAdminUser(ctx *context.Context) string {
	ua := ctx.Input.Header("user-agent")
	return fmt.Sprintf("%x",sha1.Sum([]byte(fmt.Sprintf("%d%s%s",adminUser.Id,adminUser.Username,ua))))
}

//获取已授权url
func (adminUser *AdminUser)GetAuthUrl() map[string]interface{} {
	var (
		urlArr orm.ParamsList
	)
	authUrl := make(map[string]interface{})

	o := orm.NewOrm()
	qs := o.QueryTable(new(AdminRole))

	_,err := qs.Filter("id__in",strings.Split(adminUser.Role,",")).Filter("status",1).ValuesFlat(&urlArr,"url")
	if err == nil{
		urlIdStr := ""
		for k, row := range urlArr {
			urlStr,ok := row.(string)
			if ok{
				if k == 0{
					urlIdStr = urlStr
				}else {
					urlIdStr += ","+urlStr
				}
			}
		}
		urlIdArr := strings.Split(urlIdStr,",")

		var authUrlArr orm.ParamsList

		if len(urlIdStr) > 0{
			o = orm.NewOrm()
			qs = o.QueryTable(new(AdminMenu))
			_,err := qs.Filter("id__in",urlIdArr).ValuesFlat(&authUrlArr,"url")
			if err == nil{
				for k,row := range authUrlArr{
					val,ok := row.(string)
					if ok{
						authUrl[val] = k
					}
				}
			}
		}
		return authUrl
	}
	return authUrl
}
