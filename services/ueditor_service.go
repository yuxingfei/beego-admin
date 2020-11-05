package services

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/google/uuid"
	"gopkg.in/ini.v1"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type UeditorService struct {
}

//获取配置信息
func (*UeditorService) GetConfig() map[string]interface{} {
	cfg, err := ini.Load("conf/ueditor.conf")
	configKeys := cfg.Section("ueditor").KeyStrings()

	result := make(map[string]interface{})
	if err != nil {
		result["state"] = "请求地址出错"
		return result
	}

	for _, key := range configKeys {
		value := cfg.Section("ueditor").Key(key).String()
		arr := strings.Split(value, "|")
		lenArr := len(arr)
		if lenArr < 1 {
			result[key] = ""
		} else if lenArr > 1 {
			result[key] = arr
		} else {
			result[key] = arr[0]
		}
	}
	return result
}

//上传图片
func (this *UeditorService) UploadImage(ctx *context.Context) map[string]interface{} {
	fieldName := beego.AppConfig.String("ueditor::imageFieldName")
	if fieldName == "" {
		return map[string]interface{}{
			"state": "not found field ueditor::imageFieldName.",
		}
	}
	result := this.upFile(fieldName, ctx)
	return result
}

//上传文件
func (*UeditorService) upFile(fieldName string, ctx *context.Context) map[string]interface{} {
	result := make(map[string]interface{})
	file, h, err := ctx.Request.FormFile(fieldName)
	if err != nil {
		result["state"] = err.Error()
		return result
	}
	defer file.Close()

	//自定义文件验证
	err = validateForAttachment(h)
	if err != nil {
		result["state"] = err.Error()
		return result
	}

	//数据表写入
	saveName := uuid.New().String()
	//后缀带. (.png)
	fileExt := path.Ext(h.Filename)
	savePath := "static/uploads/ueditor/" + saveName + fileExt
	saveRealDir := filepath.ToSlash(beego.AppPath + "/static/uploads/ueditor/")

	_, err = os.Stat(saveRealDir)
	if err != nil {
		err = os.MkdirAll(saveRealDir, os.ModePerm)
	}

	saveUrl := "/static/uploads/ueditor/" + saveName + fileExt

	f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		result["state"] = err.Error()
		return result
	}
	defer f.Close()
	_, err = io.Copy(f, file)

	if err != nil {
		result["state"] = err.Error()
		return result
	}

	result = map[string]interface{}{
		"state":    "SUCCESS",
		"url":      saveUrl,
		"title":    saveName + fileExt,
		"original": h.Filename,
		"type":     strings.TrimLeft(fileExt, "."),
		"size":     h.Size,
	}

	return result
}

//列出图片
func (this *UeditorService) ListImage(get url.Values) map[string]interface{} {
	result := make(map[string]interface{})
	allowFiles := beego.AppConfig.String("ueditor::imageManagerAllowFiles")
	listSize := beego.AppConfig.String("ueditor::imageManagerListSize")
	if allowFiles == "" || listSize == "" || len(get) <= 0 {
		result["state"] = "config params error."
	}

	listSizeInt,_ := strconv.Atoi(listSize)
	result = this.fileList(allowFiles,listSizeInt,get)

	return result
}

//列出图片
func (this *UeditorService) fileList(allowFiles string, listSize int, get url.Values) map[string]interface{} {
	result := make(map[string]interface{})
	var sizeInt, startInt int

	dir := "/static/uploads/ueditor/"

	//获取参数
	size := get.Get("size")
	if size == "" {
		sizeInt = listSize
	} else {
		sizeInt, _ = strconv.Atoi(size)
	}

	start := get.Get("start")
	if start == "" {
		startInt = 0
	} else {
		startInt, _ = strconv.Atoi(start)
	}

	end := startInt + sizeInt

	//获取文件列表
	files := this.getFiles(dir, allowFiles)
	if files == nil || len(files) <= 0 {
		result = map[string]interface{}{
			"state": "no match file",
			"list":  map[string]interface{}{},
			"start": startInt,
			"total": 0,
		}

		return result
	}

	//获取指定范围的列表
	lenFiles := len(files)

	result = map[string]interface{}{
		"state": "SUCCESS",
		"start": startInt,
		"total": lenFiles,
	}

	if startInt > lenFiles || end < 0 || startInt > end{
		result["list"] = map[string]interface{}{}
	}

	endInt := 0

	if end > lenFiles {
		endInt = lenFiles
	}else {
		endInt = end
	}

	result["list"] = files[startInt:endInt]

	return result
}

func (this *UeditorService) getFiles(dir, allowFiles string) []map[string]string {
	path := filepath.ToSlash(beego.AppPath + dir)
	var filesArr []map[string]string

	if !strings.HasPrefix(path, "/") {
		path = path + "/"
	}

	_, err := os.Stat(path)
	if err != nil {
		return nil
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}

	for _, file := range files {
		if file.IsDir() {
			childFilesArr := this.getFiles(filepath.ToSlash(path+file.Name()+"/"), allowFiles)
			if len(childFilesArr) > 0 {
				filesArr = append(filesArr, childFilesArr...)
			}
		} else {
			filesArr = append(filesArr, map[string]string{
				"url":   dir + file.Name(),
				"mtime": strconv.Itoa(int(file.ModTime().Unix())),
			})
		}
	}

	return filesArr

}
