package controllers

import (
	"fmt"
	"github.com/astaxie/beego/context"
)

type IndexController struct {
	baseController
}

func (this *IndexController) Index()  {
	this.Layout = "public/base.html"

	this.Data["flash_session_success_message"] = this.FlashSession("success_message",this.Ctx)
	this.Data["flash_session_error_message"] = this.FlashSession("error_message",this.Ctx)
	this.Data["flash_session_url"] = this.FlashSession("url",this.Ctx)
	this.TplName = "index/index.html"
}

func (this *IndexController)FlashSession(key string,ctx *context.Context) string {
	val,ok := this.GetSession(key).(string)
	if !ok{
		return ""
	}else{
		this.DelSession(key)
		fmt.Println("flash_session",val)
		return val
	}
}
