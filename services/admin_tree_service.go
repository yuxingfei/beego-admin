package services

import (
	"beego-admin/models"
	"beego-admin/utils"
	"github.com/astaxie/beego/orm"
	"sort"
	"strconv"
	"strings"
)

type AdminTreeService struct {
	Ret   string
	Html  string
	Array map[int]orm.Params
	Text  map[string]interface{}
}

var (
	icon  []string = []string{"│", "├", "└"}
	space string   = "&nbsp;&nbsp;"
)

//获取左侧菜单
func (adminTreeService *AdminTreeService) GetLeftMenu(requestUrl string, user models.AdminUser) string {
	menu := user.GetShowMenu()
	maxLevel := 0
	currentId := 1
	parentIds := []int{0}

	for _, v := range menu {
		if v["Url"].(string) == requestUrl {
			parentIds = adminTreeService.getMenuParent(menu, int(v["Id"].(int64)), []int{})
			currentId = int(v["Id"].(int64))
		}
	}

	if len(parentIds) == 0 {
		parentIds = []int{0}
	}

	for k, v := range menu {
		menu[k]["Url"] = "/" + v["Url"].(string)
		tempLevel := adminTreeService.GetLevel(int(v["Id"].(int64)), menu, 0)
		menu[k]["Level"] = tempLevel
		if maxLevel <= tempLevel {
			maxLevel = tempLevel
		}
	}

	adminTreeService.initTree(menu)

	textBaseOne := "<li class='treeview"
	textHover := " active"
	textBaseTwo := `'><a href='javascript:void(0);'>
	<i class='fa $icon'></i>
	<span>
	$name
	</span>
	<span class='pull-right-container'><i class='fa fa-angle-left pull-right'></i></span>
	</a><ul class='treeview-menu`
	textOpen := " menu-open"
	textBaseThree := "'>"

	textBaseFour := "<li"
	textHoverLi := " class='active'"
	textBaseFive := `>
	<a href='$url'>
	<i class='fa $icon'></i>
	<span>$name</span>
	</a>
	</li>`

	text0 := textBaseOne + textBaseTwo + textBaseThree
	text1 := textBaseOne + textHover + textBaseTwo + textOpen + textBaseThree
	text2 := "</ul></li>"
	textCurrent := textBaseFour + textHoverLi + textBaseFive
	textOther := textBaseFour + textBaseFive

	adminTreeService.Text = make(map[string]interface{})
	for i := 0; i <= maxLevel; i++ {
		adminTreeService.Text[strconv.Itoa(i)] = []string{text0, text1, text2}
	}
	adminTreeService.Text["current"] = textCurrent
	adminTreeService.Text["other"] = textOther

	return adminTreeService.getAuthTree(0, currentId, parentIds)

}

//获取父级菜单
func (adminTreeService *AdminTreeService) getMenuParent(menu map[int]orm.Params, myId int, parentIds []int) []int {
	for _, a := range menu {
		if int(a["Id"].(int64)) == myId && int(a["ParentId"].(int64)) != 0 {
			parentIds = append(parentIds, int(a["ParentId"].(int64)))
			parentIds = adminTreeService.getMenuParent(menu, int(a["ParentId"].(int64)), parentIds)
		}
	}
	if len(parentIds) > 0 {
		return parentIds
	} else {
		return []int{}
	}
}

//递归获取级别
func (adminTreeService *AdminTreeService) GetLevel(id int, menu map[int]orm.Params, i int) int {
	v, ok := menu[id]["ParentId"].(int64)
	if (!ok || int(v) == 0) || id == int(v) {
		return i
	}
	i++
	return adminTreeService.GetLevel(int(v), menu, i)
}

func (adminTreeService *AdminTreeService) initTree(menu map[int]orm.Params) {
	adminTreeService.Array = make(map[int]orm.Params)
	adminTreeService.Array = menu
	adminTreeService.Ret = ""
	adminTreeService.Html = ""
}

//获取后台左侧菜单
func (adminTreeService *AdminTreeService) getAuthTree(myId int, currentId int, parentIds []int) string {
	nStr := ""
	child := adminTreeService.getChild(myId)
	if len(child) > 0 {
		menu := make(map[string]interface{})
		//取key最小的一个，防止for range随机取，导致每次菜单顺序不同
		sortId := 99999
		for k, _ := range child {
			if k < sortId {
				sortId = k
			}
		}
		menu = child[sortId]

		//获取当前等级的html
		var textHtmlInterface interface{}
		if adminTreeService.Text[strconv.Itoa(menu["Level"].(int))] != "" {
			//[]string类型
			textHtmlInterface = adminTreeService.Text[strconv.Itoa(menu["Level"].(int))]
		} else {
			//string类型
			textHtmlInterface = adminTreeService.Text["other"]
		}

		//child排序，防止菜单位置每次都不同
		var childKeys []int
		for k := range child {
			childKeys = append(childKeys, k)
		}
		sort.Ints(childKeys)
		for _, key := range childKeys {
			k := key
			v := child[key]

			if len(adminTreeService.getChild(k)) > 0 {
				textHtmlArr := textHtmlInterface.([]string)
				if utils.KeyInArrayForInt(parentIds, k) {
					nStr = adminTreeService.strReplace(textHtmlArr[1], v)
					adminTreeService.Html += nStr
				} else {
					nStr = adminTreeService.strReplace(textHtmlArr[0], v)
					adminTreeService.Html += nStr
				}
				adminTreeService.getAuthTree(k, currentId, parentIds)
				nStr = adminTreeService.strReplace(textHtmlArr[2], v)
				adminTreeService.Html += nStr
			} else if k == currentId {
				a := adminTreeService.Text["current"].(string)
				nStr = adminTreeService.strReplace(a, v)
				adminTreeService.Html += nStr
			} else {
				a := adminTreeService.Text["other"].(string)
				nStr = adminTreeService.strReplace(a, v)
				adminTreeService.Html += nStr
			}
		}
	}
	return adminTreeService.Html
}

