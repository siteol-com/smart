package middleware

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/utils"
	"siteol.com/smart/src/service"
)

// CommMiddleWare 开放中间件
func CommMiddleWare(c *gin.Context) {
	// 生成请求唯一标志
	traceID := utils.TraceID()
	url := c.Request.URL.Path
	log.InfoTF(traceID, "Req URL = %s", url)
	c.Set(constant.ContextTraceID, traceID)
	// 设置语言
	setLang(c)
	// 读取路由配置，可能会NG 404 路由找不到
	router, routerNg := setRouter(c, url, traceID)
	if routerNg {
		service.JsonRes(c, baseModel.PathErr)
		return
	}
	// 请求日志处理，可能会NG 500 系统异常
	reqNg := setReq(c, router, url, traceID)
	if reqNg {
		service.JsonRes(c, baseModel.SysErr)
		return
	}
	// 读取鉴权信息
	if router.NeedAuth {
		// TODO
	}
	c.Next()
}
