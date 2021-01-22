package services

import (
	"beego-admin/formvalidate"
	"beego-admin/models"
	"beego-admin/utils/page"
	"github.com/astaxie/beego/orm"
	"net/url"
	"time"
)

// UserLevelService struct
type UserLevelService struct {
	BaseService
}

// GetPaginateData 通过分页获取user_level
func (uls *UserLevelService) GetPaginateData(listRows int, params url.Values) ([]*models.UserLevel, page.Pagination) {
	//搜索、查询字段赋值
	uls.SearchField = append(uls.SearchField, new(models.UserLevel).SearchField()...)

	var userLevel []*models.UserLevel
	o := orm.NewOrm().QueryTable(new(models.UserLevel))
	_, err := uls.PaginateAndScopeWhere(o, listRows, params).All(&userLevel)
	if err != nil {
		return nil, uls.Pagination
	}
	return userLevel, uls.Pagination
}

// GetExportData 获取导出数据
func (uls *UserLevelService) GetExportData(params url.Values) []*models.UserLevel {
	//搜索、查询字段赋值
	uls.SearchField = append(uls.SearchField, new(models.UserLevel).SearchField()...)
	var userLevel []*models.UserLevel
	o := orm.NewOrm().QueryTable(new(models.UserLevel))
	_, err := uls.ScopeWhere(o, params).All(&userLevel)
	if err != nil {
		return nil
	}
	return userLevel
}

// Create 新增用户等级
func (*UserLevelService) Create(form *formvalidate.UserLevelForm) int {
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
	}
	return 0
}

// Update 更新用户等级
func (*UserLevelService) Update(form *formvalidate.UserLevelForm) int {
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
		}
		return 0
	}
	return 0
}

// GetUserLevelById 根据id获取一条user_level数据
func (*UserLevelService) GetUserLevelById(id int) *models.UserLevel {
	o := orm.NewOrm()
	userLevel := models.UserLevel{Id: id}
	err := o.Read(&userLevel)
	if err != nil {
		return nil
	}
	return &userLevel
}

// GetUserLevel 获取所有用户等级
func (*UserLevelService) GetUserLevel() []*models.UserLevel {
	var userLevels []*models.UserLevel
	_, err := orm.NewOrm().QueryTable(new(models.UserLevel)).All(&userLevels)
	if err == nil {
		return userLevels
	}
	return nil
}

// Enable 启用
func (*UserLevelService) Enable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.UserLevel)).Filter("id__in", ids).Update(orm.Params{
		"status": 1,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Disable 禁用
func (*UserLevelService) Disable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.UserLevel)).Filter("id__in", ids).Update(orm.Params{
		"status": 0,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Del 删除
func (*UserLevelService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.UserLevel)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}
