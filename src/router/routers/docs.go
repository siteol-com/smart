package routers

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/service/comm/commHander"
)

// DocsRouter API文档路由
func DocsRouter(router *gin.Engine) {
	// API 本地调试路由
	// initSwagger
	docsRouter := router.Group("/docs")
	{
		// Swagger资源文件
		docsRouter.POST("/sample", commHander.Sample)
		// Swagger资源文件
		docsRouter.GET("/file/*any", commHander.DocsFile)
		// ReDoc
		docsRouter.GET("/redoc/*any", commHander.ReDoc)
		// Swagger范本
		docsRouter.GET("/swagger/*any", commHander.SwaggerDoc)
	}
}
