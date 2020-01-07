package services

import "github.com/article-publish-server/models"

type ImageService interface {
	ArticleCoverUptoken() *models.Ret
	ImageUptoken() *models.Ret
	ImageUEUptoken() *models.Ret
	//ImageUploadCb(body *qnuploader.UploadImageCbBody) *models.Ret
}
