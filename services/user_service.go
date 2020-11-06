package services

import (
	"beego-admin/form_validate"
	"beego-admin/models"
	"beego-admin/utils"
	beego_pagination "beego-admin/utils/beego-pagination"
	"encoding/base64"
	"github.com/astaxie/beego/orm"
	"net/url"
	"time"
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

//新增用户
func (*UserService) Create(form *form_validate.UserForm) int {
	user := models.User{
		Username:    form.Username,
		Nickname:    form.Nickname,
		UserLevelId: form.UserLevelId,
		Mobile:      form.Mobile,
		Description: form.Description,
		Status:      int8(form.Status),
		CreateTime:  int(time.Now().Unix()),
		UpdateTime:  int(time.Now().Unix()),
	}
	if form.Avatar != "" {
		user.Avatar = form.Avatar
	}

	//密码加密
	newPasswordForHash, err := utils.PasswordHash(form.Password)
	if err != nil {
		return 0
	}
	user.Password = base64.StdEncoding.EncodeToString([]byte(newPasswordForHash))

	id, err := orm.NewOrm().Insert(&user)

	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

//根据id获取一条user数据
func (*UserService) GetUserById(id int) *models.User {
	o := orm.NewOrm()
	user := models.User{Id: id}
	err := o.Read(&user)
	if err != nil {
		return nil
	}
	return &user
}

//更新用户
func (*UserService) Update(form *form_validate.UserForm) int {
	o := orm.NewOrm()
	user := models.User{Id: form.Id}
	if o.Read(&user) == nil {

		//判断密码是否相等
		if user.Password != form.Password {
			newPasswordForHash, err := utils.PasswordHash(form.Password)
			if err == nil {
				user.Password = base64.StdEncoding.EncodeToString([]byte(newPasswordForHash))
			}
		}

		user.Username = form.Username
		user.Nickname = form.Nickname
		user.UserLevelId = form.UserLevelId
		user.Mobile = form.Mobile
		user.Description = form.Description
		user.Status = int8(form.Status)
		user.UpdateTime = int(time.Now().Unix())

		if form.Avatar != "" {
			user.Avatar = form.Avatar
		}
		num, err := o.Update(&user)

		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
}

//启用
func (*UserService) Enable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.User)).Filter("id__in", ids).Update(orm.Params{
		"status": 1,
	})
	if err == nil {
		return int(num)
	} else {
		return 0
	}
}

//禁用
func (*UserService) Disable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.User)).Filter("id__in", ids).Update(orm.Params{
		"status": 0,
	})
	if err == nil {
		return int(num)
	} else {
		return 0
	}
}

//删除
func (*UserService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.User)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}

//获取导出数据
func (this *UserService) GetExportData(params url.Values) []*models.User {
	//搜索、查询字段赋值
	this.SearchField = append(this.SearchField, new(models.User).SearchField()...)
	var user []*models.User
	o := orm.NewOrm().QueryTable(new(models.User))
	_, err := this.ScopeWhere(o, params).All(&user)
	if err != nil {
		return nil
	} else {
		return user
	}
}
