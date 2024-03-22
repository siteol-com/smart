package platHander

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/plat/platServer"
)

// GetSysConfig	godoc
// @id			GetSysConfig 系统配置详情
// @Summary		系统配置详情
// @Description	查询系统配置详情
// @Router		/plat/sysConfig/get [post]
// @Tags		数据系统配置
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq	true							"请求"
// @Success		200	{object}	baseModel.ResBody{data=model.SysConfigGetRes}	"响应成功"
func GetSysConfig(c *gin.Context) {
	// traceID 日志追踪
	traceID := c.GetString(constant.ContextTraceID)
	// 执行查询
	service.JsonRes(c, platServer.GetSysConfig(traceID))
}

// EditSysConfig	godoc
// @id				EditSysConfig 系统配置编辑
// @Summary			系统配置编辑
// @Description		在系统配置分组下编辑系统配置
// @Router			/plat/sysConfig/edit [post]
// @Tags			数据系统配置
// @Accept			json
// @Produce			json
// @Security		Token
// @Param			req	body		model.SysConfigEditReq	true	"请求"
// @Success			200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditSysConfig(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &model.SysConfigEditReq{})
	if err == nil {
		req := reqObj.(*model.SysConfigEditReq)
		// 执行创建
		service.JsonRes(c, platServer.EditSysConfig(traceID, req))
	}
}
