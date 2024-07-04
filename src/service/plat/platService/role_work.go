package platService

import (
	"encoding/json"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/common/mysql/actuator"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/redis"
	"strings"
)

// 业务层数据处理函数
// 抽取到独立文件中仅便于Server层阅读（没有特别意义）

// 解析数据库错误
func checkRoleDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "name_uni") {
			// 角色名唯一
			return baseModel.Fail(constant.RoleNameUniNG)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}

// 字典分页查询对象
func rolePageQuery(req *platModel.RolePageReq) (query *actuator.Query) {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query = actuator.InitQuery()
	if req.Name != "" {
		query.Like("name", req.Name)
	}
	query.Eq("status", constant.StatusOpen)
	query.Desc("id")
	query.LimitByPage(req.Current, req.PageSize)
	return
}

// SyncRoleCache 同步全部角色
func SyncRoleCache(traceID string) (err error) {
	allList, err := platDB.RoleTable.GetAll()
	if err != nil {
		log.ErrorTF(traceID, "SyncRoleCache GetRole Fail . Err Is : %v", err)
		return
	}
	if len(allList) == 0 {
		log.WarnTF(traceID, "SyncRoleCache GetRole Empty")
		return
	}
	// 组装缓存对象
	resCacheMap := make(map[uint64]string, len(allList))
	for _, res := range allList {
		resCacheMap[res.Id] = res.Name
	}
	// 写入缓存 无超期
	err = redis.Set(constant.CacheRoles, resCacheMap, 0)
	if err != nil {
		log.InfoTF(traceID, "SyncRoleCache Fail . Err Is : %v", err)
	}
	return
}

// getRoleCache 获取角色缓存
func getRoleCache(traceID string) (res map[uint64]string, err error) {
	str, err := redis.Get(constant.CacheRoles)
	if err == nil {
		resCacheMap := make(map[uint64]string, 0)
		err = json.Unmarshal([]byte(str), &resCacheMap)
		if err == nil {
			res = resCacheMap
		} else {
			log.ErrorTF(traceID, "Unmarshal RouterCache Fail . Err Is : %v", err)
		}
	} else {
		log.ErrorTF(traceID, "Get RouterCache Fail . Err Is : %v", err)
	}
	return
}
