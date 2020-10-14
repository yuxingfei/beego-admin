package services

import (
	"beego-admin/form_validate"
	"beego-admin/models"
	"github.com/astaxie/beego/orm"
)

type AdminMenuService struct {
}

//根据url获取admin_menu数据
func (*AdminMenuService) GetAdminMenuByUrl(url string) *models.AdminMenu {
	var adminMenu models.AdminMenu
	err := orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("url", url).One(&adminMenu)
	if err == nil {
		return &adminMenu
	} else {
		return nil
	}
}

//获取admin_menu 总数
func (*AdminMenuService) GetCount() int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminMenu)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

//获取所有菜单
func (*AdminMenuService) AllMenu() []*models.AdminMenu {
	var adminMenus []*models.AdminMenu
	_, err := orm.NewOrm().QueryTable(new(models.AdminMenu)).OrderBy("sort_id", "id").All(&adminMenus)
	if err == nil {
		return adminMenus
	} else {
		return nil
	}
}

//除去当前id之外的所有菜单id
func (*AdminMenuService) Menu(currentId int) []orm.Params {
	var adminMenusMap []orm.Params
	orm.NewOrm().QueryTable(new(models.AdminMenu)).Exclude("id", currentId).OrderBy("sort_id", "id").Values(&adminMenusMap, "id", "parent_id", "name", "sort_id")
	return adminMenusMap
}

//创建菜单
func (*AdminMenuService) Create(form *form_validate.AdminMenuForm) (id int64, err error) {
	adminMenu := models.AdminMenu{
		ParentId:  form.ParentId,
		Name:      form.Name,
		Url:       form.Url,
		Icon:      form.Icon,
		IsShow:    form.IsShow,
		SortId:    form.SortId,
		LogMethod: form.LogMethod,
	}

	return orm.NewOrm().Insert(&adminMenu)
}

//更新菜单
func (*AdminMenuService) Update(form *form_validate.AdminMenuForm) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("id", form.Id).Update(orm.Params{
		"parent_id":  form.ParentId,
		"name":       form.Name,
		"url":        form.Url,
		"icon":       form.Icon,
		"is_show":    form.IsShow,
		"sort_id":    form.SortId,
		"log_method": form.LogMethod,
	})
	if err == nil {
		return int(num)
	} else {
		return 0
	}
}

//Url验重
func (*AdminMenuService) IsExistUrl(url string, id int) bool {
	if id == 0 {
		return orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("url", url).Exist()
	} else {
		return orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("url", url).Exclude("id", id).Exist()
	}
}

//判断是否有子菜单
func (*AdminMenuService) IsChildMenu(ids []int) bool {
	return orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("parent_id__in", ids).Exist()
}

//删除菜单
func (*AdminMenuService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}

//通过id获取菜单信息
func (*AdminMenuService) GetAdminMenuById(id int) *models.AdminMenu {
	var adminMenu models.AdminMenu
	err := orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("id", id).One(&adminMenu)
	if err == nil {
		return &adminMenu
	} else {
		return nil
	}
}
