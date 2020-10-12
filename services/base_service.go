package services

import (
	utils2 "beego-admin/utils"
	beego_pagination "beego-admin/utils/beego-pagination"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strings"
	"time"
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
func (this *BaseService) Paginate(seter orm.QuerySeter, listRows int, parameters url.Values) orm.QuerySeter {
	var pagination beego_pagination.Pagination
	qs := pagination.Paginate(seter, listRows, parameters)
	this.Pagination = pagination
	return qs
}

//查询处理
func (this *BaseService) ScopeWhere(seter orm.QuerySeter, parameters url.Values) orm.QuerySeter {

	//关键词like搜索
	keywords := parameters.Get("_keywords")
	cond := orm.NewCondition()
	if keywords != "" && len(this.SearchField) > 0 {
		for _, v := range this.SearchField {
			cond = cond.Or(v+"__icontains", keywords)
		}
	}

	//字段条件查询
	if len(this.WhereField) > 0 && len(parameters) > 0 {
		for k, v := range parameters {
			if v[0] != "" && utils2.KeyInArrayForString(this.WhereField, k) {
				cond = cond.And(k, v[0])
			}
		}
	}

	//时间范围查询
	if len(this.TimeField) > 0 && len(parameters) > 0 {
		for key, value := range parameters {
			if value[0] != "" && utils2.KeyInArrayForString(this.TimeField, key) {
				timeRange := strings.Split(value[0], " - ")
				startTimeStr := timeRange[0]
				endTimeStr := timeRange[1]

				loc, _ := time.LoadLocation("Local")
				startTime, err := time.ParseInLocation("2006-01-02 15:04:05", startTimeStr, loc)

				if err == nil {
					unixStartTime := startTime.Unix()
					if len(endTimeStr) == 10 {
						endTimeStr += "23:59:59"
					}

					endTime, err := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, loc)
					if err == nil {
						unixEndTime := endTime.Unix()
						cond = cond.And(key+"__gte", unixStartTime).And(key+"__lte", unixEndTime)
					}
				}
			}
		}
	}

	//将条件语句拼装到主语句中
	seter = seter.SetCond(cond)

	//排序
	order := parameters.Get("_order")
	by := parameters.Get("_by")
	if order == "" {
		order = "id"
	}
	if by == "" {
		by = "-"
	}

	//排序
	seter = seter.OrderBy(by + order)

	return seter
}

//分页和查询合并，多用于首页列表展示、搜索
func (this *BaseService) PaginateAndScopeWhere(seter orm.QuerySeter, listRows int, parameters url.Values) orm.QuerySeter {
	return this.Paginate(this.ScopeWhere(seter, parameters), listRows, parameters)
}
