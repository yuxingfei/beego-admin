package services

import (
	"beego-admin/form_validate"
	"beego-admin/models"
	beego_pagination "beego-admin/utils/beego-pagination"
	"github.com/astaxie/beego/orm"
	"net/url"
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

//获取所有adminrole
func (this *AdminRoleService) GetAllData(listRows int, params url.Values) ([]*models.AdminRole, beego_pagination.Pagination) {
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
		Name:form.Name,
		Description:form.Description,
		Url:"1,2,18",
		Status:form.Status,
	}

	insertId, err := orm.NewOrm().Insert(&adminRole)
	if err != nil{
		return 0
	}else {
		return int(insertId)
	}
}