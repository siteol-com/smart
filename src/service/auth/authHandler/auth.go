package authHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/auth/authService"
)

// AuthLogin 	godoc
// @id			AuthLogin 账号登陆
// @Summary		账号密码登陆
// @Description	账号密码登陆
// @Router		/auth/login [post]
// @Tags		开放接口
// @Accept		json
// @Produce		json
// @Param		req	body		baseModel.AccountLoginReq		true				"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.AccountLoginRes}	"响应成功"
func AuthLogin(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.AccountLoginReq{})
	if err == nil {
		req := reqObj.(*baseModel.AccountLoginReq)
		// 执行创建
		service.JsonRes(c, authService.AuthLogin(traceID, req))
	}
}

// AuthReset 	godoc
// @id			AuthReset 账号密码重置
// @Summary		账号密码重置
// @Description	账号密码重置
// @Router		/auth/reset [post]
// @Tags		开放接口
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.AccountResetReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}		"响应成功"
func AuthReset(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.AccountResetReq{})
	if err == nil {
		req := reqObj.(*baseModel.AccountResetReq)
		// 读取当前登陆用户
		authUser := service.GetAuthUser(c)
		// 执行创建
		service.JsonRes(c, authService.AuthReset(traceID, req, authUser))
	}
}
