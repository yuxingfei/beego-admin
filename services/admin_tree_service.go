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
	Ret string
	Html string
	Array map[int]orm.Params
	Text map[string]interface{}
}

//初始化

//获取左侧菜单
func (adminTreeService *AdminTreeService) GetLeftMenu(requestUrl string,user models.AdminUser) string {
	menu := user.GetShowMenu()
	maxLevel := 0
	currentId := 1
	parentIds := []int{0}

	for _,v := range menu{
		if v["Url"].(string) == requestUrl{
			parentIds = adminTreeService.getMenuParent(menu,int(v["Id"].(int64)),[]int{})
			currentId = int(v["Id"].(int64))
		}
	}

	if len(parentIds) == 0{
		parentIds = []int{0}
	}

	for k,v := range menu{
		menu[k]["Url"] = "/" + v["Url"].(string)
		tempLevel := adminTreeService.getLevel(int(v["Id"].(int64)),menu,0)
		menu[k]["Level"] = tempLevel
		if maxLevel <= tempLevel{
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
	textOpen       := " menu-open"
	textBaseThree := "'>"

	textBaseFour := "<li"
	textHoverLi  := " class='active'"
	textBaseFive := `>
	<a href='$url'>
	<i class='fa $icon'></i>
	<span>$name</span>
	</a>
	</li>`

	text0       := textBaseOne + textBaseTwo + textBaseThree
	text1       := textBaseOne + textHover + textBaseTwo + textOpen + textBaseThree
	text2       := "</ul></li>"
	textCurrent := textBaseFour + textHoverLi + textBaseFive
	textOther   := textBaseFour + textBaseFive

	adminTreeService.Text = make(map[string]interface{})
	for i := 0; i <= maxLevel;i++{
		adminTreeService.Text[strconv.Itoa(i)] = []string{text0,text1,text2}
	}
	adminTreeService.Text["current"] = textCurrent
	adminTreeService.Text["other"] = textOther

	return adminTreeService.getAuthTree(0,currentId,parentIds)

}

//获取父级菜单
func (adminTreeService *AdminTreeService) getMenuParent(menu map[int]orm.Params,myId int,parentIds []int) []int {
	for _,a := range menu{
		if int(a["Id"].(int64)) == myId && int(a["ParentId"].(int64)) != 0{
			parentIds = append(parentIds,int(a["ParentId"].(int64)))
			parentIds = adminTreeService.getMenuParent(menu,int(a["ParentId"].(int64)),parentIds)
		}
	}
	if len(parentIds) > 0{
		return parentIds
	}else{
		return []int{}
	}
}

//递归获取级别
func (adminTreeService *AdminTreeService) getLevel(id int,menu map[int]orm.Params,i int) int {
	v,ok := menu[id]["ParentId"].(int64)
	if (!ok || int(v) == 0) || id == int(v){
		return i
	}
	i++
	return adminTreeService.getLevel(int(v),menu,i)
}

func (adminTreeService *AdminTreeService) initTree(menu map[int]orm.Params)  {
	adminTreeService.Array = make(map[int]orm.Params)
	adminTreeService.Array = menu
	adminTreeService.Ret = ""
	adminTreeService.Html = ""
}

//获取后台左侧菜单
func (adminTreeService *AdminTreeService) getAuthTree(myId int,currentId int,parentIds []int) string {
	nStr := ""
	child := adminTreeService.getChild(myId)
	if len(child) > 0{
		menu := make(map[string]interface{})
		//取key最小的一个，防止for range随机取，导致每次菜单顺序不同
		sortId := 99999
		for k,_ := range child{
			if k < sortId{
				sortId = k
			}
		}
		menu = child[sortId]

		//获取当前等级的html
		var textHtmlInterface interface{}
		if adminTreeService.Text[strconv.Itoa(menu["Level"].(int))] != ""{
			//[]string类型
			textHtmlInterface = adminTreeService.Text[strconv.Itoa(menu["Level"].(int))]
		}else{
			//string类型
			textHtmlInterface = adminTreeService.Text["other"]
		}

		//child排序，防止菜单位置每次都不同
		var childKeys []int
		for k := range child{
			childKeys = append(childKeys,k)
		}
		sort.Ints(childKeys)
		for _,key := range childKeys{
			k := key
			v := child[key]

			if len(adminTreeService.getChild(k)) > 0{
				textHtmlArr := textHtmlInterface.([]string)
				if utils.KeyInArrayForInt(parentIds,k){
					nStr = adminTreeService.strReplace(textHtmlArr[1],v)
					adminTreeService.Html += nStr
				}else {
					nStr = adminTreeService.strReplace(textHtmlArr[0],v)
					adminTreeService.Html += nStr
				}
				adminTreeService.getAuthTree(k,currentId,parentIds)
				nStr = adminTreeService.strReplace(textHtmlArr[2],v)
				adminTreeService.Html += nStr
			}else if k == currentId {
				a := adminTreeService.Text["current"].(string)
				nStr = adminTreeService.strReplace(a,v)
				adminTreeService.Html += nStr
			}else {
				a := adminTreeService.Text["other"].(string)
				nStr = adminTreeService.strReplace(a,v)
				adminTreeService.Html += nStr
			}
		}
	}
	return adminTreeService.Html
}

//得到子级数组
func (adminTreeService *AdminTreeService) getChild(pid int) map[int]map[string]interface{} {
	result := make(map[int]map[string]interface{})
	for k,v := range adminTreeService.Array{
		if int(v["ParentId"].(int64)) == pid{
			result[k] = v
		}
	}
	return result
}

//替换字符串
func (adminTreeService *AdminTreeService) strReplace(str string,m map[string]interface{}) string {
	str = strings.ReplaceAll(str,"$icon",m["Icon"].(string))
	str = strings.ReplaceAll(str,"$name",m["Name"].(string))
	str = strings.ReplaceAll(str,"$url",m["Url"].(string))
	return str
}
