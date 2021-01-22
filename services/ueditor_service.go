package services

import (
	"beego-admin/global"
	"encoding/base64"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/google/uuid"
	"gopkg.in/ini.v1"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// UeditorService struct
type UeditorService struct {
}

// GetConfig 获取配置信息
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

// UploadImage 上传图片
func (us *UeditorService) UploadImage(ctx *context.Context) map[string]interface{} {
	fieldName := global.BA_CONFIG.Ueditor.ImageFieldName
	if fieldName == "" {
		return map[string]interface{}{
			"state": "not found field ueditor::imageFieldName.",
		}
	}
	return us.upFile(fieldName, ctx)
}

// upFile 上传文件
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

	//存储名
	saveName := uuid.New().String()
	//后缀带. (.png)
	fileExt := path.Ext(h.Filename)
	savePath := "static/uploads/ueditor/" + saveName + fileExt
	saveRealDir := filepath.ToSlash(beego.AppPath + "/static/uploads/ueditor/")

	_, err = os.Stat(saveRealDir)
	if err != nil {
		err = os.MkdirAll(saveRealDir, os.ModePerm)
	}

	saveURL := "/static/uploads/ueditor/" + saveName + fileExt

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
		"url":      saveURL,
		"title":    saveName + fileExt,
		"original": h.Filename,
		"type":     strings.TrimLeft(fileExt, "."),
		"size":     h.Size,
	}

	return result
}

// ListImage 列出图片
func (us *UeditorService) ListImage(get url.Values) map[string]interface{} {
	result := make(map[string]interface{})
	allowFiles := global.BA_CONFIG.Ueditor.ImageManagerAllowFiles
	//ext前面的.号去掉
	allowFiles = strings.ReplaceAll(allowFiles, ".", "")
	listSize := global.BA_CONFIG.Ueditor.ImageManagerListSize
	if allowFiles == "" || listSize == "" || len(get) <= 0 {
		result["state"] = "config params error."
	}

	listSizeInt, _ := strconv.Atoi(listSize)
	result = us.fileList(allowFiles, listSizeInt, get)

	return result
}

// fileList 列出图片
func (us *UeditorService) fileList(allowFiles string, listSize int, get url.Values) map[string]interface{} {
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
	files := us.getFiles(dir, allowFiles)
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

	if startInt > lenFiles || end < 0 || startInt > end {
		result["list"] = map[string]interface{}{}
	}

	endInt := 0

	if end > lenFiles {
		endInt = lenFiles
	} else {
		endInt = end
	}

	result["list"] = files[startInt:endInt]

	return result
}

// getFiles 遍历获取目录下的指定类型的文件
func (us *UeditorService) getFiles(dir, allowFiles string) []map[string]string {
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
			childFilesArr := us.getFiles(filepath.ToSlash(path+file.Name()+"/"), allowFiles)
			if len(childFilesArr) > 0 {
				filesArr = append(filesArr, childFilesArr...)
			}
		} else {
			if !strings.Contains(allowFiles, strings.ToLower(strings.TrimLeft(filepath.Ext(file.Name()), "."))) {
				continue
			}
			filesArr = append(filesArr, map[string]string{
				"url":   dir + file.Name(),
				"mtime": strconv.Itoa(int(file.ModTime().Unix())),
			})
		}
	}

	return filesArr
}

// UploadVideo 上传视频
func (us *UeditorService) UploadVideo(ctx *context.Context) map[string]interface{} {
	fieldName := global.BA_CONFIG.Ueditor.VideoFieldName
	if fieldName == "" {
		return map[string]interface{}{
			"state": "not found field ueditor::videoFieldName.",
		}
	}

	return us.upFile(fieldName, ctx)
}

// UploadFile 上传文件
func (us *UeditorService) UploadFile(ctx *context.Context) map[string]interface{} {
	fieldName := global.BA_CONFIG.Ueditor.FileFieldName
	if fieldName == "" {
		return map[string]interface{}{
			"state": "not found field ueditor::fileFieldName.",
		}
	}

	return us.upFile(fieldName, ctx)
}

