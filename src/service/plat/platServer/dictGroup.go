package platServer

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/platDb"
)

// ReadDictGroup 字典分组读取
func ReadDictGroup(traceID, local string) *baseModel.ResBody {
	// 查询全部
	dbRes, err := platDb.DictGroupTable.FindAll()
	if err != nil {
		log.ErrorTF(traceID, "ListDictGroup Fail. Err : %s", err)
		return baseModel.Fail(constant.DictGroupGetNG)
	}
	dictGroupList := make([]*baseModel.SelectRes, 0)
	dictGroupMap := make(map[string]string, 0)
	// 循环并执行语言选择
	for _, d := range dbRes {
		// 未启用
		if d.Status != constant.StatusOpen {
			continue
		}
		var label string
		// 文言翻译
		switch local {
		case "en":
			label = d.NameEn
		default:
			label = d.Name
		}
		dictGroupList = append(dictGroupList, &baseModel.SelectRes{Label: label, Value: d.Key})
		dictGroupMap[d.Key] = label
	}
	// 处理响应
	resData := &model.DictGroupReadRes{
		List: dictGroupList,
		Map:  dictGroupMap,
	}
	return baseModel.SuccessUnPop(resData)
}
