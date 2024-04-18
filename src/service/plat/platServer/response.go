package platServer

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/common/mysql/platDb"
)

// NextResponseVal 响应码的Val建议
func NextResponseVal(traceID string, req *platModel.ResponseNextValReq) *baseModel.ResBody {
	responseCode, err := responseMakeCode(req.ServiceCode, req.Type)
	if err != nil {
		log.ErrorTF(traceID, "NextResponseVal Fail . Err Is : %v", err)
		return baseModel.Fail(constant.ResponseGetNG)
	}
	return baseModel.SuccessUnPop(responseCode)
}

// AddResponse 创建响应码
func AddResponse(traceID string, req *platModel.ResponseAddReq) *baseModel.ResBody {
	// 创建对象初始化
	dbReq := req.ToDbReq()
	// 计算响应码
	responseCode, err := responseMakeCode(req.ServiceCode, req.Type)
	if err != nil {
		log.ErrorTF(traceID, "NextResponseVal Fail . Err Is : %v", err)
		return baseModel.Fail(constant.ResponseGetNG)
	}
	dbReq.Code = responseCode
	err = platDb.ResponseTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddResponseFail . Err Is : %v", err)
		// 解析数据库错误
		return checkResponseDBErr(err)
	}
	// 如果插入Code和响应Code不同，提示一下
	if req.Code != responseCode {
		// 响应码创建成功,实际响应码为{{code}}
		return baseModel.Success(constant.ResponseAddSSWNC, struct {
			Code string `json:"code"`
		}{Code: responseCode})
	}
	// 异步更新缓存
	go func() { _ = SyncResponseCache(traceID) }()
	return baseModel.Success(constant.ResponseAddSS, true)
}

// PageResponse 查询响应码分页
func PageResponse(traceID string, req *platModel.ResponsePageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := platDb.ResponseTable.Page(responsePageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PageResponseFail . Err Is : %v", err)
		return baseModel.Fail(constant.ResponseGetNG)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(platModel.ToResponsePageRes(list), total))
}

// GetResponse 响应码详情
func GetResponse(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := platDb.ResponseTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetResponseFail . Err Is : %v", err)
		return baseModel.Fail(constant.ResponseGetNG)
	}
	return baseModel.SuccessUnPop(platModel.ToResponseGetRes(&res))
}

// EditResponse 编辑响应码
func EditResponse(traceID string, req *platModel.ResponseEditReq) *baseModel.ResBody {
	dbReq, err := platDb.ResponseTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetResponseFail . Err Is : %v", err)
		return baseModel.Fail(constant.ResponseGetNG)
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = platDb.ResponseTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditResponse %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkResponseDBErr(err)
	}
	// 异步更新缓存
	go func() { _ = SyncResponseCache(traceID) }()
	return baseModel.Success(constant.ResponseEditSS, true)
}

// DelResponse 响应码封存
func DelResponse(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := platDb.ResponseTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetResponseFail . Err Is : %v", err)
		return baseModel.Fail(constant.ResponseGetNG)
	}
	// 响应码禁止刪除
	if dbReq.Mark == constant.StatusLock {
		log.ErrorTF(traceID, "EditDel %d Fail . Can not Edit", dbReq.Id)
		return baseModel.Fail(constant.ResponseMarkNG)
	}
	dbReq.Status = constant.StatusClose
	err = platDb.ResponseTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "DelResponse %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkResponseDBErr(err)
	}
	// 异步更新缓存
	go func() { _ = SyncResponseCache(traceID) }()
	return baseModel.Success(constant.ResponseDelSS, true)
}
