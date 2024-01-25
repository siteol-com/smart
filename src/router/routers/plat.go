package routers

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/router/middleware"
	"siteol.com/smart/src/service/plat/platHander"
)

// PlatRouter 平台业务路由
func PlatRouter(router *gin.Engine) {
	platRouter := router.Group("/plat") // TODO 授权中间件
	{
		// 字典分组相关
		dictGroupRouter := platRouter.Group("/dictGroup", middleware.OpenMiddleWare)
		{
			dictGroupRouter.POST("/list", platHander.ListDictGroup)
		}

		// 字典相关
		dictRouter := platRouter.Group("/dict", middleware.OpenMiddleWare)
		{
			dictRouter.POST("/read", platHander.ReadDict)
		}
	}
}
