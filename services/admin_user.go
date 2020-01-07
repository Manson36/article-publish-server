package services

import (
	"github.com/article-publish-server/datamodels"
	"github.com/article-publish-server/models"
	"github.com/article-publish-server/repositories"
	"github.com/article-publish-server/utils"
	"log"
	"strings"
)

type AdminUserService interface {
	Create(*models.AdminUserAddReqBody) *models.Ret
	Login(*models.AdminUserLoginReqBody) *models.Ret
}

type adminUserService struct {
	repo repositories.AdminUserRepository
}

func NewAdminUserService() AdminUserService {
	repo := repositories.NewAdminUserRepository()
	return adminUserService{repo: repo}
}

func (a adminUserService) Create(body *models.AdminUserAddReqBody) *models.Ret {
	id, err := utils.GetInt64ID() //这里是啥》》？？？
	if err != nil {
		log.Println("创建账号时，获取生成id错误：", err.Error())
		return &models.Ret{Code: 500, Msg: "创建账号时生成id错误"}
	}

	switch body.Platform {
	case datamodels.ZingglobalPlatform, datamodels.ZHIdreamPaltform, datamodels.HealthEnginePlatform:
	default:
		return &models.Ret{Code: 400, Msg: "请输入正确的平台类型"}
	}

	//trimSpace是怎么获取内容？？？
	pwd := strings.TrimSpace(body.Password)
	if pwd == "" {
		return &models.Ret{Code: 400, Msg: "请输入管理员密码"}
	}

	email := strings.TrimSpace(body.Email)
	if email == "" {
		return &models.Ret{Code: 400, Msg: "请输入邮箱"}
	}

	nickName := strings.TrimSpace(body.NickName)
	if nickName == "" {
		return &models.Ret{Code: 400, Msg: "请输入昵称"}
	}

	//？？？？？
	pwdInfo := utils.GenPwdAndSalt(pwd)
	user := datamodels.AdminUser{
		ID:           id,
		NickName:     nickName,
		Email:        email,
		Password:     pwd,
		Salt:         pwdInfo.Salt,
		AdminType:    2,
		PlatformType: body.Platform,
	}

	if body.IsAdmin {
		user.AdminType = 1
	}

	u, err := a.repo.Get(`removed IS NOT TRUE AND email = ? AND platform_type= ?`, email, body.Platform)
	if err != nil {
		log.Println("用户信息获取失败，数据库错误，errmsg：", err.Error())
		return &models.Ret{Code: 500, Msg: "用户信息获取失败, 请与平台联系"}
	}

	if u != nil {
		return &models.Ret{Code: 400, Msg: "该账号存在"}
	}

	if err := a.repo.Create(&user); err != nil {
		log.Println("创建用户错误：", err.Error())
		return &models.Ret{Code: 500, Msg: "创建用户错误"}
	}

	return &models.Ret{Code: 200, Msg: "用户创建成功", Data: user}
}

func (a adminUserService) Login(body *models.AdminUserLoginReqBody) *models.Ret {
	//switch body.Platform {
	//case datamodels.ZingglobalPlatform, datamodels.ZhidreamPlatform, datamodels.HealthEnginePlatform:
	//default:
	//	return &models.Ret{Code: 400, Msg: "请输入正确的平台类型"}
	//}

	// user, err := a.repo.Get("email = ? AND platform_type = ? AND removed IS NOT TRUE", body.Email, body.Platform)
	user, err := a.repo.Get("email = ? AND removed IS NOT TRUE", body.Email)
	if err != nil {
		log.Println("管理用户信息获取失败：", err.Error())
		return &models.Ret{Code: 500, Msg: "管理用户信息获取失败，请与平台联系"}
	}

	if user == nil {
		return &models.Ret{Code: 400, Msg: "该用户不存在"}
	}

	//这个需要强记吗？？？
	if utils.HashPwdWithSalt(body.Password, user.Salt) != user.Password {
		return &models.Ret{Code: 400, Msg: "密码输入错误"}
	}

	//里面包含jwt，怎么使用？？？
	claims := models.AdminUserCustomClaims{}
	claims.UserID = user.ID
	token, err := claims.Sign()
	if err != nil {
		log.Println("管理用户登陆token生成失败：", err.Error())
		return &models.Ret{Code: 500, Msg: "管理用户登陆token生成失败"}
	}

	data := models.AdminUserLoginResBody{
		User:  user,
		Token: token,
	}

	return &models.Ret{Code: 200, Msg: "登陆成功", Data: data}
}
