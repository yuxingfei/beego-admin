package controllers

import (
	beego "github.com/beego/beego/v2/adapter"
)

// MainController struct
type MainController struct {
	beego.Controller
}

// Get 显示
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
