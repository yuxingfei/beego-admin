package services

import (
	"beego-admin/form_validate"
	"beego-admin/models"
	beego_pagination "beego-admin/utils/beego-pagination"
	"github.com/astaxie/beego/orm"
	"net/url"
	"time"
)

type UserLevelService struct {
	BaseService
}

//通过分页获取user_level
func (this *UserLevelService) GetPaginateData(listRows int, params url.Values) ([]*models.UserLevel, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	this.SearchField = append(this.SearchField, new(models.UserLevel).SearchField()...)

	var userLevel []*models.UserLevel
	o := orm.NewOrm().QueryTable(new(models.UserLevel))
	_, err := this.PaginateAndScopeWhere(o, listRows, params).All(&userLevel)
	if err != nil {
		return nil, this.Pagination
	} else {
		return userLevel, this.Pagination
	}
}

//新增用户等级
func (*UserLevelService) Create(form *form_validate.UserLevelForm) int {
	userLevel := models.UserLevel{
		Name:        form.Name,
		Description: form.Description,
		Status:      int8(form.Status),
		CreateTime:  int(time.Now().Unix()),
		UpdateTime:  int(time.Now().Unix()),
	}
	if form.Img != "" {
		userLevel.Img = form.Img
	}
	id, err := orm.NewOrm().Insert(&userLevel)

	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

//更新用户等级
func (*UserLevelService) Update(form *form_validate.UserLevelForm) int {
	o := orm.NewOrm()
	userLevel := models.UserLevel{Id: form.Id}
	if o.Read(&userLevel) == nil {
		userLevel.Name = form.Name
		userLevel.Description = form.Description
		userLevel.Status = int8(form.Status)
		userLevel.UpdateTime = int(time.Now().Unix())
		if form.Img != "" {
			userLevel.Img = form.Img
		}
		userLevel.Name = form.Name
		num, err := o.Update(&userLevel)
		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
}

//根据id获取一条user_level数据
func (*UserLevelService) GetUserLevelById(id int) *models.UserLevel {
	o := orm.NewOrm()
	userLevel := models.UserLevel{Id: id}
	err := o.Read(&userLevel)
	if err != nil {
		return nil
	}
	return &userLevel
}

//启用
func (*UserLevelService) Enable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.UserLevel)).Filter("id__in", ids).Update(orm.Params{
		"status": 1,
	})
	if err == nil {
		return int(num)
	} else {
		return 0
	}
}

//禁用
func (*UserLevelService) Disable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.UserLevel)).Filter("id__in", ids).Update(orm.Params{
		"status": 0,
	})
	if err == nil {
		return int(num)
	} else {
		return 0
	}
}

//删除
func (*UserLevelService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.UserLevel)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}
