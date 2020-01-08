package services

import (
	"github.com/article-publish-server/datamodels"
	"github.com/article-publish-server/models"
	"github.com/article-publish-server/repositories"
	"github.com/article-publish-server/utils"
	"github.com/article-publish-server/utils/qnuploader"
	"log"
	"mime/multipart"
)

type ImageService interface {
	ArticleCoverUptoken() *models.Ret
	ImageUptoken() *models.Ret
	ImageUEUptoken() *models.Ret
	ImageUploadCb(body *qnuploader.UploadImageCbBody) *models.Ret //又是七牛？？
	CreateImageByUploadBody(body *qnuploader.UploadImageCbBody) (*datamodels.Image, *models.Ret)
	GetList(body *models.ImageListReqBody) *models.Ret
	RemovedImage(body *models.ImageRemoveReqBody) *models.Ret
	UEImageUpload(file *multipart.FileHeader) map[string]interface{}
}

type imageService struct {
	uploader *qnuploader.Uploader
	repo     repositories.ImageRepository
	fileRepo repositories.FileRepository
}

func NewImageService() ImageService {
	uploader := qnuploader.NewUploader(nil)
	return &imageService{
		uploader: uploader,
		repo:     repositories.NewImageRepository(),
		fileRepo: repositories.NewFileRepository(),
	}
}

//???
func (i imageService) ArticleCoverUptoken() *models.Ret {
	return &models.Ret{
		Code: 200,
		Msg:  "获取文章封面图片上传凭证成功",
		Data: map[string]string{
			"uptoken": i.uploader.GetImageUptoken(
				i.uploader.CallbackURI+"article/upload/cb",
				nil,
			),
		},
	}
}

func (i imageService) ImageUptoken() *models.Ret {
	return &models.Ret{
		Code: 200,
		Msg:  "获取图片上传凭证成功",
		Data: map[string]string{
			"uptoken": i.uploader.GetImageUptoken(
				i.uploader.CallbackURI+"image/upload/cb", nil),
		},
	}
}

func (i imageService) ImageUEUptoken() *models.Ret {
	return &models.Ret{
		Code: 200,
		Msg:  "获取UE图片上传凭证成功",
		Data: map[string]string{
			"uptoken": i.uploader.GetImageUptoken(
				i.uploader.CallbackURI+"image/upload/ue/cb", nil),
		},
	}
}

func (i imageService) ImageUploadCb(body *qnuploader.UploadImageCbBody) *models.Ret {
	panic("implement me")
}

func (i imageService) CreateImageByUploadBody(body *qnuploader.UploadImageCbBody) (*datamodels.Image, *models.Ret) {
	panic("implement me")
}

func (i imageService) GetList(body *models.ImageListReqBody) *models.Ret {
	panic("implement me")
}

func (i imageService) RemovedImage(body *models.ImageRemoveReqBody) *models.Ret {
	panic("implement me")
}

//???
func (i imageService) UEImageUpload(file *multipart.FileHeader) map[string]interface{} {
	uptoken := i.uploader.GetImageUptoken("", nil)
	body := qnuploader.UploadImageCbBody{}

	key := "ow/platform/image" + utils.GetRandomString(32)
	params := map[string]string{
		"x:uploader": "1",
		"x:filename": file.Filename,
		"x:platform": "1",
	}

	if err := i.uploader.UploadFormFile(&body, params, uptoken, key, file); err != nil {
		log.Println("编辑器上传图片失败，errmsg：", err.Error())
		return map[string]interface{}{"state": "FAIL", "msg": "编辑器上传图片失败，请与平台联系"}
	}

	img, ret := i.CreateImageByUploadBody(&body)
	if ret != nil {
		return map[string]interface{}{"state": "FAIL", "msg": ret.Msg}
	}

	return map[string]interface{}{"state": "SUCCESS", "url": img.Path, "title": img.Name, "original": img.Name}
}