// ListFile 列出文件
func (us *UeditorService) ListFile(get url.Values) map[string]interface{} {
	result := make(map[string]interface{})
	allowFiles := global.BA_CONFIG.Ueditor.FileManagerAllowFiles
	//ext前面的.号去掉
	allowFiles = strings.ReplaceAll(allowFiles, ".", "")
	listSize := global.BA_CONFIG.Ueditor.FileManagerListSize
	if allowFiles == "" || listSize == 0 || len(get) <= 0 {
		result["state"] = "config params error."
	}

	listSizeInt := listSize
	result = us.fileList(allowFiles, listSizeInt, get)

	return result
}

// UploadScrawl 上传涂鸦
func (us *UeditorService) UploadScrawl(get url.Values) map[string]interface{} {
	result := make(map[string]interface{})
	pathFormat := global.BA_CONFIG.Ueditor.ScrawlPathFormat
	maxSize := global.BA_CONFIG.Ueditor.ScrawlMaxSize
	allowFiles := global.BA_CONFIG.Ueditor.ScrawlAllowFiles
	//ext前面的.号去掉
	allowFiles = strings.ReplaceAll(allowFiles, ".", "")
	oriName := global.BA_CONFIG.Ueditor.ScrawlFieldName

	if pathFormat == "" || maxSize == 0 || allowFiles == "" || oriName == "" {
		result["state"] = "config params error."
		return result
	}

	config := map[string]string{
		"pathFormat": pathFormat,
		"maxSize":    strconv.Itoa(maxSize),
		"allowFiles": allowFiles,
		"oriName":    oriName,
	}

	base64Data := get.Get(oriName)
	return us.upBase64(config, base64Data)
}

// upBase64 处理base64编码的图片上传
func (us *UeditorService) upBase64(config map[string]string, base64Data string) map[string]interface{} {
	result := make(map[string]interface{})
	imgByte, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		result["state"] = "picture content get error. err:" + err.Error()
		return result
	}

	path := "/static/uploads/ueditor/scrawl/"
	dirName := filepath.ToSlash(beego.AppPath + path)
	file := make(map[string]string)
	file["filesize"] = strconv.Itoa(len(string(imgByte)))
	file["oriName"] = config["oriName"]
	file["ext"] = ".png"
	file["name"] = uuid.New().String() + file["ext"]
	file["fullName"] = dirName + file["name"]
	file["urlName"] = path + file["name"]

	fullName := file["fullName"]

	//检查文件大小是否超出限制
	fileSizeInt, _ := strconv.Atoi(file["filesize"])
	maxSizeInt, _ := strconv.Atoi(config["maxSize"])
	if fileSizeInt >= maxSizeInt {
		result["state"] = "文件大小超出网站限制"
		return result
	}

	//创建目录失败
	_, err = os.Stat(dirName)
	if err != nil {
		err = os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			result["state"] = "目录创建失败"
			return result
		}
	}

	//写入文件
	f, err := os.OpenFile(fullName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		result["state"] = "写入文件内容错误,err :" + err.Error()
		return result
	}
	defer f.Close()
	_, err = f.Write(imgByte)
	if err != nil {
		result["state"] = "写入文件失败,err :" + err.Error()
		return result
	}

	result = map[string]interface{}{
		"state":    "SUCCESS",
		"url":      file["urlName"],
		"title":    file["name"],
		"original": file["oriName"],
		"type":     file["ext"],
		"size":     file["filesize"],
	}

	return result
}

