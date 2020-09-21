package utils

import (
	"github.com/dchest/captcha"
)

//获取验证码id
func CaptchaId() string {
	return captcha.NewLen(4)
}

//模仿php的in_array,判断是否存在map中
func KeyInMap(key string,m map[string]interface{}) bool {
	_,ok := m[key]
	if ok{
		return true
	}else{
		return false
	}
}
