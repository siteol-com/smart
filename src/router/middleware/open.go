package middleware

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model"
	"siteol.com/smart/src/common/utils"
	"siteol.com/smart/src/service"
)

// OpenMiddleWare 开放中间件
func OpenMiddleWare(c *gin.Context) {
	// 生成请求唯一标志
	traceID := utils.TraceID()
	url := c.Request.URL.Path
	log.InfoTF(traceID, "Req URL = %s", url)
	c.Set(constant.ContextTraceID, traceID)
	// TODO 设置路由配置
	setRouter(c, url)
	// 设置语言
	setLang(c)
	// 是否是中间件拒绝
	middleRes := true
	defer func() {
		// TODO 中间件拒绝的特殊处理
		if middleRes {
			service.JsonRes(c, model.SysErr)
		}
	}()
	// 读取鉴权信息
	// TODO
	// 其他中间件或控制层响应，无需本层特殊处理
	middleRes = false
	c.Next()
}