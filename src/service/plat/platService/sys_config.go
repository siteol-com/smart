package platService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/redis"
)

// GetSysConfig 系统配置详情
func GetSysConfig(traceID string) *baseModel.ResBody {
	res, err := platDB.SysConfigTable.FindOneById(1)
	if err != nil {
		log.ErrorTF(traceID, "GetSysConfig Fail . Err Is : %v", err)
		return baseModel.Fail(constant.SysConfigGetNG)
	}
	return baseModel.SuccessUnPop(platModel.ToSysConfigGetRes(&res))
}

// EditSysConfig 编辑系统配置
func EditSysConfig(traceID string, req *platModel.SysConfigEditReq) *baseModel.ResBody {
	dbReq, err := platDB.SysConfigTable.FindOneById(1)
	if err != nil {
		log.ErrorTF(traceID, "GetSysConfig Fail . Err Is : %v", err)
		return baseModel.Fail(constant.SysConfigGetNG)
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = platDB.SysConfigTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditSysConfig %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误 - 不涉及
		return baseModel.ResFail
	}
	// 异步更新缓存
	go func() { _ = SyncSysConfigCache(traceID) }()
	return baseModel.Success(constant.SysConfigEditSS, true)
}

// SyncSysConfigCache 同步系统配置
func SyncSysConfigCache(traceID string) (err error) {
	db, err := platDB.SysConfigTable.FindOneById(1)
	if err != nil {
		log.ErrorTF(traceID, "GetSysConfig Fail . Err Is : %v", err)
		return
	}
	err = redis.Set(constant.CacheSysConf, db, 0)
	if err != nil {
		log.ErrorTF(traceID, "SetSysConfigCache Fail . Err Is : %v", err)
	}
	return
}
