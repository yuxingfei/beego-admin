//Package template 自定义模板函数
package template

import (
	beego "github.com/beego/beego/v2/adapter"
	"math"
	"strconv"
	"time"
)

func init() {
	beego.AddFuncMap("UnixTimeForFormat", UnixTimeForFormat)
	beego.AddFuncMap("FormatSize", FormatSize)
}

// UnixTimeForFormat 时间轴转时间字符串
func UnixTimeForFormat(timeUnix int) string {
	//转化所需模板
	timeLayout := "2006-01-02 15:04:05"
	return time.Unix(int64(timeUnix), 0).Format(timeLayout)
}

// FormatSize 格式化文件大小单位
func FormatSize(size, delimiter string) string {
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		return ""
	}
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	var i int
	for i = 0; sizeInt >= 1024 && i < 5; i++ {
		sizeInt /= 1024
	}
	return strconv.FormatFloat(math.Round(float64(sizeInt)), 'f', -1, 64) + delimiter + units[i]
}
