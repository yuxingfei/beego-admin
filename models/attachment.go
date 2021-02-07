package models

import (
	"beego-admin/global"
	"beego-admin/utils"
	"github.com/beego/beego/v2/client/orm"
	"math"
	"strconv"
)

// Attachment struct
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

// TableName 自定义table 名称
func (*Attachment) TableName() string {
	return "attachment"
}

// SearchField 定义模型的可搜索字段
func (*Attachment) SearchField() []string {
	return []string{}
}

// NoDeletionId 禁止删除的数据id
func (*Attachment) NoDeletionId() []int {
	return []int{}
}

// WhereField 定义模型可作为条件的字段
func (*Attachment) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*Attachment) TimeField() []string {
	return []string{}
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(Attachment))
}

// FileType 属性定义
func (*Attachment) FileType() map[string][]string {
	return map[string][]string{
		"图片":   {"jpg", "bmp", "png", "jpeg", "gif", "svg"},
		"文档":   {"txt", "doc", "docx", "xls", "xlsx", "pdf"},
		"压缩文件": {"rar", "zip", "7z", "tar"},
		"音视":   {"mp3", "ogg", "flac", "wma", "ape"},
		"视频":   {"mp4", "wmv", "avi", "rmvb", "mov", "mpg"},
	}
}

// FileThumb 属性定义
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

// GetSize 格式化大小
func (attachment *Attachment) GetSize() string {
	size := float64(attachment.Size)
	units := []string{" B", " KB", " MB", " GB", " TB"}
	var i int
	for i = 0; size >= 1024 && i < 4; i++ {
		size /= 1024
	}
	return strconv.FormatFloat(math.Round(size), 'f', -1, 64) + units[i]
}

// GetFileType 文件分类
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

// GetThumbnail 文件预览
func (attachment *Attachment) GetThumbnail() string {
	thumbnail := global.BA_CONFIG.Attachment.ThumbPath + "unknown.svg"
	extension := attachment.Extension
	thumbPath := global.BA_CONFIG.Attachment.ThumbPath

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
