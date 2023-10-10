package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin/blog-service/docs"
	"github.com/gin/blog-service/global"
	"github.com/gin/blog-service/internal/middleware"
	"github.com/gin/blog-service/internal/routers/api/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter() *gin.Engine {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(middleware.Translations())
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	e.POST("/auth", v1.GetAuth)

	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := v1.NewUpload()
	e.POST("/upload/file", upload.UploadFile)
	e.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	apiv1 := e.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)

	}

	return e
}
