package services

import (
	"beego-admin/formvalidate"
	"beego-admin/models"
	"beego-admin/utils/page"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strings"
)

// AdminRoleService struct
type AdminRoleService struct {
	BaseService
}

// GetCount 获取admin_role 总数
func (*AdminRoleService) GetCount() int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminRole)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// GetAllData 获取所有admin role
func (*AdminRoleService) GetAllData() []*models.AdminRole {
	var adminRoles []*models.AdminRole
	orm.NewOrm().QueryTable(new(models.AdminRole)).All(&adminRoles)
	return adminRoles
}

// GetPaginateData 分页获取adminrole
func (ars *AdminRoleService) GetPaginateData(listRows int, params url.Values) ([]*models.AdminRole, page.Pagination) {
	//搜索、查询字段赋值
	ars.SearchField = append(ars.SearchField, new(models.AdminRole).SearchField()...)

	var adminRole []*models.AdminRole
	o := orm.NewOrm().QueryTable(new(models.AdminRole))
	_, err := ars.PaginateAndScopeWhere(o, listRows, params).All(&adminRole)
	if err != nil {
		return nil, ars.Pagination
	}
	return adminRole, ars.Pagination
}

// IsExistName 名称验重
func (*AdminRoleService) IsExistName(name string, id int) bool {
	if id == 0 {
		return orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("name", name).Exist()
	}
	return orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("name", name).Exclude("id", id).Exist()
}

// Create 创建角色
func (*AdminRoleService) Create(form *formvalidate.AdminRoleForm) int {
	adminRole := models.AdminRole{
		Name:        form.Name,
		Description: form.Description,
		Url:         "1,2,18",
		Status:      form.Status,
	}

	insertID, err := orm.NewOrm().Insert(&adminRole)
	if err != nil {
		return 0
	}
	return int(insertID)
}

// GetAdminRoleById 通过id获取菜单信息
func (*AdminRoleService) GetAdminRoleById(id int) *models.AdminRole {
	var adminRole models.AdminRole
	err := orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("id", id).One(&adminRole)
	if err == nil {
		return &adminRole
	}
	return nil
}

// Update 更新角色信息
func (*AdminRoleService) Update(form *formvalidate.AdminRoleForm) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("id", form.Id).Update(orm.Params{
		"name":        form.Name,
		"description": form.Description,
		"status":      form.Status,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Del 删除角色
func (*AdminRoleService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

// Enable 启用角色
func (*AdminRoleService) Enable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("id__in", ids).Update(orm.Params{
		"status": 1,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Disable 禁用角色
func (*AdminRoleService) Disable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("id__in", ids).Update(orm.Params{
		"status": 0,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// StoreAccess 授权菜单
func (*AdminRoleService) StoreAccess(id int, url []string) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("id", id).Update(orm.Params{
		"url": strings.Join(url, ","),
	})
	if err == nil {
		return int(num)
	}
	return 0
}
