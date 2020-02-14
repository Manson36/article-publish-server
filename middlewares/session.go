package middlewares

import (
	"github.com/article-publish-server/config"
	"github.com/article-publish-server/models"
	"github.com/article-publish-server/repositories"
	"github.com/article-publish-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Session(basePath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		errFunc := func(code int, msg string) {
			c.Abort() //？？使用终止
			c.JSON(http.StatusOK, models.Ret{Code: code, Msg: msg, TokenInvalid: true})
		}

		if utils.IndexOfWithString(
			utils.SessionWithList(),
			utils.RelativePath(c, basePath),
		) != -1 {
			c.Next()
			return
		}

		cookie, err := c.Cookie(config.Web.TokenKey)
		if err != nil {
			errFunc(400, "当前用户还未登录")
			return
		}

		au := &models.AdminUserCustomClaims{}
		if err := au.Parse(cookie); err != nil || au.UserID == 0 {
			errFunc(400, "用户登陆已失效，请重新登陆")
			return
		}

		repo := repositories.NewAdminUserRepository()
		user, err := repo.GetInfoFromCacheByID(au.UserID)
		if err != nil {
			errFunc(500, "用户缓存信息获取失败")
			return
		}

		if user == nil {
			user, err = repo.Get("id = ?", au.UserID)
			if err != nil {
				errFunc(500, "用户信息获取失败，请与平台联系")
				return
			}

			if user == nil {
				errFunc(400, "用户信息不存在，请重新登陆")
			}

			_ = repo.SetInfoToCache
		}

		if user.Removed {
			errFunc(400, "用户已被移除")
			return
		}

		if user.Disabled {
			errFunc(500, "用户已被禁用")
			return
		}

		c.Set("session", user)
		c.Next()
	}
}
