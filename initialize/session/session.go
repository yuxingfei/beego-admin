package session

import beego "github.com/beego/beego/v2/adapter"

func init() {
	//session 过期时间，默认值是 3600 秒
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 7200

	//session 默认存在客户端的 cookie 的时间，默认值是 3600 秒
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 7200
}
