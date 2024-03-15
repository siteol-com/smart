package platServer

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/actuator"
	"siteol.com/smart/src/common/mysql/platDb"
	"strings"
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
		return baseModel.Fail(constant.DictGetNG, nil)
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
		return checkDictDBErr(err)
	}
	return baseModel.Success(constant.DictAddSS, true)
}

// PageDict 查询字典分页
func PageDict(traceID string, req *model.DictPageReq) *baseModel.ResBody {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query := actuator.InitQuery()
	if req.GroupKey != "" {
		query.Eq("group_key", req.GroupKey)
		query.Asc("sort")
	} else {
		query.Desc("id")
	}
	if req.Label != "" {
		var label string
		switch req.Local {
		case "en":
			label = "label_en"
		default:
			label = "label"
		}
		query.Like(label, req.Label)
	}
	query.Eq("status", constant.StatusOpen)
	query.LimitByPage(req.Current, req.PageSize)
	// 查询分页
	total, list, err := platDb.DictTable.Page(query)
	if err != nil {
		log.ErrorTF(traceID, "PageDict Fail . Err Is : %v", err)
		return baseModel.Fail(constant.DictGetNG, nil)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(model.ToDictPageRes(list), total))
}

// GetDict 字典详情
func GetDict(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := platDb.DictTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetDict Fail . Err Is : %v", err)
		return baseModel.Fail(constant.DictGetNG, nil)
	}
	return baseModel.SuccessUnPop(res)
}

// EditDict 编辑字典
func EditDict(traceID string, req *model.DictEditReq) *baseModel.ResBody {
	dbReq, err := platDb.DictTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetDict Fail . Err Is : %v", err)
		return baseModel.Fail(constant.DictGetNG, nil)
	}
	// 字段禁止编辑
	if dbReq.Mark == constant.StatusLock {
		log.ErrorTF(traceID, "EditDict %d Fail . Can not Edit", dbReq.Id)
		return baseModel.Fail(constant.DictMarkNG, nil)
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = platDb.DictTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditDict %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkDictDBErr(err)
	}
	return baseModel.Success(constant.DictEditSS, true)
}

// BroDict 字典分组列表
func BroDict(traceID string, req *model.DictBroReq) *baseModel.ResBody {
	// 如果查询key不为空
	dictList, err := platDb.DictTable.FindByObjectSort(&platDb.Dict{GroupKey: req.GroupKey})
	if err != nil {
		log.WarnTF(traceID, "BroDict Fail . GroupKey Query By : %s , Err is : %v", req.GroupKey, err)
		return baseModel.Fail(constant.DictGetNG, nil)
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
		return baseModel.Fail(constant.DictSortNG, nil)
	}
	return baseModel.Success(constant.DictSortSS, true)
}

// DelDict 字典封存
func DelDict(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := platDb.DictTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetDict Fail . Err Is : %v", err)
		return baseModel.Fail(constant.DictGetNG, nil)
	}
	// 字段禁止编辑
	if dbReq.Mark == constant.StatusLock {
		log.ErrorTF(traceID, "EditDel %d Fail . Can not Edit", dbReq.Id)
		return baseModel.Fail(constant.DictMarkNG, nil)
	}
	dbReq.Status = constant.StatusClose
	err = platDb.DictTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "DelDict %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkDictDBErr(err)
	}
	return baseModel.Success(constant.DictDelSS, true)
}

// 数据库查询列表转换选择列表与翻译列表
func dictListToBro(list []*platDb.Dict, local string) []*baseModel.SortRes {
	broList := make([]*baseModel.SortRes, 0)
	for _, dict := range list {
		// 未启用字典
		if dict.Status != constant.StatusOpen {
			continue
		}
		// 如果可选择，创建选择数据
		if dict.Choose == constant.StatusOpen {
			// 翻译label
			var label string
			switch local {
			case "en":
				label = dict.LabelEn
			default:
				label = dict.Label
			}
			broList = append(broList, &baseModel.SortRes{
				Id:   dict.Id,
				Name: label,
				Sort: dict.Sort,
			})
		}
	}
	return broList
}

// 数据库查询列表转换选择列表与翻译列表
func dictListToRead(list []*platDb.Dict, local string) ([]*baseModel.SelectRes, map[string]string) {
	labelList := make([]*baseModel.SelectRes, 0)
	valueMap := make(map[string]string, 0)
	for _, dict := range list {
		// 翻译label
		var label string
		switch local {
		case "en":
			label = dict.LabelEn
		default:
			label = dict.Label
		}
		// 创建Map数据
		valueMap[dict.Val] = label
		// 未启用字典（可以翻译，不可以选择）
		if dict.Status != constant.StatusOpen {
			continue
		}
		// 如果可选择，创建选择数据
		if dict.Choose == constant.StatusOpen {
			labelList = append(labelList, &baseModel.SelectRes{Label: label, Value: dict.Val})
		}
	}
	return labelList, valueMap
}

// 解析数据库错误
func checkDictDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "dict_uni") {
			// 字典分组下字典值唯一
			return baseModel.Fail(constant.DictUniNG, nil)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}
