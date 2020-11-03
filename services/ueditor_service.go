package services

import (
	"github.com/astaxie/beego"
	"strings"
)

type UeditorService struct {

}

//获取配置信息
func (*UeditorService) GetConfig() map[string]interface{} {
	ueditorConfig,err := beego.AppConfig.GetSection("ueditor")
	result := make(map[string]interface{})
	if err != nil{
		result["state"] = "请求地址出错"
		return result
	}

	for key,value := range ueditorConfig{
		arr := strings.Split(value,"|")
		lenArr := len(arr)
		if lenArr <  1{
			result[key] = ""
		}else if lenArr > 1 {
			result[key] = arr
		}else {
			result[key] = arr[0]
		}
	}

	return result
}

//上传图片
func (*UeditorService) UploadImage()  {

}