//得到子级数组
func (adminTreeService *AdminTreeService) getChild(pid int) map[int]map[string]interface{} {
	result := make(map[int]map[string]interface{})
	for k, v := range adminTreeService.Array {
		parentId, ok := v["ParentId"].(int64)
		var parentIdInt int
		if ok {
			parentIdInt = int(parentId)
		} else {
			parentIdInt = v["ParentId"].(int)
		}

		if parentIdInt == pid {
			result[k] = v
		}
	}
	return result
}

//替换字符串
func (adminTreeService *AdminTreeService) strReplace(str string, m map[string]interface{}) string {
	str = strings.ReplaceAll(str, "$icon", m["Icon"].(string))
	str = strings.ReplaceAll(str, "$name", m["Name"].(string))
	str = strings.ReplaceAll(str, "$url", m["Url"].(string))
	return str
}

//得到树型结构
//myId 表示获得这个ID下的所有子级
//str 生成树型结构的基本代码，例如："<option value=\$id \$selected>\$spacer\$name</option>"
//sid 被选中的ID，比如在做树型下拉框的时候需要用到
func (adminTreeService *AdminTreeService) GetTree(myId int, str string, sid int, adds string, strGroup string) string {
	number := 1
	child := adminTreeService.getChild(myId)

	if len(child) > 0 {
		total := len(child)

		//child排序
		var ids []int
		for id, _ := range child {
			ids = append(ids, id)
		}
		sort.Ints(ids)

		for _, id := range ids {
			value := child[id]

			j := ""
			k := ""
			if number == total {
				j += icon[2]
			} else {
				j += icon[1]
				if adds != "" {
					k = icon[0]
				} else {
					k = ""
				}
			}

			spacer := ""
			if adds != "" {
				spacer = adds + j
			}
			selected := ""
			if id == sid {
				selected = "selected"
			}
			nStr := ""
			parentIdInt, ok := value["ParentId"].(int)
			if !ok {
				parentIdInt = int(value["ParentId"].(int64))
			}
			if 0 == parentIdInt && strGroup != "" {
				nStr = strGroup
			} else {
				nStr = str
			}

			//id orm转换可能是int或者int64类型，兼容
			idInt, ok := value["Id"].(int)
			if !ok {
				idInt64, ok := value["Id"].(int64)
				if ok {
					idInt = int(idInt64)
					nStr = strings.ReplaceAll(nStr, "$id", strconv.Itoa(idInt))
				}
			} else {
				nStr = strings.ReplaceAll(nStr, "$id", strconv.Itoa(idInt))
			}

			levelInt, ok := value["Level"].(int)
			if !ok {
				levelInt64, ok := value["Level"].(int64)
				if ok {
					levelInt = int(levelInt64)
					nStr = strings.ReplaceAll(nStr, "$level", strconv.Itoa(levelInt))
				}
			} else {
				nStr = strings.ReplaceAll(nStr, "$level", strconv.Itoa(levelInt))
			}

			sortIdInt, ok := value["SortId"].(int)
			if !ok {
				sortIdInt64, ok := value["SortId"].(int64)
				if ok {
					sortIdInt = int(sortIdInt64)
					nStr = strings.ReplaceAll(nStr, "$sort_id", strconv.Itoa(sortIdInt))
				}
			} else {
				nStr = strings.ReplaceAll(nStr, "$sort_id", strconv.Itoa(sortIdInt))
			}

			parentIdNodeStringValue, ok := value["ParentIdNode"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$parent_id_node", parentIdNodeStringValue)
			}

			nStr = strings.ReplaceAll(nStr, "$spacer", spacer)
			nStr = strings.ReplaceAll(nStr, "$selected", selected)

			nameValue, ok := value["Name"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$name", nameValue)
			}

			urlValue, ok := value["Url"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$url", urlValue)
			}

			nStr = strings.ReplaceAll(nStr, "$parent_id", strconv.Itoa(parentIdInt))

			iconValue, ok := value["Icon"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$icon", iconValue)
			}

			isShowValue, ok := value["IsShow"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$is_show", isShowValue)
			}

			logMethodValue, ok := value["LogMethod"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$log_method", logMethodValue)
			}

			strManageValue, ok := value["StrManage"].(string)
			if ok {
				strManageValue = strings.ReplaceAll(strManageValue, "\\", "")
				nStr = strings.ReplaceAll(nStr, "$str_manage", strManageValue)
			}

			adminTreeService.Ret += nStr
			adminTreeService.GetTree(id, str, sid, adds+k+space, strGroup)

			number++
		}

	}

	return adminTreeService.Ret
}

//菜单选择 select树形选择
func (adminTreeService *AdminTreeService) Menu(selected int, currentId int) string {
	var adminMenuService AdminMenuService
	result := adminMenuService.Menu(currentId)
	resultKey := make(map[int]orm.Params)
	if result != nil {
		for _, r := range result {
			idInt, ok := r["Id"].(int)
			if !ok {
				idInt = int(r["Id"].(int64))
			}
			resultKey[idInt] = r
			if idInt == selected {
				resultKey[idInt]["selected"] = "selected"
			} else {
				resultKey[idInt]["selected"] = ""
			}
		}

		str := `<option value='$id' $selected >$spacer $name</option>`
		adminTreeService.initTree(resultKey)
		return adminTreeService.GetTree(0, str, selected, "", "")
	} else {
		return ""
	}
}
