package models

import (
	"crypto/sha1"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"strings"
)

// AdminUser struct
type AdminUser struct {
	Id         int    `orm:"column(id);auto;size(11)" description:"表ID" json:"id"`
	Username   string `orm:"column(username);size(30)" description:"用户名" json:"username"`
	Password   string `orm:"column(password);size(255)" description:"密码" json:"password"`
	Nickname   string `orm:"column(nickname);size(30)" description:"昵称" json:"nickname"`
	Avatar     string `orm:"column(avatar);size(255)" description:"头像" json:"avatar"`
	Role       string `orm:"column(role);size(200)" description:"角色" json:"role"`
	Status     int8   `orm:"column(status);size(1)" description:"是否启用 0：否 1：是" json:"status"`
	DeleteTime int    `orm:"column(delete_time);;size(10);default(0)" description:"删除时间" json:"delete_time"`
}

// TableName 自定义table 名称
func (*AdminUser) TableName() string {
	return "admin_user"
}

// SearchField 定义模型的可搜索字段
func (*AdminUser) SearchField() []string {
	return []string{"nickname", "username"}
}

// NoDeletionId 禁止删除的数据id
func (*AdminUser) NoDeletionId() []int {
	return []int{}
}

// WhereField 定义模型可作为条件的字段
func (*AdminUser) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*AdminUser) TimeField() []string {
	return []string{}
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(AdminUser))
}

// GetSignStrByAdminUser 获取加密字符串，用在登录的时候加密处理
func (adminUser *AdminUser) GetSignStrByAdminUser(ctx *context.Context) string {
	ua := ctx.Input.Header("user-agent")
	return fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%d%s%s", adminUser.Id, adminUser.Username, ua))))
}

// GetAuthUrl 获取已授权url
func (adminUser *AdminUser) GetAuthUrl() map[string]interface{} {
	var (
		urlArr orm.ParamsList
	)
	authURL := make(map[string]interface{})

	o := orm.NewOrm()
	qs := o.QueryTable(new(AdminRole))

	_, err := qs.Filter("id__in", strings.Split(adminUser.Role, ",")).Filter("status", 1).ValuesFlat(&urlArr, "url")
	if err == nil {
		urlIDStr := ""
		for k, row := range urlArr {
			urlStr, ok := row.(string)
			if ok {
				if k == 0 {
					urlIDStr = urlStr
				} else {
					urlIDStr += "," + urlStr
				}
			}
		}
		urlIDArr := strings.Split(urlIDStr, ",")

		var authURLArr orm.ParamsList

		if len(urlIDStr) > 0 {
			o = orm.NewOrm()
			qs = o.QueryTable(new(AdminMenu))
			_, err := qs.Filter("id__in", urlIDArr).ValuesFlat(&authURLArr, "url")
			if err == nil {
				for k, row := range authURLArr {
					val, ok := row.(string)
					if ok {
						authURL[val] = k
					}
				}
			}
		}
		return authURL
	}
	return authURL
}

// GetShowMenu 获取当前用户已授权的显示菜单
func (adminUser *AdminUser) GetShowMenu() map[int]orm.Params {
	var maps []orm.Params
	returnMaps := make(map[int]orm.Params)
	o := orm.NewOrm()

	if adminUser.Id == 1 {
		_, err := o.QueryTable(new(AdminMenu)).Filter("is_show", 1).OrderBy("sort_id", "id").Values(&maps, "id", "parent_id", "name", "url", "icon", "sort_id")
		if err == nil {
			for _, m := range maps {
				returnMaps[int(m["Id"].(int64))] = m
			}
			return returnMaps
		}
		return map[int]orm.Params{}
	}

	var list orm.ParamsList
	_, err := o.QueryTable(new(AdminRole)).Filter("id__in", strings.Split(adminUser.Role, ",")).Filter("status", 1).ValuesFlat(&list, "url")
	if err == nil {
		var urlIDArr []string
		for _, m := range list {
			urlIDArr = append(urlIDArr, strings.Split(m.(string), ",")...)
		}
		_, err := o.QueryTable(new(AdminMenu)).Filter("id__in", urlIDArr).Filter("is_show", 1).OrderBy("sort_id", "id").Values(&maps, "id", "parent_id", "name", "url", "icon", "sort_id")
		if err == nil {
			for _, m := range maps {
				returnMaps[int(m["Id"].(int64))] = m
			}
			return returnMaps
		}
		return map[int]orm.Params{}
	}
	return map[int]orm.Params{}

}

// GetRoleText 用户角色名称
func (adminUser *AdminUser) GetRoleText() map[int]*AdminRole {
	roleIDArr := strings.Split(adminUser.Role, ",")
	var adminRole []*AdminRole
	_, err := orm.NewOrm().QueryTable(new(AdminRole)).Filter("id__in", roleIDArr, "id", "name").All(&adminRole)
	if err != nil {
		return nil
	}
	adminRoleMap := make(map[int]*AdminRole)
	for _, v := range adminRole {
		adminRoleMap[v.Id] = v
	}
	return adminRoleMap
}

// GetAdminUser 获取所有用户
func (*AdminUser) GetAdminUser() []*AdminUser {
	var adminUsers []*AdminUser
	_, err := orm.NewOrm().QueryTable(new(AdminUser)).All(&adminUsers)
	if err == nil {
		return adminUsers
	}
	return nil
}
