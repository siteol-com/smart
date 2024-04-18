package translator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"siteol.com/smart/src/common/model/baseModel"
	"strings"

	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model"
	"siteol.com/smart/src/common/redis"
)

// ReturnMsgTrans 执行响应码 => 响应文言 翻译
func ReturnMsgTrans(res *baseModel.ResBody, c *gin.Context, router *model.CacheRouter, traceID string) {
	// 语言类型
	lang := c.GetString(constant.ContextLang)
	// 非400校验错误执行翻译
	if res.HttpCode != http.StatusBadRequest {
		// 执行翻译
		tableMsgTrans(res, lang, traceID)
	}
	// 响应日志
	printBts := []byte("{}")
	if !router.ResPrint {
		printBts = []byte("{ Res Set Not Print}")
	} else {
		// 如需打印日志
		resBts, _ := json.Marshal(res)
		// JSON序列化
		printBts = resBts
	}
	log.InfoTF(traceID, "RespBody: %s", printBts)
}

// tableMsgTrans 执行Msg翻译
func tableMsgTrans(res *baseModel.ResBody, lang, traceID string) {
	// 读取缓存
	tranStr, err := redis.Get(constant.CacheResTrans)
	if err != nil {
		log.WarnTF(traceID, "GetTransLangCacheMap Fail . Err Is : %v", err)
		// 出错不翻译
		return

	}
	transMap := make(map[string]map[string]string)
	err = json.Unmarshal([]byte(tranStr), &transMap)
	if err != nil {
		log.ErrorTF(traceID, "JsonUnmarshal TransMap Fail . Err Is : %v", err)
		// 出错不翻译
		return
	}
	// 读取配置
	codeMap, ok := transMap[res.Code]
	// 未能匹配的响应码，复用基础文言
	if !ok {
		useCode := constant.Success
		if strings.HasPrefix(res.Code, "F") {
			useCode = constant.Fail
		}
		if strings.HasPrefix(res.Code, "E") {
			useCode = constant.Error
		}
		codeMap, ok = transMap[useCode]
	}
	// 执行语言读取
	if ok {
		langTemple, lok := codeMap[lang]
		if lok {
			// 检查是否有变量，有则进行变量替换
			if strings.Index(langTemple, "}}") > -1 {
				res.Msg = TableValReplace(langTemple, res.Data)
			} else {
				res.Msg = langTemple
			}
		}
	}
	// 无相关翻译
	return
}

// TableValReplace 执行变量替换
// 实际生效，{{name}}修改成功 => 米虫修改成功
func TableValReplace(temple string, data any) string {
	if data == nil {
		return temple
	}
	dataStr, err := json.Marshal(data)
	if err != nil {
		return temple
	}
	dataMap := make(map[string]any)
	err = json.Unmarshal([]byte(dataStr), &dataMap)
	if err != nil {
		return temple
	}
	// 提取模板中的变量数据
	valArray := getTempleVal(temple)
	// 提取参数并替换
	for _, val := range valArray {
		valObj, ok := dataMap[val]
		if ok && valObj != nil {
			// 变量存在值则替换
			temple = strings.ReplaceAll(temple, "{{"+val+"}}", fmt.Sprintf("%v", valObj))
		}
	}
	return temple
}

// 提取模板中的变量数据
func getTempleVal(temple string) []string {
	valArray := make([]string, 0)
	strS := strings.Split(temple, "{{")
	for _, i := range strS {
		if strings.Index(i, "}}") > -1 {
			valStr := i[:strings.Index(i, "}}")]
			valArray = append(valArray, valStr)
		}
	}
	return valArray
}
