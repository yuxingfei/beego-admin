package services

import (
	"beego-admin/global"
	"beego-admin/models"
	"beego-admin/utils"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// AttachmentService struct
type AttachmentService struct {
}

// Upload 上传单个文件
func (*AttachmentService) Upload(ctx *context.Context, name string, adminUserId int, userId int) (*models.Attachment, error) {
	file, h, err := ctx.Request.FormFile(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//自定义文件验证
	err = validateForAttachment(h)
	if err != nil {
		return nil, err
	}

	//数据表写入
	saveName := uuid.New().String()
	//后缀带. (.png)
	fileExt := path.Ext(h.Filename)
	savePath := global.BA_CONFIG.Attachment.Path + saveName + fileExt
	saveRealDir := filepath.ToSlash(beego.AppPath + "/" + global.BA_CONFIG.Attachment.Path)

	_, err = os.Stat(saveRealDir)
	if err != nil {
		err = os.MkdirAll(saveRealDir, os.ModePerm)
	}

	saveURL := "/" + global.BA_CONFIG.Attachment.Url + saveName + fileExt

	f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	io.Copy(f, file)

	attachmentInfo := models.Attachment{
		AdminUserId:  adminUserId,
		UserId:       userId,
		OriginalName: h.Filename,
		SaveName:     saveName,
		SavePath:     saveRealDir + saveName + fileExt,
		Url:          saveURL,
		Extension:    strings.TrimLeft(fileExt, "."),
		Mime:         h.Header.Get("Content-Type"),
		Size:         h.Size,
		Md5:          utils.GetMd5String(saveName),
		Sha1:         utils.GetSha1String(saveName),
		CreateTime:   int(time.Now().Unix()),
		UpdateTime:   int(time.Now().Unix()),
	}

	insertID, err := orm.NewOrm().Insert(&attachmentInfo)
	if err == nil {
		attachmentInfo.Id = int(insertID)
		return &attachmentInfo, nil
	}
	return nil, err
}

// UploadMulti 上传多个文件
func (*AttachmentService) UploadMulti(ctx *context.Context, name string, adminUserId int, userId int) ([]*models.Attachment, error) {
	var result []*models.Attachment
	//GetFiles return multi-upload files
	files, ok := ctx.Request.MultipartForm.File[name]
	if !ok {
		return nil, http.ErrMissingFile
	}

	for i := range files {
		h := files[i]
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			return nil, err
		}
		////create destination file making sure the path is writeable.
		//dst, err := os.Create("upload/" + files[i].Filename)
		//defer dst.Close()
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		////copy the uploaded file to the destination file
		//if _, err := io.Copy(dst, file); err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}

		//自定义文件验证
		err = validateForAttachment(h)
		if err != nil {
			return nil, err
		}

		//数据表写入
		saveName := uuid.New().String()
		//后缀带. (.png)
		fileExt := path.Ext(h.Filename)
		savePath := global.BA_CONFIG.Attachment.Path + saveName + fileExt
		saveRealDir := filepath.ToSlash(beego.AppPath + "/" + global.BA_CONFIG.Attachment.Path)

		_, err = os.Stat(saveRealDir)
		if err != nil {
			err = os.MkdirAll(saveRealDir, os.ModePerm)
		}

		saveURL := "/" + global.BA_CONFIG.Attachment.Url + saveName + fileExt

		f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		io.Copy(f, file)

		attachmentInfo := models.Attachment{
			AdminUserId:  adminUserId,
			UserId:       userId,
			OriginalName: h.Filename,
			SaveName:     saveName,
			SavePath:     saveRealDir + saveName + fileExt,
			Url:          saveURL,
			Extension:    strings.TrimLeft(fileExt, "."),
			Mime:         h.Header.Get("Content-Type"),
			Size:         h.Size,
			Md5:          utils.GetMd5String(saveName),
			Sha1:         utils.GetSha1String(saveName),
			CreateTime:   int(time.Now().Unix()),
			UpdateTime:   int(time.Now().Unix()),
		}

		insertID, err := orm.NewOrm().Insert(&attachmentInfo)
		if err == nil {
			attachmentInfo.Id = int(insertID)
			//return &attachmentInfo, nil
			result = append(result, &attachmentInfo)
		} else {
			return nil, err
		}
	}

	//返回上传文件的对象信息
	if result != nil {
		return result, nil
	}
	return nil, errors.New("无法获取文件")

}

// validateForAttachment attachment自定义验证
func validateForAttachment(h *multipart.FileHeader) error {
	validateSize, _ :=  strconv.Atoi(global.BA_CONFIG.Attachment.ValidateSize)
	validateExt := global.BA_CONFIG.Attachment.ValidateExt
	if int(h.Size) > validateSize {
		return errors.New("文件超过限制大小")
	}

	if !strings.Contains(validateExt, strings.ToLower(strings.TrimLeft(path.Ext(h.Filename), "."))) {
		return errors.New("不支持的文件格式")
	}

	return nil
}
