package middleware

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model"
)

// 设置语言上下文
func setLang(c *gin.Context) {
	lang := c.GetHeader(constant.ContextLang)
	if lang == "" || lang == "null" {
		lang = "zh-CN"
	}
	c.Set(constant.ContextLang, lang)
}

// 设置路由上下文
func setRouter(c *gin.Context, url string) {
	// TODO 缓存处理
	router := model.CacheRouterNormal
	c.Set(constant.ContextRouterC, router)
}
