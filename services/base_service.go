package services

import (
	beego_pagination "beego-admin/utils/beego-pagination"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type BaseService struct {
	Pagination beego_pagination.Pagination
}

//共用分页
func (this *BaseService)Paginate(seter orm.QuerySeter,listRows int,parameters url.Values) orm.QuerySeter {
	var pagination beego_pagination.Pagination
	qs := pagination.Paginate(seter,listRows,parameters)
	this.Pagination = pagination
	return qs
}
