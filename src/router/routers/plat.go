package routers

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/router/middleware"
	"siteol.com/smart/src/service/plat/platHandler"
)

// PlatRouter 平台业务路由
func PlatRouter(router *gin.Engine) {
	platRouter := router.Group("/plat", middleware.CommMiddleWare) // TODO 授权中间件
	{
		// 字典相关
		dictRouter := platRouter.Group("/dict")
		{
			dictRouter.POST("/read", platHandler.ReadDict)
			dictRouter.POST("/nextVal", platHandler.NextDictVal)
			dictRouter.POST("/add", platHandler.AddDict)
			dictRouter.POST("/page", platHandler.PageDict)
			dictRouter.POST("/get", platHandler.GetDict)
			dictRouter.POST("/edit", platHandler.EditDict)
			dictRouter.POST("/bro", platHandler.BroDict)
			dictRouter.POST("/sort", platHandler.SortDict)
			dictRouter.POST("/del", platHandler.DelDict)
		}
		// 响应码相关
		responseRouter := platRouter.Group("/response")
		{
			responseRouter.POST("/nextVal", platHandler.NextResponseVal)
			responseRouter.POST("/add", platHandler.AddResponse)
			responseRouter.POST("/page", platHandler.PageResponse)
			responseRouter.POST("/get", platHandler.GetResponse)
			responseRouter.POST("/edit", platHandler.EditResponse)
			responseRouter.POST("/del", platHandler.DelResponse)
		}
		// 接口路由相关
		routerRouter := platRouter.Group("/router")
		{
			routerRouter.POST("/add", platHandler.AddRouter)
			routerRouter.POST("/page", platHandler.PageRouter)
			routerRouter.POST("/get", platHandler.GetRouter)
			routerRouter.POST("/edit", platHandler.EditRouter)
			routerRouter.POST("/del", platHandler.DelRouter)
		}
		// 权限相关
		permissionRouter := platRouter.Group("/permission")
		{
			permissionRouter.POST("/add", platHandler.AddPermission)
			permissionRouter.POST("/tree", platHandler.TreePermission)
			permissionRouter.POST("/get", platHandler.GetPermission)
			permissionRouter.POST("/edit", platHandler.EditPermission)
			permissionRouter.POST("/del", platHandler.DelPermission)
			permissionRouter.POST("/bro", platHandler.BroPermission)
			permissionRouter.POST("/sort", platHandler.SortPermission)
		}
		// 角色相关
		roleRouter := platRouter.Group("/role")
		{
			roleRouter.POST("/add", platHandler.AddRole)
			roleRouter.POST("/page", platHandler.PageRole)
			roleRouter.POST("/get", platHandler.GetRole)
			roleRouter.POST("/edit", platHandler.EditRole)
			roleRouter.POST("/del", platHandler.DelRole)
		}
		// 系统配置相关
		sysConfigRouter := platRouter.Group("/sysConfig")
		{
			sysConfigRouter.POST("/get", platHandler.GetSysConfig)
			sysConfigRouter.POST("/edit", platHandler.EditSysConfig)
		}
	}
}
