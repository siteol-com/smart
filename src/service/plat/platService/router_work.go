package platService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/common/mysql/actuator"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/redis"
	"siteol.com/smart/src/common/utils"
	"strings"
)

// 业务层数据处理函数
// 抽取到独立文件中仅便于Server层阅读（没有特别意义）

// 解析数据库错误
func checkRouterDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "url_uni") {
			// 路由地址全局唯一
			return baseModel.Fail(constant.RouterUniUrlNG)
		}
		if strings.Contains(errStr, "name_uni") {
			// 路由地址全局唯一
			return baseModel.Fail(constant.RouterUniNameNG)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}

// 路由分页查询对象
func routerPageQuery(req *platModel.RouterPageReq) (query *actuator.Query) {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query = actuator.InitQuery()
	if req.Name != "" {
		query.Like("name", req.Name)
	}
	if req.Url != "" {
		query.Like("url", req.Url)
	}
	if req.ServiceCode != "" {
		query.Eq("service_code", req.ServiceCode)
	}
	if req.Type != "" {
		query.Eq("type", req.Type)
	}
	query.Eq("status", constant.StatusOpen)
	query.Desc("id")
	query.LimitByPage(req.Current, req.PageSize)
	return
}

// SyncRouterCache 同步响应码配置
func SyncRouterCache(traceID string) (err error) {
	allList, err := platDB.RouterTable.GetAll()
	if err != nil {
		log.ErrorTF(traceID, "SyncRouterCache GetRouter Fail . Err Is : %v", err)
		return
	}
	if len(allList) == 0 {
		log.WarnTF(traceID, "SyncRouterCache GetRouter Empty")
		return
	}
	// 组装缓存对象
	resCacheMap := make(map[string]*baseModel.CacheRouter, 0)
	for _, res := range allList {
		// 路由是物理删除
		cache := &baseModel.CacheRouter{
			Id:        res.Id,
			NeedAuth:  utils.StatusBool(res.Type),
			LogInDb:   utils.StatusBool(res.LogInDb),
			ReqPrint:  utils.StatusBool(res.ReqLogPrint),
			ReqSecure: utils.ArrayStr(res.ReqLogSecure),
			ResPrint:  utils.StatusBool(res.ResLogPrint),
			ResSecure: utils.ArrayStr(res.ResLogSecure),
		}
		resCacheMap[res.Url] = cache
	}
	// 写入缓存 无超期
	err = redis.Set(constant.CacheRouters, resCacheMap, 0)
	if err != nil {
		log.InfoTF(traceID, "SyncRouterCache Fail . Err Is : %v", err)
	}
	return
}
