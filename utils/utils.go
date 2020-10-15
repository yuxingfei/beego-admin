package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/dchest/captcha"
	"golang.org/x/crypto/bcrypt"
	"io"
	"strconv"
)

type CaptchaResponse struct {
	CaptchaId  string
	CaptchaUrl string
}

//获取验证码
func GetCaptcha() *CaptchaResponse {
	captchaId := captcha.NewLen(4)
	return &CaptchaResponse{
		CaptchaId:  captchaId,
		CaptchaUrl: fmt.Sprintf("/admin/auth/captcha/%s.png", captchaId),
	}
}

//模仿php的array_key_exists,判断是否存在map中
func KeyInMap(key string, m map[string]interface{}) bool {
	_, ok := m[key]
	if ok {
		return true
	} else {
		return false
	}
}

//模仿php的in_array,判断是否存在string数组中
func InArrayForString(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

//模仿php的in_array,判断是否存在int数组中
func InArrayForInt(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

//php的函数password_hash
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

//php的函数password_verify
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//int数组转string数组
func IntArrToStringArr(arr []int) []string {
	var stringArr []string
	for _, v := range arr {
		stringArr = append(stringArr, strconv.Itoa(v))
	}
	return stringArr
}

//对字符串进行MD5哈希
func GetMd5String(str string) string {
	t := md5.New()
	io.WriteString(t, str)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//对字符串进行SHA1哈希
func GetSha1String(str string) string {
	t := sha1.New()
	io.WriteString(t, str)
	return fmt.Sprintf("%x", t.Sum(nil))
}
