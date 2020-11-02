package services

import (
	"beego-admin/models"
	beego_pagination "beego-admin/utils/beego-pagination"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type UserService struct {
	BaseService
}

//通过分页获取user
func (this *UserService) GetPaginateData(listRows int, params url.Values) ([]*models.User, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	this.SearchField = append(this.SearchField, new(models.User).SearchField()...)

	var users []*models.User
	o := orm.NewOrm().QueryTable(new(models.User))
	_, err := this.PaginateAndScopeWhere(o, listRows, params).All(&users)
	if err != nil {
		return nil, this.Pagination
	} else {
		return users, this.Pagination
	}
}
