package platHander

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/plat/platServer"
)

// ListDictGroup godoc
// @id			 ListDictGroup字典分组列表
// @Summary      字典分组列表
// @Description  获取字典分组列表，改列表为内置列表
// @Router       /plat/dictGroup/list [post]
// @Tags         数据字典
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        Lang header string false "语言，不传默认为zh-CN"
// @Success      200 {object} model.ResBody{data=[]model.SelectRes} "响应成功"
func ListDictGroup(c *gin.Context) {
	// traceID 日志追踪
	traceID := c.GetString(constant.ContextTraceID)
	// 语言获取
	local := service.GetLocal(c)
	service.JsonRes(c, platServer.ListDictGroup(traceID, local))
}
