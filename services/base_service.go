package services

import (
	beego_pagination "beego-admin/utils/beego-pagination"
	"fmt"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type BaseService struct {
	//可搜索字段
	SearchField []string
	//可作为条件的字段
	WhereField []string
	//可做为时间范围查询的字段
	TimeField []string
	//禁止删除的数据id
	NoDeletionId []int
	//分页
	Pagination beego_pagination.Pagination
}

//分页处理
func (this *BaseService)Paginate(seter orm.QuerySeter,listRows int,parameters url.Values) orm.QuerySeter {
	var pagination beego_pagination.Pagination
	qs := pagination.Paginate(seter,listRows,parameters)
	this.Pagination = pagination
	return qs
}

//查询处理
func (this *BaseService) ScopeWhere(seter orm.QuerySeter) orm.QuerySeter {
	fmt.Println("this.WhereField = ",this.WhereField)
	return seter
}

//分页和查询合并，多用于首页列表展示、搜索
func (this *BaseService)PaginateAndScopeWhere(seter orm.QuerySeter,listRows int,parameters url.Values) orm.QuerySeter {
	//关键词like搜索

	//字段条件查询

	//时间范围查询

	//排序

	return this.ScopeWhere(this.Paginate(seter,listRows,parameters))
}
