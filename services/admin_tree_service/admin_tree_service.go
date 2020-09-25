package admin_tree_service

import (
	"beego-admin/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

//初始化
var (
	ret string
	html string
    array map[int]orm.Params
	text map[string]interface{}
)

//获取左侧菜单
func GetLeftMenu(requestUrl string,user models.AdminUser) map[string]interface{} {
	menu := user.GetShowMenu()
	fmt.Println("menu = ",menu)
	maxLevel := 0
	currentId := 1
	parentIds := []int{0}

	for _,v := range menu{
		if v["Url"].(string) == requestUrl{
			parentIds = getMenuParent(menu,int(v["Id"].(int64)),[]int{})
			currentId = int(v["Id"].(int64))
		}
	}

	if len(parentIds) == 0{
		parentIds = []int{0}
	}

	for k,v := range menu{
		menu[k]["Url"] = "/" + v["Url"].(string)
		tempLevel := getLevel(int(v["Id"].(int64)),menu,0)
		menu[k]["Level"] = tempLevel
		if maxLevel <= tempLevel{
			maxLevel = tempLevel
		}
	}

	initTree(menu)

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

	//text
	text = make(map[string]interface{})

	for i := 0; i < maxLevel;i++{
		text[strconv.Itoa(i)] = []string{text0,text1,text2}
	}
	text["current"] = textCurrent
	text["other"] = textOther

	fmt.Println(`$this->get_authTree(0, $current_id, $parent_ids);`)








	fmt.Println("newmenu = ",menu)

	fmt.Println("max_level=",maxLevel)
	fmt.Println("current_id=",currentId)
	fmt.Println("parent_ids=",parentIds)



	return map[string]interface{}{}
}

//获取父级菜单
func getMenuParent(menu map[int]orm.Params,myId int,parentIds []int) []int {
	for _,a := range menu{
		if int(a["Id"].(int64)) == myId && int(a["ParentId"].(int64)) != 0{
			parentIds = append(parentIds,int(a["ParentId"].(int64)))
			parentIds = getMenuParent(menu,int(a["ParentId"].(int64)),parentIds)
		}
	}
	if len(parentIds) > 0{
		return parentIds
	}else{
		return []int{}
	}
}

//递归获取级别
func getLevel(id int,menu map[int]orm.Params,i int) int {
	v,ok := menu[id]["ParentId"].(int64)
	if (!ok || int(v) == 0) || id == int(v){
		return i
	}
	i++
	return getLevel(int(v),menu,i)
}

func initTree(menu map[int]orm.Params)  {
	array = make(map[int]orm.Params)
	array = menu
	ret = ""
	html = ""
}

//获取后台左侧菜单
func getAuthTree(myId int,currentId int,parentIds []int)  {
	nStr := ""
	child := getChild(myId)
	if len(child) > 0{
		
	}

}

//得到子级数组
func getChild(pid int) map[int]map[string]interface{} {
	result := make(map[int]map[string]interface{})
	for k,v := range array{
		if int(v["ParentId"].(int64)) == pid{
			result[k] = v
		}
	}
	return result
}
