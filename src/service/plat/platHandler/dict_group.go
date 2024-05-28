package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/plat/platService"
)

// ReadDictGroup	godoc
// @id				ReadDictGroup 字典分组读取
// @Summary			字典分组读取
// @Description		获取字典分组下拉列表以及关联键值Map
// @Router			/plat/dictGroup/read [post]
// @Tags			数据字典
// @Accept			json
// @Produce			json
// @Security		Token
// @Param			Lang	header		string	false										"语言，不传默认为zh-CN"
// @Success			200		{object}	baseModel.ResBody{data=platModel.DictGroupReadRes}	"响应成功"
func ReadDictGroup(c *gin.Context) {
	// traceID 日志追踪
	traceID := c.GetString(constant.ContextTraceID)
	// 语言获取
	local := service.GetLocal(c)
	service.JsonRes(c, platService.ReadDictGroup(traceID, local))
}
