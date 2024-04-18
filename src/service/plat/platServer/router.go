package platServer

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/common/mysql/platDb"
)

// AddRouter 创建路由接口
func AddRouter(traceID string, req *platModel.RouterAddReq) *baseModel.ResBody {
	// 创建对象初始化
	dbReq := req.ToDbReq()
	err := platDb.RouterTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddRouterFail . Err Is : %v", err)
		// 解析数据库错误
		return checkRouterDBErr(err)
	}
	// 异步更新缓存
	go func() { _ = SyncRouterCache(traceID) }()
	return baseModel.Success(constant.RouterAddSS, true)
}

// PageRouter 查询路由接口分页
func PageRouter(traceID string, req *platModel.RouterPageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := platDb.RouterTable.Page(routerPageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PageRouterFail . Err Is : %v", err)
		return baseModel.Fail(constant.RouterGetNG)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(platModel.ToRouterPageRes(list), total))
}

// GetRouter 路由接口详情
func GetRouter(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := platDb.RouterTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetRouterFail . Err Is : %v", err)
		return baseModel.Fail(constant.RouterGetNG)
	}
	return baseModel.SuccessUnPop(platModel.ToRouterGetRes(&res))
}

// EditRouter 编辑路由接口
func EditRouter(traceID string, req *platModel.RouterEditReq) *baseModel.ResBody {
	dbReq, err := platDb.RouterTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetRouterFail . Err Is : %v", err)
		return baseModel.Fail(constant.RouterGetNG)
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = platDb.RouterTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditRouter %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkRouterDBErr(err)
	}
	// 异步更新缓存
	go func() { _ = SyncRouterCache(traceID) }()
	return baseModel.Success(constant.RouterEditSS, true)
}

// DelRouter 路由接口封存
func DelRouter(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := platDb.RouterTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetRouterFail . Err Is : %v", err)
		return baseModel.Fail(constant.RouterGetNG)
	}
	// 路由接口禁止刪除
	if dbReq.Mark == constant.StatusLock {
		log.ErrorTF(traceID, "EditDel %d Fail . Can not Edit", dbReq.Id)
		return baseModel.Fail(constant.RouterMarkNG)
	}
	// 路由是物理删除
	err = platDb.RouterTable.DeleteOne(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelRouter %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkRouterDBErr(err)
	}
	// 异步更新缓存
	go func() { _ = SyncRouterCache(traceID) }()
	return baseModel.Success(constant.RouterDelSS, true)
}
