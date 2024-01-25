package platServer

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model"
	"siteol.com/smart/src/common/mysql/platDb"
)

// ListDictGroup 字典分组列表
func ListDictGroup(traceID, local string) *model.ResBody {
	// 查询全部
	dbRes, err := platDb.DictGroupTable.FindAll()
	if err != nil {
		log.ErrorTF(traceID, "ListDictGroup Fail. Err : %s", err)
		return model.Fail(constant.DictGroupGetNG, nil)
	}
	// 处理响应
	resData := make([]*model.SelectRes, len(dbRes))
	// 循环并执行语言选择
	for i, d := range dbRes {
		var label string
		// 文言翻译
		switch local {
		case "en":
			label = d.NameEn
		default:
			label = d.Name
		}
		resData[i] = &model.SelectRes{
			Label: label,
			Value: d.Key,
		}
	}
	return model.SuccessUnPop(resData)
}
