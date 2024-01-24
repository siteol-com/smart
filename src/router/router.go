package router

import (
	"io"
	"siteol.com/smart/src/router/middleware"
	"siteol.com/smart/src/router/routers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	gin.ForceConsoleColor() // 颜色日志
	// 基础路由
	router := gin.Default()
	// 公共的Panic中间件
	router.Use(middleware.Recover)
	// API文档（示例文档）
	routers.DocsRouter(router)
	//// 开放路由
	//subRouter2.OpenRouter(router)
	//// 平台路由
	//subRouter2.PlatFormRouter(router)
	return router
}
