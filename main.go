package main

import (
	"github.com/article-publish-server/config"
	"github.com/article-publish-server/controllers"
	"github.com/article-publish-server/middlewares"
	"github.com/article-publish-server/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	if err := web(); err != nil {
		log.Fatal(err.Error())
	}
}

func web() error {
	switch config.Server.Mode {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "dev":
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	//basePath是指什么？？？session待定解决？？？
	r.Use(middlewares.Session(r.BasePath()))
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, config.Server.Name+": pong")
	}).GET("/name", func(c *gin.Context) {
		c.String(http.StatusOK, config.Server.Name)
	})

	//admin_user
	adminUserController := controllers.AdminUserController{
		Service: services.NewAdminUserService(),
	}
	r.
		Group("admin/user").
		POST("/login", adminUserController.Login).
		POST("/create", adminUserController.Create).
		POST("/info", adminUserController.Info)

	//image
	imageController := controllers.ImageController{}
	r.
		Group("/image").
		GET("/ue", imageController.GetUEConfig).
		POST("/ue", imageController.UploadUEFile).
		POST("/uptoken", imageController.UpToken).
		POST("/upload/cb", imageController.UploadCb).
		POST("/list", imageController.GetList).
		POST("/remove", imageController.Remove)

	// article
	articleController := controllers.ArticleController{}
	r.
		Group("/article").
		POST("/uptoken", articleController.Uptoken).
		POST("/upload/cb", articleController.UploadCb).
		POST("/create", articleController.Create).
		POST("/remove", articleController.Remove).
		POST("/update", articleController.Update).
		POST("/info", articleController.Info).
		POST("/list", articleController.List)

	// tag
	tagController := controllers.TagController{}
	r.
		Group("/tag").
		POST("/create", tagController.Create).
		POST("/remove", tagController.Remove).
		POST("/update", tagController.Update).
		POST("/info", tagController.Get).
		POST("/list", tagController.ListAll)

	return r.Run(":" + config.Web.Port)
}
