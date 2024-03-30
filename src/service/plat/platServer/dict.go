package platServer

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/actuator"
	"siteol.com/smart/src/common/mysql/platDb"
)

// ReadDict  读取字典
func ReadDict(traceID string, req *model.DictReadReq) *baseModel.ResBody {
	// 如果查询key不为空
	if len(req.GroupKeys) > 0 {
		dictListMap := make(map[string][]*baseModel.SelectRes, len(req.GroupKeys))
		dictValueMap := make(map[string]map[string]string, len(req.GroupKeys))
		// 遍历查询
		for _, groupKey := range req.GroupKeys {
			// 默认的Sort排序处理
			dictList, err := platDb.DictTable.FindByObjectSort(&platDb.Dict{GroupKey: groupKey})
			if err != nil {
				log.WarnTF(traceID, "ListDict Fail . GroupKey Query By : %s , Err is : %v", groupKey, err)
				dictListMap[groupKey] = make([]*baseModel.SelectRes, 0)
				dictValueMap[groupKey] = make(map[string]string, 0)
				continue
			}
			dictListMap[groupKey], dictValueMap[groupKey] = dictListToRead(dictList, req.Local)
		}
		return baseModel.SuccessUnPop(model.DictReadRes{List: dictListMap, Map: dictValueMap})
	}
	return baseModel.SuccessUnPop(nil)
}

// NextDictVal 字典的Val建议
func NextDictVal(traceID string, req *model.DictNextValReq) *baseModel.ResBody {
	// 组装Query
	query := actuator.InitQuery()
	query.Eq("group_key", req.GroupKey)
	total, err := platDb.DictTable.CountByQuery(query)
	if err != nil {
		log.ErrorTF(traceID, "NextDictVal Fail . Err Is : %v", err)
		return baseModel.Fail(constant.DictGetNG)
	}
	return baseModel.SuccessUnPop(total)
}

// AddDict 创建字典
func AddDict(traceID string, req *model.DictAddReq) *baseModel.ResBody {
	// 创建对象初始化
	dbReq := req.ToDbReq()
	err := platDb.DictTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddDict Fail . Err Is : %v", err)
		// 解析数据库错误
		return dictCheckDBErr(err)
	}
	return baseModel.Success(constant.DictAddSS, true)
}

// PageDict 查询字典分页
func PageDict(traceID string, req *model.DictPageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := platDb.DictTable.Page(dictPageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PageDict Fail . Err Is : %v", err)
		return baseModel.Fail(constant.DictGetNG)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(model.ToDictPageRes(list), total))
}

// GetDict 字典详情
func GetDict(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := platDb.DictTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetDict Fail . Err Is : %v", err)
		return baseModel.Fail(constant.DictGetNG)
	}
	return baseModel.SuccessUnPop(model.ToDictGetRes(&res))
}

// EditDict 编辑字典
func EditDict(traceID string, req *model.DictEditReq) *baseModel.ResBody {
	dbReq, err := platDb.DictTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetDict Fail . Err Is : %v", err)
		return baseModel.Fail(constant.DictGetNG)
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = platDb.DictTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditDict %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return dictCheckDBErr(err)
	}
	return baseModel.Success(constant.DictEditSS, true)
}

// BroDict 字典分组列表
func BroDict(traceID string, req *model.DictBroReq) *baseModel.ResBody {
	// 如果查询key不为空（只查询启用的数据）
	dictList, err := platDb.DictTable.FindByObjectSort(&platDb.Dict{GroupKey: req.GroupKey, Common: platDb.Common{Status: constant.StatusOpen}})
	if err != nil {
		log.WarnTF(traceID, "BroDict Fail . GroupKey Query By : %s , Err is : %v", req.GroupKey, err)
		return baseModel.Fail(constant.DictGetNG)
	}
	return baseModel.SuccessUnPop(dictListToBro(dictList, req.Local))
}

// SortDict 字典排序处理
func SortDict(traceID string, req []*baseModel.SortReq) *baseModel.ResBody {
	if len(req) == 0 {
		return baseModel.ResFail
	}
	err := platDb.DictTable.SortWithTransaction(req)
	if err != nil {
		log.ErrorTF(traceID, "SortDict Fail .  Err is : %s", err)
		return baseModel.Fail(constant.DictSortNG)
	}
	return baseModel.Success(constant.DictSortSS, true)
}

// DelDict 字典封存
func DelDict(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := platDb.DictTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetDict Fail . Err Is : %v", err)
		return baseModel.Fail(constant.DictGetNG)
	}
	// 字典禁止刪除
	if dbReq.Mark == constant.StatusLock {
		log.ErrorTF(traceID, "EditDel %d Fail . Can not Edit", dbReq.Id)
		return baseModel.Fail(constant.DictMarkNG)
	}
	dbReq.Status = constant.StatusClose
	err = platDb.DictTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "DelDict %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return dictCheckDBErr(err)
	}
	return baseModel.Success(constant.DictDelSS, true)
}