// CatchImage 抓取远程文件
func (us *UeditorService) CatchImage(ctx *context.Context) map[string]interface{} {
	result := make(map[string]interface{})
	pathFormat := global.BA_CONFIG.Ueditor.CatcherPathFormat
	maxSize := global.BA_CONFIG.Ueditor.CatcherMaxSize
	allowFiles := global.BA_CONFIG.Ueditor.CatcherAllowFiles
	//ext前面的.号去掉
	allowFiles = strings.ReplaceAll(allowFiles, ".", "")
	oriName := "remote.png"

	if pathFormat == "" || maxSize == 0 || allowFiles == "" {
		result["state"] = "config params error."
		return result
	}

	config := map[string]string{
		"pathFormat": pathFormat,
		"maxSize":    strconv.Itoa(maxSize),
		"allowFiles": allowFiles,
		"oriName":    oriName,
	}

	fieldName := global.BA_CONFIG.Ueditor.CatcherFieldName

	source := make([]string, 0)
	ctx.Input.Bind(&source, fieldName)

	var list []map[string]string
	//没有数据
	if len(source) <= 0 {
		result = map[string]interface{}{
			"state": "ERROR",
			"list":  list,
		}
		return result
	}

	for _, imgURL := range source {
		info := us.saveRemote(config, imgURL)
		if info["state"] == "SUCCESS" {
			list = append(list, map[string]string{
				"state":    info["state"],
				"url":      info["url"],
				"size":     info["size"],
				"title":    info["title"],
				"original": info["original"],
				"source":   imgURL,
			})
		} else {
			list = append(list, map[string]string{
				"state":    info["state"],
				"url":      "",
				"size":     "",
				"title":    "",
				"original": "",
				"source":   imgURL,
			})
		}

	}

	result = map[string]interface{}{
		"state": "SUCCESS",
		"list":  list,
	}

	return result
}

// saveRemote 抓取远程图片
func (*UeditorService) saveRemote(config map[string]string, fieldName string) map[string]string {
	result := make(map[string]string)
	imgURL := strings.ReplaceAll(fieldName, "&amp;", "&")

	if imgURL == "" {
		result["state"] = "链接为空"
		return result
	}

	//http开头验证
	if !strings.HasPrefix(imgURL, "http") {
		result["state"] = "链接不是http链接"
		return result
	}

	//获取请求头并检测死链
	response, err := http.Get(imgURL)
	defer response.Body.Close()
	if err != nil || response.StatusCode != 200 {
		result["state"] = "链接不可用"
		return result
	}

	//格式验证(扩展名验证和Content-Type验证)
	if !strings.Contains(response.Header.Get("Content-Type"), "image") {
		result["state"] = "链接contentType不正确"
		return result
	}

	fileType := strings.TrimLeft(filepath.Ext(imgURL), ".")
	if fileType == "" || !strings.Contains(config["allowFiles"], fileType) {
		result["state"] = "链接url后缀不正确"
		return result
	}

	path := "/static/uploads/ueditor/remote/"
	dirName := filepath.ToSlash(beego.AppPath + path)

	file := make(map[string]string)
	file["oriName"] = filepath.Ext(imgURL)
	file["filesize"] = "0"
	file["ext"] = file["oriName"]
	file["name"] = uuid.New().String() + file["ext"]
	file["fullName"] = dirName + file["name"]
	file["urlName"] = path + file["name"]

	//检查文件大小是否超出限制
	fileSizeInt, _ := strconv.Atoi(file["filesize"])
	maxSizeInt, _ := strconv.Atoi(config["maxSize"])
	if fileSizeInt >= maxSizeInt {
		result["state"] = "文件大小超出网站限制"
		return result
	}

	//创建目录失败
	_, err = os.Stat(dirName)
	if err != nil {
		err = os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			result["state"] = "目录创建失败"
			return result
		}
	}

	//写入文件
	img := response.Body
	f, err := os.Create(file["fullName"])

	if err != nil {
		result["state"] = "写入文件失败,err :" + err.Error()
		return result
	}

	w, err := io.Copy(f, img)
	if err != nil {
		result["state"] = "写入文件失败,err :" + err.Error()
		return result
	}

	file["filesize"] = strconv.Itoa(int(w))

	result = map[string]string{
		"state":    "SUCCESS",
		"url":      file["urlName"],
		"title":    file["name"],
		"original": file["oriName"],
		"type":     file["ext"],
		"size":     file["filesize"],
	}

	return result

}
