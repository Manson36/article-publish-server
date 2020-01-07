package controllers

import (
	"github.com/article-publish-server/config"
	"github.com/article-publish-server/datamodels"
	"github.com/article-publish-server/models"
	"github.com/article-publish-server/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AdminUserController struct {
	Service services.AdminUserService
}

func (a *AdminUserController) Login(c *gin.Context) {
	body := models.AdminUserLoginReqBody{}
	//问题1：shouldBindJSON???验证
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusOK, models.Ret{Code: 400, Msg: "用户登陆参数解析错误", Data: nil})
		return
	}

	var hostname string
	//问题2：tokendomain作用？？
	if config.Web.TokenDomain != "" {
		hostname = config.Web.TokenDomain
	} else {
		hostname = strings.Split(c.Request.Host, ":")[0]
	}

	data := a.Service.Login(&body)
	if data.Code == 200 {
		token := data.Data.(models.AdminUserLoginResBody).Token
		c.SetCookie(config.Web.TokenKey, token, 3600*24*config.Web.ExpiresAt, "/", hostname, false, false)
	}

	c.JSON(http.StatusOK, data)
}

func (a *AdminUserController) Create(c *gin.Context) {
	body := models.AdminUserAddReqBody{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusOK, models.Ret{Code: 400, Msg: "创建管理用户参数解析失败", Data: nil})
		return
	}

	c.JSON(http.StatusOK, a.Service.Create(&body))
}

func (a *AdminUserController) Info(c *gin.Context) {
	u, ok := c.Get("session")
	if !ok {
		c.JSON(http.StatusOK, models.Ret{Code: 400, Msg: "用户未登录，请登陆", Data: nil, TokenInvalid: true})
		return
	}

	user, ok := u.(*datamodels.AdminUser)
	if !ok {
		c.JSON(http.StatusOK, models.Ret{Code: 400, Msg: "用户未登录，请登录", Data: nil, TokenInvalid: true})
		return
	}

	if user == nil {
		c.JSON(http.StatusOK, models.Ret{Code: 400, Msg: "用户未登录，请登录", Data: nil, TokenInvalid: true})
		return
	}

	c.JSON(http.StatusOK, models.Ret{Code: 200, Msg: "用户信息获取成功", Data: *user})
	return
}
