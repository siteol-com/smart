package model

// DictReadReq 字典读取请求
type DictReadReq struct {
	GroupKeys []string `json:"groupKeys" binding:"required" example:"serviceCode,responseType"` // 需要查询的字典分组
	Local     string   `json:"-"`                                                               // 字典语言
}

// DictReadRes 字典读取响应
type DictReadRes struct {
	List map[string][]*SelectRes      `json:"list"` // 字典下拉列表 {'serviceCode':"[{'label':'基础','value':'1'}]"}
	Map  map[string]map[string]string `json:"map"`  // 字典翻译Map {'serviceCode':{'1':'基础'}}
}
