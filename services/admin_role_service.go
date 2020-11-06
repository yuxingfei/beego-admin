package services

import (
	"beego-admin/form_validate"
	"beego-admin/models"
	beego_pagination "beego-admin/utils/beego-pagination"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strings"
)

type AdminRoleService struct {
	BaseService
}

//获取admin_role 总数
func (*AdminRoleService) GetCount() int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminRole)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

//获取所有admin role
func (this *AdminRoleService) GetAllData() []*models.AdminRole {
	var adminRoles []*models.AdminRole
	orm.NewOrm().QueryTable(new(models.AdminRole)).All(&adminRoles)
	return adminRoles
}

//分页获取adminrole
func (this *AdminRoleService) GetPaginateData(listRows int, params url.Values) ([]*models.AdminRole, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	this.SearchField = append(this.SearchField, new(models.AdminRole).SearchField()...)

	var adminRole []*models.AdminRole
	o := orm.NewOrm().QueryTable(new(models.AdminRole))
	_, err := this.PaginateAndScopeWhere(o, listRows, params).All(&adminRole)
	if err != nil {
		return nil, this.Pagination
	} else {
		return adminRole, this.Pagination
	}
}

//名称验重
func (*AdminRoleService) IsExistName(name string, id int) bool {
	if id == 0 {
		return orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("name", name).Exist()
	} else {
		return orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("name", name).Exclude("id", id).Exist()
	}
}

//创建角色
func (*AdminRoleService) Create(form *form_validate.AdminRoleForm) int {
	adminRole := models.AdminRole{
		Name:        form.Name,
		Description: form.Description,
		Url:         "1,2,18",
		Status:      form.Status,
	}

	insertId, err := orm.NewOrm().Insert(&adminRole)
	if err != nil {
		return 0
	} else {
		return int(insertId)
	}
}

//通过id获取菜单信息
func (*AdminRoleService) GetAdminRoleById(id int) *models.AdminRole {
	var adminRole models.AdminRole
	err := orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("id", id).One(&adminRole)
	if err == nil {
		return &adminRole
	} else {
		return nil
	}
}

//更新角色信息
func (*AdminRoleService) Update(form *form_validate.AdminRoleForm) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("id", form.Id).Update(orm.Params{
		"name":        form.Name,
		"description": form.Description,
		"status":      form.Status,
	})
	if err == nil {
		return int(num)
	} else {
		return 0
	}
}

//删除角色
func (*AdminRoleService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}

//启用角色
func (*AdminRoleService) Enable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("id__in", ids).Update(orm.Params{
		"status": 1,
	})
	if err == nil {
		return int(num)
	} else {
		return 0
	}
}

//禁用角色
func (*AdminRoleService) Disable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("id__in", ids).Update(orm.Params{
		"status": 0,
	})
	if err == nil {
		return int(num)
	} else {
		return 0
	}
}

//授权菜单
func (*AdminRoleService) StoreAccess(id int, url []string) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("id", id).Update(orm.Params{
		"url": strings.Join(url, ","),
	})
	if err == nil {
		return int(num)
	} else {
		return 0
	}
}
