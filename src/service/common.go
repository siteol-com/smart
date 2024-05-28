package service

import (
	"errors"
	"siteol.com/smart/src/common/model/baseModel"
	"strings"

	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/utils/translator"
	"siteol.com/smart/src/common/validate"
)

// ValidateReqObj 读取并验证请求数据（并处理响应）
func ValidateReqObj(c *gin.Context, req any) (traceID string, reqObj any, err error) {
	// traceID 日志追踪
	traceID = c.GetString(constant.ContextTraceID)
	// 校验并且 解析请求数据
	res, reqObj := validate.Readable(c, req)
	if res != nil {
		err = errors.New(res.Msg)
		// 处理响应
		JsonRes(c, res)
	}
	return
}

// GetLocal 从上下文获取请求的语言
func GetLocal(c *gin.Context) string {
	local := c.GetString(constant.ContextLang)
	lineIndex := strings.Index(local, "-")
	if lineIndex > 0 {
		local = local[:lineIndex]
	}
	return local
}

// GetRouterConf 从上下文获取登录用户授权机构体
func GetRouterConf(c *gin.Context) *baseModel.CacheRouter {
	obj, ok := c.Get(constant.ContextRouterC)
	if ok {
		router := &baseModel.CacheRouter{}
		router = obj.(*baseModel.CacheRouter)
		return router
	}
	// 空白对象
	return baseModel.CacheRouterNormal
}

// JsonRes 执行Json响应 包含响应日志处理
func JsonRes(c *gin.Context, res *baseModel.ResBody) {
	// traceID 日志追踪
	traceID := c.GetString(constant.ContextTraceID)
	// 获取路由配置
	router := GetRouterConf(c)
	// 对Res进行翻译
	translator.ReturnMsgTrans(res, c, router, traceID)
	c.JSON(res.HttpCode, res)
}
