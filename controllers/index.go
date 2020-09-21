package controllers

type IndexController struct {
	baseController
}

func (this *IndexController) Index()  {
	this.Data["json"] = "json"

	this.ServeJSON()
}
