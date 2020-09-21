package response

import (
	"beego-admin/global"
	"github.com/astaxie/beego/context"
)

const (
	ERROR   = 0
	SUCCESS = 1
)

//响应参数结构体
type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
	Url string `json:"url"`
	Wait int `json:"wait"`
}

//返回结果辅助函数
func Result(code int,msg string,data interface{},url string,wait int,header map[string]string,ctx *context.Context)  {
	if ctx.Request.Method == "POST"{
		result := Response{
			Code:code,
			Msg:msg,
			Data:data,
			Url:url,
			Wait:wait,
		}

		if len(header) > 0{
			for k,v := range header{
				ctx.Output.Header(k,v)
			}
		}

		ctx.Output.JSON(result,false,false)
	}

	if url == ""{
		url = ctx.Request.Referer()
		if  url == ""{
			url = "/admin/index/index"
		}
	}

	//闪存session
	if code == 0{
		ctx.Output.Session("success_message",msg)
	}else{
		ctx.Output.Session("error_message",msg)
	}

	ctx.Output.Session("url",url)
	ctx.Redirect(302,url)
}

//成功、普通返回
func Success(ctx *context.Context)  {
	Result(SUCCESS,"操作成功","",global.URL_BACK,0, map[string]string{},ctx)
}

//成功、返回自定义信息
func SuccessWithMessage(msg string,ctx *context.Context)  {
	Result(SUCCESS,msg,"",global.URL_BACK,0,map[string]string{},ctx)
}

//成功、返回自定义信息和url
func SuccessWithMessageAndUrl(msg string,url string,ctx *context.Context)  {
	Result(SUCCESS,msg,"",url,0,map[string]string{},ctx)
}

//成功、返回所有自定义信息
func SuccessWithDetailed(msg string,url string,data interface{},wait int,header map[string]string,ctx *context.Context)  {
	Result(SUCCESS,msg,data,url,wait,header,ctx)
}

//失败、普通返回
func Error(ctx *context.Context)  {
	Result(ERROR,"操作失败","",global.URL_CURRENT,0, map[string]string{},ctx)
}

//失败、返回自定义信息
func ErrorWithMessage(msg string,ctx *context.Context)  {
	Result(ERROR,msg,"",global.URL_BACK,0,map[string]string{},ctx)
}

//失败、返回自定义信息和url
func ErrorWithMessageAndUrl(msg string,url string,ctx *context.Context)  {
	Result(ERROR,msg,"",url,0,map[string]string{},ctx)
}

//失败、返回所有自定义信息
func ErrorWithDetailed(msg string,url string,data interface{},wait int,header map[string]string,ctx *context.Context)  {
	Result(ERROR,msg,data,url,wait,header,ctx)
}