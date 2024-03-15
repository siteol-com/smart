package platHander

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/model"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/plat/platServer"
)

// ReadDict      godoc
//
//	@id				ReadDict 读取字典
//	@Summary		读取字典
//	@Description	获取字典下拉列表以及关联键值Map
//	@Router			/plat/dict/read [post]
//	@Tags			数据字典
//	@Accept			json
//	@Produce		json
//	@Security		Token
//	@Param			Lang	header		string										false	"语言，不传默认为zh-CN"
//	@Param			req		body		model.DictReadReq							true	"请求"
//	@Success		200		{object}	baseModel.ResBody{data=model.DictReadRes}	"响应成功"
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

// NextDictVal   godoc
//
//	@id				NextDictVal 字典NextVal建议
//	@Summary		字典NextVal建议
//	@Description	在字典分组下读取下一个Val的建议值
//	@Router			/plat/dict/nextVal [post]
//	@Tags			数据字典
//	@Accept			json
//	@Produce		json
//	@Security		Token
//	@Param			req	body		model.DictNextValReq			true	"请求"
//	@Success		200	{object}	baseModel.ResBody{data=int64}	"响应成功"
func NextDictVal(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &model.DictNextValReq{})
	if err == nil {
		req := reqObj.(*model.DictNextValReq)
		// 执行创建
		service.JsonRes(c, platServer.NextDictVal(traceID, req))
	}
}

// AddDict godoc
//
//	@id				AddDict 字典新建
//	@Summary		字典新建
//	@Description	在字典分组下新建字典
//	@Router			/plat/dict/add [post]
//	@Tags			数据字典
//	@Accept			json
//	@Produce		json
//	@Security		Token
//	@Param			req	body		model.DictAddReq				true	"请求"
//	@Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func AddDict(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &model.DictAddReq{})
	if err == nil {
		req := reqObj.(*model.DictAddReq)
		// 执行创建
		service.JsonRes(c, platServer.AddDict(traceID, req))
	}
}

// PageDict godoc
//
//	@id				PageDict 字典分页
//	@Summary		字典分页
//	@Description	查询字典分页数据
//	@Router			/plat/dict/page [post]
//	@Tags			数据字典
//	@Accept			json
//	@Produce		json
//	@Security		Token
//	@Param			req	body		model.DictPageReq													true	"请求"
//	@Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]model.DictPageRes}}	"响应成功"
func PageDict(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &model.DictPageReq{})
	if err == nil {
		req := reqObj.(*model.DictPageReq)
		// 语言读取
		req.Local = service.GetLocal(c)
		// 执行查询
		service.JsonRes(c, platServer.PageDict(traceID, req))
	}
}

// GetDict godoc
//
//	@id				GetDict 字典详情
//	@Summary		字典详情
//	@Description	查询字典详情
//	@Router			/plat/dict/get [post]
//	@Tags			数据字典
//	@Accept			json
//	@Produce		json
//	@Security		Token
//	@Param			req	body		baseModel.IdReq												true	"请求"
//	@Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=platDb.Dict}}	"响应成功"
func GetDict(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, platServer.GetDict(traceID, req))
	}
}

// EditDict godoc
//
//	@id				EditDict 字典编辑
//	@Summary		字典编辑
//	@Description	在字典分组下编辑字典
//	@Router			/plat/dict/edit [post]
//	@Tags			数据字典
//	@Accept			json
//	@Produce		json
//	@Security		Token
//	@Param			req	body		model.DictEditReq				true	"请求"
//	@Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditDict(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &model.DictEditReq{})
	if err == nil {
		req := reqObj.(*model.DictEditReq)
		// 执行创建
		service.JsonRes(c, platServer.EditDict(traceID, req))
	}
}

// BroDict godoc
//
//	@id				BroDict 字典排序数据
//	@Summary		排序数据
//	@Description	获取字典排序数据
//	@Router			/plat/dict/bro [post]
//	@Tags			数据字典
//	@Accept			json
//	@Produce		json
//	@Security		Token
//	@Param			Lang	header		string										false	"语言，不传默认为zh-CN"
//	@Param			req		body		model.DictBroReq							true	"请求"
//	@Success		200		{object}	baseModel.ResBody{data=baseModel.SortRes}	"响应成功"
func BroDict(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &model.DictBroReq{})
	if err == nil {
		req := reqObj.(*model.DictBroReq)
		// 语言读取
		req.Local = service.GetLocal(c)
		// 执行查询
		service.JsonRes(c, platServer.BroDict(traceID, req))
	}
}

// SortDict godoc
//
//	@id				SortDict 字典排序处理
//	@Summary		排序处理
//	@Description	字典分组下字典排序处理
//	@Router			/plat/dict/sort [post]
//	@Tags			数据字典
//	@Accept			json
//	@Produce		json
//	@Security		Token
//	@Param			Lang	header		string							false	"语言，不传默认为zh-CN"
//	@Param			req		body		[]baseModel.SortReq				true	"请求"
//	@Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func SortDict(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &[]*baseModel.SortReq{})
	if err == nil {
		req := reqObj.(*[]*baseModel.SortReq)
		// 执行查询
		service.JsonRes(c, platServer.SortDict(traceID, *req))
	}
}

// DelDict godoc
//
//	@id				DelDict 字典封存
//	@Summary		字典封存
//	@Description	字典封存处理
//	@Router			/plat/dict/del [post]
//	@Tags			数据字典
//	@Accept			json
//	@Produce		json
//	@Security		Token
//	@Param			Lang	header		string							false	"语言，不传默认为zh-CN"
//	@Param			req		body		[]baseModel.IdReq				true	"请求"
//	@Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelDict(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, platServer.DelDict(traceID, req))
	}
}
