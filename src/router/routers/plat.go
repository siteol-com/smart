package routers

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/router/middleware"
	"siteol.com/smart/src/service/plat/platHander"
)

// PlatRouter 平台业务路由
func PlatRouter(router *gin.Engine) {
	platRouter := router.Group("/plat", middleware.CommMiddleWare) // TODO 授权中间件
	{
		// 字典分组相关
		dictGroupRouter := platRouter.Group("/dictGroup")
		{
			dictGroupRouter.POST("/read", platHander.ReadDictGroup)
		}
		// 字典相关
		dictRouter := platRouter.Group("/dict")
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
		// 响应码相关
		responseRouter := platRouter.Group("/response")
		{
			responseRouter.POST("/nextVal", platHander.NextResponseVal)
			responseRouter.POST("/add", platHander.AddResponse)
			responseRouter.POST("/page", platHander.PageResponse)
			responseRouter.POST("/get", platHander.GetResponse)
			responseRouter.POST("/edit", platHander.EditResponse)
			responseRouter.POST("/del", platHander.DelResponse)
		}
		// 接口路由相关
		routerRouter := platRouter.Group("/router")
		{
			routerRouter.POST("/add", platHander.AddRouter)
			routerRouter.POST("/page", platHander.PageRouter)
			routerRouter.POST("/get", platHander.GetRouter)
			routerRouter.POST("/edit", platHander.EditRouter)
			routerRouter.POST("/del", platHander.DelRouter)
		}
		// 系统配置相关
		sysConfigRouter := platRouter.Group("/sysConfig")
		{
			sysConfigRouter.POST("/get", platHander.GetSysConfig)
			sysConfigRouter.POST("/edit", platHander.EditSysConfig)
		}
	}
}
