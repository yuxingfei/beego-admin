package services

import (
	"beego-admin/formvalidate"
	"beego-admin/models"
	"github.com/astaxie/beego/orm"
)

// AdminMenuService struct
type AdminMenuService struct {
}

// GetAdminMenuByUrl 根据url获取admin_menu数据
func (*AdminMenuService) GetAdminMenuByUrl(url string) *models.AdminMenu {
	var adminMenu models.AdminMenu
	err := orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("url", url).One(&adminMenu)
	if err == nil {
		return &adminMenu
	}
	return nil
}

// GetCount 获取admin_menu 总数
func (*AdminMenuService) GetCount() int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminMenu)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// AllMenu 获取所有菜单
func (*AdminMenuService) AllMenu() []*models.AdminMenu {
	var adminMenus []*models.AdminMenu
	_, err := orm.NewOrm().QueryTable(new(models.AdminMenu)).OrderBy("sort_id", "id").All(&adminMenus)
	if err == nil {
		return adminMenus
	}
	return nil
}

// Menu 除去当前id之外的所有菜单id
func (*AdminMenuService) Menu(currentID int) []orm.Params {
	var adminMenusMap []orm.Params
	orm.NewOrm().QueryTable(new(models.AdminMenu)).Exclude("id", currentID).OrderBy("sort_id", "id").Values(&adminMenusMap, "id", "parent_id", "name", "sort_id")
	return adminMenusMap
}

// Create 创建菜单
func (*AdminMenuService) Create(form *formvalidate.AdminMenuForm) (id int64, err error) {
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

// Update 更新菜单
func (*AdminMenuService) Update(form *formvalidate.AdminMenuForm) int {
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
	}
	return 0
}

// IsExistUrl Url验重
func (*AdminMenuService) IsExistUrl(url string, id int) bool {
	if id == 0 {
		return orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("url", url).Exist()
	}
	return orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("url", url).Exclude("id", id).Exist()
}

// IsChildMenu 判断是否有子菜单
func (*AdminMenuService) IsChildMenu(ids []int) bool {
	return orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("parent_id__in", ids).Exist()
}

// Del 删除菜单
func (*AdminMenuService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

// GetAdminMenuById 通过id获取菜单信息
func (*AdminMenuService) GetAdminMenuById(id int) *models.AdminMenu {
	var adminMenu models.AdminMenu
	err := orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("id", id).One(&adminMenu)
	if err == nil {
		return &adminMenu
	}
	return nil
}
