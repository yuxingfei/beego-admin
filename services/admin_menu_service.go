package services

import (
	"beego-admin/form_validate"
	"beego-admin/models"
	"github.com/astaxie/beego/orm"
	"strconv"
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

//生成菜单树
func (*AdminMenuService) AdminMenuTree() string {
	//查询所有菜单
	var adminMenus []*models.AdminMenu
	_, err := orm.NewOrm().QueryTable(new(models.AdminMenu)).OrderBy("sort_id", "id").All(&adminMenus)
	if err == nil {
		result := make(map[int]orm.Params)
		var adminTreeService AdminTreeService
		for _, adminMenu := range adminMenus {
			n := adminMenu.Id
			//初始化orm.Params map类型
			if result[n] == nil {
				result[n] = make(orm.Params)
			}

			result[n]["Id"] = adminMenu.Id
			result[n]["ParentId"] = adminMenu.ParentId
			result[n]["Name"] = adminMenu.Name
			result[n]["Url"] = adminMenu.Url
			result[n]["Icon"] = adminMenu.Icon
			result[n]["IsShow"] = adminMenu.IsShow
			result[n]["SortId"] = adminMenu.SortId

			result[n]["Level"] = adminTreeService.GetLevel(adminMenu.Id, result, 0)
			if adminMenu.ParentId > 0 {
				result[n]["ParentIdNode"] = ` class="child-of-node-` + strconv.Itoa(adminMenu.ParentId) + `"`
			} else {
				result[n]["ParentIdNode"] = ""
			}
			result[n]["StrManage"] = `<a href="/admin/admin_menu/edit?id=` + strconv.Itoa(adminMenu.Id) + `" class="btn btn-primary btn-xs" title="修改" data-toggle="tooltip"><i class="fa fa-pencil"></i></a> <a class="btn btn-danger btn-xs AjaxButton" data-id="` + strconv.Itoa(adminMenu.Id) + `" data-url="del"  data-confirm-title="删除确认" data-confirm-content=\'您确定要删除ID为 <span class="text-red"> ` + strconv.Itoa(adminMenu.Id) + ` </span> 的数据吗\'  data-toggle="tooltip" title="删除"><i class="fa fa-trash"></i></a>`
			if adminMenu.IsShow == 1 {
				result[n]["IsShow"] = "显示"
			} else {
				result[n]["IsShow"] = "隐藏"
			}
			result[n]["LogMethod"] = adminMenu.LogMethod
		}
		str := `<tr id='node-$id' data-level='$level' $parent_id_node><td><input type='checkbox' onclick='checkThis(this)'
                     name='data-checkbox' data-id='$id' class='checkbox data-list-check' value='$id' placeholder='选择/取消'>
                    </td><td>$id</td><td>$spacer$name</td><td>$url</td>
                    <td>$parent_id</td><td><i class='fa $icon'></i><span>($icon)</span></td>
                    <td>$sort_id</td><td>$is_show</td><td>$log_method</td><td class='td-do'>$str_manage</td></tr>`

		adminTreeService.initTree(result)

		return adminTreeService.GetTree(0, str, 0, "", "")

	} else {
		return ""
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
