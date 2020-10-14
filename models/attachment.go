package models

import (
	"beego-admin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"strconv"
)

type Attachment struct {
	Id           int    `orm:"column(id);auto;size(11)" description:"表ID"`
	AdminUserId  int    `orm:"column(admin_user_id);size(11);default(0)" description:"后台用户id"`
	UserId       int    `orm:"column(user_id);size(11);default(0)" description:"前台用户ID"`
	OriginalName string `orm:"column(original_name);size(200)" description:"原文件名"`
	SaveName     string `orm:"column(save_name);size(200)" description:"保存文件名"`
	SavePath     string `orm:"column(save_path);size(255)" description:"系统完整路径"`
	Url          string `orm:"column(url);size(255)" description:"图片访问路径"`
	Extension    string `orm:"column(extension);size(100)" description:"后缀"`
	Mime         string `orm:"column(mime);size(100)" description:"类型"`
	Size         int64  `orm:"column(size);size(20);default(0)" description:"大小"`
	Md5          string `orm:"column(md5);size(32);default(\"\")" description:"MD5"`
	Sha1         string `orm:"column(sha1);size(40);default(\"\")" description:"SHA1"`
	CreateTime   int    `orm:"column(create_time);size(10);default(0)" description:"操作时间"`
	UpdateTime   int    `orm:"column(update_time);size(10);default(0)" description:"更新时间"`
	DeleteTime   int    `orm:"column(delete_time);size(10);default(0)" description:"删除时间"`
}

//自定义table 名称
func (*Attachment) TableName() string {
	return "attachment"
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(Attachment))
}

func (*Attachment) FileType() map[string][]string {
	return map[string][]string{
		"图片":   {"jpg", "bmp", "png", "jpeg", "gif", "svg"},
		"文档":   {"txt", "doc", "docx", "xls", "xlsx", "pdf"},
		"压缩文件": {"rar", "zip", "7z", "tar"},
		"音视":   {"mp3", "ogg", "flac", "wma", "ape"},
		"视频":   {"mp4", "wmv", "avi", "rmvb", "mov", "mpg"},
	}
}

func (*Attachment) FileThumb() map[string][]string {
	return map[string][]string{
		"picture":      {"jpg", "bmp", "png", "jpeg", "gif", "svg"},
		"txt.svg":      {"txt", "pdf"},
		"pdf.svg":      {"pdf"},
		"word.svg":     {"doc", "docx"},
		"excel.svg":    {"xls", "xlsx"},
		"archives.svg": {"rar", "zip", "7z", "tar"},
		"audio.svg":    {"mp3", "ogg", "flac", "wma", "ape"},
		"video.svg":    {"mp4", "wmv", "avi", "rmvb", "mov", "mpg"},
	}
}

//格式化大小
func (attachment *Attachment) GetSize() string {
	size := float64(attachment.Size)
	units := []string{" B", " KB", " MB", " GB", " TB"}
	var i int
	for i = 0; size >= 1024 && i < 4; i++ {
		size /= 1024
	}
	return strconv.FormatFloat(math.Round(size), 'f', -1, 64) + units[i]
}

//文件分类
func (attachment *Attachment) GetFileType() string {
	typeName := "其他"
	extension := attachment.Extension
	for name, arr := range attachment.FileType() {
		if utils.InArrayForString(arr, extension) {
			typeName = name
			break
		}
	}
	return typeName
}

//文件预览
func (attachment *Attachment) GetThumbnail() string {
	thumbnail := beego.AppConfig.String("attachment::thumb_path") + "unknown.svg"
	extension := attachment.Extension
	thumbPath := beego.AppConfig.String("attachment::thumb_path")

	fileThumb := map[string][]string{
		"picture":                  {"jpg", "bmp", "png", "jpeg", "gif", "svg"},
		thumbPath + "txt.svg":      {"txt", "pdf"},
		thumbPath + "pdf.svg":      {"pdf"},
		thumbPath + "word.svg":     {"doc", "docx"},
		thumbPath + "excel.svg":    {"xls", "xlsx"},
		thumbPath + "archives.svg": {"rar", "zip", "7z", "tar"},
		thumbPath + "audio.svg":    {"mp3", "ogg", "flac", "wma", "ape"},
		thumbPath + "video.svg":    {"mp4", "wmv", "avi", "rmvb", "mov", "mpg"},
	}

	for name, arr := range fileThumb {
		if utils.InArrayForString(arr, extension) {
			if name == "picture" {
				thumbnail = attachment.Url
			} else {
				thumbnail = name
			}
			break
		}
	}

	return thumbnail
}
