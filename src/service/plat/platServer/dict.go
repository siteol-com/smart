package platServer

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model"
	"siteol.com/smart/src/common/mysql/platDb"
)

// ReadDict  读取字典
func ReadDict(traceID string, req *model.DictReadReq) *model.ResBody {
	// 如果查询key不为空
	if len(req.GroupKeys) > 0 {
		dictListMap := make(map[string][]*model.SelectRes, len(req.GroupKeys))
		dictValueMap := make(map[string]map[string]string, len(req.GroupKeys))
		// 遍历查询
		for _, groupKey := range req.GroupKeys {
			// 默认的Sort排序处理
			dictList, err := platDb.DictTable.FindByObjectSort(&platDb.Dict{GroupKey: groupKey})
			if err != nil {
				log.WarnTF(traceID, "ListDict Fail . GroupKey Query By : %s , Err is : %v", groupKey, err)
				dictListMap[groupKey] = make([]*model.SelectRes, 0)
				dictValueMap[groupKey] = make(map[string]string, 0)
				continue
			}
			dictListMap[groupKey], dictValueMap[groupKey] = dictListToRead(dictList, req.Local)
		}
		return model.SuccessUnPop(model.DictReadRes{List: dictListMap, Map: dictValueMap})
	}
	return model.SuccessUnPop(nil)
}

// 数据库查询列表转换选择列表与翻译列表
func dictListToRead(list []*platDb.Dict, local string) ([]*model.SelectRes, map[string]string) {
	labelList := make([]*model.SelectRes, 0)
	valueMap := make(map[string]string, 0)
	for _, dict := range list {
		// 未启用字典
		if dict.Status != constant.StatusOpen {
			continue
		}
		// 翻译label
		var label string
		switch local {
		case "en":
			label = dict.LabelEn
		default:
			label = dict.Label
		}
		// 如果可选择，创建选择数据
		if dict.Choose == constant.StatusOpen {
			labelList = append(labelList, &model.SelectRes{Label: label, Value: dict.Val})
		}
		// 创建Map数据
		valueMap[dict.Val] = label
	}
	return labelList, valueMap
}
