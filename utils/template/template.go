//自定义模板函数
package template

import (
	"github.com/astaxie/beego"
	"time"
)

func init() {
	beego.AddFuncMap("UnixTimeForFormat", UnixTimeForFormat)
}

//时间轴转时间字符串
func UnixTimeForFormat(timeUnix int) string {
	//转化所需模板
	timeLayout := "2006-01-02 15:04:05"
	return time.Unix(int64(timeUnix), 0).Format(timeLayout)
}
