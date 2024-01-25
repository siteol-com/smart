package platHander

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/model"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/plat/platServer"
)

// ReadDict godoc
// @id			 ReadDict读取字典
// @Summary      读取字典
// @Description  获取字典下拉列表以及关联键值Map
// @Router       /plat/dict/read [post]
// @Tags         数据字典
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        Lang header string false "语言，不传默认为zh-CN"
// @Param        req body model.DictReadReq true "请求"
// @Success      200 {object} model.ResBody{data=model.DictReadRes} "响应成功"
func ReadDict(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &model.DictReadReq{})
	if err == nil {
		req := reqObj.(*model.DictReadReq)
		// 语言读取
		req.Local = service.GetLocal(c)
		// 执行查询
		service.JsonRes(c, platServer.ReadDict(traceID, req))
	}
}
