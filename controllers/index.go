package controllers

type IndexController struct {
	baseController
}

func (this *IndexController) Index()  {

	m := map[string]interface{}{
		"msg":"ok",
	}
	this.Data["json"] = m
	this.ServeJSON()
}
