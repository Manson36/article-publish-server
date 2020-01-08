package services

import (
	"github.com/article-publish-server/datamodels"
	"github.com/article-publish-server/models"
	"github.com/article-publish-server/repositories"
	"github.com/article-publish-server/utils"
	"log"
	"strings"
)

type TagService interface {
	Create(body *models.TagCreateReqBody) *models.Ret
	Remove(body *models.TagRemoveReqBody) *models.Ret
	Get(body *models.TagInfoReqBody) *models.Ret
	Update(body *models.TagUpdateReqBody) *models.Ret
	ListAll(body *models.TagListReqBody) *models.Ret
}

type tagService struct {
	repo repositories.TagRepository
}

func NewTagService() TagService {
	return &tagService{repo: repositories.NewTagRepository()}
}

func (t tagService) Create(body *models.TagCreateReqBody) *models.Ret {
	switch body.Platform {
	case datamodels.ZingglobalPlatform, datamodels.ZhidreamPlatform, datamodels.HealthEnginePlatform:
	default:
		return &models.Ret{Code: 400, Msg: "请输入正确的平台类型"}
	}

	name := strings.TrimSpace(body.Name)
	if name == "" {
		return &models.Ret{Code: 400, Msg: "请输入标签标题"}
	}

	tag, err := t.repo.Get("removed IS NOT TRUE AND platform_type = ? AND name = ?", body.Platform, body.Name)
	if err != nil {
		log.Println("标签信息获取失败，数据库错误，errmsg：", err.Error())
		return &models.Ret{Code: 500, Msg: "标签信息获取失败，请与平台联系"}
	}

	if tag != nil {
		return &models.Ret{Code: 400, Msg: "标签名称重复"}
	}

	id, err := utils.GetInt64ID()
	if err != nil {
		log.Println("创建标签时，获取id信息错误，errmsg：", err.Error())
		return &models.Ret{Code: 501, Msg: "创建标签时，获取id信息错误，请与平台联系"}
	}

	_tag := datamodels.Tag{ID: id, Name: name, PlatformType: body.Platform}
	if err := t.repo.Create(&_tag); err != nil {
		log.Println("创建标签失败，errmsg：", err.Error())
		return &models.Ret{Code: 501, Msg: "创建标签失败，请与平台联系"}
	}

	return &models.Ret{Code: 200, Msg: "标签添加成功", Data: _tag}
}

func (t tagService) Remove(body *models.TagRemoveReqBody) *models.Ret {
	switch body.Platform {
	case datamodels.ZingglobalPlatform, datamodels.ZhidreamPlatform, datamodels.HealthEnginePlatform:
	default:
		return &models.Ret{Code: 400, Msg: "请输入正确的平台类型"}
	}

	if body.ID == 0 {
		return &models.Ret{Code: 400, Msg: "请传入正确的标签唯一标识"}
	}

	if err := t.repo.Remove("id = ? AND platform = ?", body.ID, body.Platform); err != nil {
		log.Println("标签删除失败，数据库错误，errmsg:", err.Error())
		return &models.Ret{Code: 500, Msg: "标签删除失败，请与平台联系"}
	}

	return &models.Ret{Code: 200, Msg: "标签删除成功"}
}

func (t tagService) Get(body *models.TagInfoReqBody) *models.Ret {
	switch body.Platform {
	case datamodels.ZingglobalPlatform, datamodels.ZhidreamPlatform, datamodels.HealthEnginePlatform:
	default:
		return &models.Ret{Code: 400, Msg: "请输入正确的平台类型"}
	}

	tag, err := t.repo.Get("removed IS NOT TRUE AND platform_type = ? AND id = ?", body.Platform, body.ID)
	if err != nil {
		log.Println("获取标签信息失败，数据库错误, errmsg:", err.Error())
		return &models.Ret{Code: 500, Msg: "标签信息获取失败，请与平台联系"}
	}

	return &models.Ret{Code: 200, Msg: "标签信息获取成功", Data: tag}
}

func (t tagService) Update(body *models.TagUpdateReqBody) *models.Ret {
	switch body.Platform {
	case datamodels.ZingglobalPlatform, datamodels.ZhidreamPlatform, datamodels.HealthEnginePlatform:
	default:
		return &models.Ret{Code: 400, Msg: "请输入正确的平台类型"}
	}

	if body.ID == 0 {
		return &models.Ret{Code: 400, Msg: "请输入正确的标签唯一标识"}
	}

	name := strings.TrimSpace(body.Name)
	tag, err := t.repo.Get("removed IS NOT TRUE AND id <> ? AND platform_type = ? AND name = ?", body.ID, body.Platform, name)
	if err != nil {
		log.Println("标签更新标签信息获取失败，数据库错误，errmsg：", err.Error())
		return &models.Ret{Code: 500, Msg: "标签信息获取失败，请与平台联系"}
	}

	if tag != nil {
		return &models.Ret{Code: 400, Msg: "标签名称重复"}
	}

	if err := t.repo.Update(map[string]interface{}{"name": name}, `id = ? AND platform_type = ?`, body.ID, body.Platform); err != nil {
		log.Println("标签更新失败，数据库错误，errmsg：", err.Error())
		return &models.Ret{Code: 500, Msg: "标签更新失败，请与平台联系"}
	}

	return &models.Ret{Code: 200, Msg: "标签信息更新成功"}
}

func (t tagService) ListAll(body *models.TagListReqBody) *models.Ret {
	switch body.Platform {
	case datamodels.ZingglobalPlatform, datamodels.ZhidreamPlatform, datamodels.HealthEnginePlatform:
	default:
		return &models.Ret{Code: 400, Msg: "请输入正确的平台类型"}
	}

	tags, err := t.repo.ListAll(`created_at ASC`, `platform_type = ? AND removed IS NOT TRUE`, body.Platform)
	if err != nil {
		log.Println("标签列表获取失败，数据库错误，errmsg：", err.Error())
		return &models.Ret{Code: 500, Msg: "标签列表获取失败，请与平台联系"}
	}

	if len(tags) == 0 {
		tags = make([]datamodels.Tag, 0)
	}

	return &models.Ret{Code: 200, Msg: "标签列表获取成功", Data: tags}
}
