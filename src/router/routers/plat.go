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
			dictGroupRouter.POST("/read", platHander.ReadDictGroup)
		}
		// 字典相关
		dictRouter := platRouter.Group("/dict", middleware.OpenMiddleWare)
		{
			dictRouter.POST("/read", platHander.ReadDict)
			dictRouter.POST("/nextVal", platHander.NextDictVal)
			dictRouter.POST("/add", platHander.AddDict)
			dictRouter.POST("/page", platHander.PageDict)
			dictRouter.POST("/get", platHander.GetDict)
			dictRouter.POST("/edit", platHander.EditDict)
			dictRouter.POST("/bro", platHander.BroDict)
			dictRouter.POST("/sort", platHander.SortDict)
			dictRouter.POST("/del", platHander.DelDict)
		}
		// 系统配置相关
		sysConfigRouter := platRouter.Group("/sysConfig", middleware.OpenMiddleWare)
		{
			sysConfigRouter.POST("/get", platHander.GetSysConfig)
			sysConfigRouter.POST("/edit", platHander.EditSysConfig)
		}
	}
}
