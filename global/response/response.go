package response

import (
	"beego-admin/global"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
)

const (
	ERROR   = 0
	SUCCESS = 1
)

// Response 响应参数结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Url  string      `json:"url"`
	Wait int         `json:"wait"`
}

// Result 返回结果辅助函数
func Result(code int, msg string, data interface{}, url string, wait int, header map[string]string, ctx *context.Context) {
	if ctx.Input.IsPost() {
		result := Response{
			Code: code,
			Msg:  msg,
			Data: data,
			Url:  url,
			Wait: wait,
		}

		if len(header) > 0 {
			for k, v := range header {
				ctx.Output.Header(k, v)
			}
		}

		ctx.Output.JSON(result, false, false)

		//Controller中this.StopRun()用法
		panic(beego.ErrAbort)
	}

	if url == "" {
		url = ctx.Request.Referer()
		if url == "" {
			url = "/admin/index/index"
		}
	}

	ctx.Redirect(http.StatusFound, url)
}

// Success 成功、普通返回
func Success(ctx *context.Context) {
	Result(SUCCESS, "操作成功", "", global.URL_BACK, 0, map[string]string{}, ctx)
}

// SuccessWithMessage 成功、返回自定义信息
func SuccessWithMessage(msg string, ctx *context.Context) {
	Result(SUCCESS, msg, "", global.URL_BACK, 0, map[string]string{}, ctx)
}

// SuccessWithMessageAndUrl 成功、返回自定义信息和url
func SuccessWithMessageAndUrl(msg string, url string, ctx *context.Context) {
	Result(SUCCESS, msg, "", url, 0, map[string]string{}, ctx)
}

// SuccessWithDetailed 成功、返回所有自定义信息
func SuccessWithDetailed(msg string, url string, data interface{}, wait int, header map[string]string, ctx *context.Context) {
	Result(SUCCESS, msg, data, url, wait, header, ctx)
}

// Error 失败、普通返回
func Error(ctx *context.Context) {
	Result(ERROR, "操作失败", "", global.URL_CURRENT, 0, map[string]string{}, ctx)
}

// ErrorWithMessage 失败、返回自定义信息
func ErrorWithMessage(msg string, ctx *context.Context) {
	Result(ERROR, msg, "", global.URL_CURRENT, 0, map[string]string{}, ctx)
}

// ErrorWithMessageAndUrl 失败、返回自定义信息和url
func ErrorWithMessageAndUrl(msg string, url string, ctx *context.Context) {
	Result(ERROR, msg, "", url, 0, map[string]string{}, ctx)
}

// ErrorWithDetailed 失败、返回所有自定义信息
func ErrorWithDetailed(msg string, url string, data interface{}, wait int, header map[string]string, ctx *context.Context) {
	Result(ERROR, msg, data, url, wait, header, ctx)
}
