package model

import "siteol.com/smart/src/common/model/baseModel"

// DictGroupReadRes 字典分组读取响应
type DictGroupReadRes struct {
	List []*baseModel.SelectRes `json:"list"`                                            // 字典分组下拉列表 [{'label':'业务模块','value':'serviceCode'}]
	Map  map[string]string      `json:"map" example:"{'serviceCode':'Business module'}"` // 字典分组翻译Map
}

// DictReadReq 字典读取请求
type DictReadReq struct {
	GroupKeys []string `json:"groupKeys" binding:"required" example:"serviceCode,responseType"` // 需要查询的字典分组
	Local     string   `json:"-"`                                                               // 字典语言
}

// DictReadRes 字典读取响应
type DictReadRes struct {
	List map[string][]*baseModel.SelectRes `json:"list"` // 字典下拉列表 {'serviceCode':"[{'label':'基础','value':'1'}]"}
	Map  map[string]map[string]string      `json:"map"`  // 字典翻译Map {'serviceCode':{'1':'基础'}}
}

// DictNextValReq 字典建议Val请求
type DictNextValReq struct {
	GroupKey string `json:"groupKey" binding:"required" example:"serviceCode"` // 字典分组KEY
}

// DictBroReq 字典分组列表请求
type DictBroReq struct {
	GroupKey string `json:"groupKey" binding:"required" example:"serviceCode"` // 字典分组KEY
	Local    string `json:"-"`                                                 // 字典语言
}

// DictDoReq 字典处理请求（共通）
type DictDoReq struct {
	Label   string `json:"label" binding:"required" example:"PlatBase"`         // 字段名称
	LabelEn string `json:"labelEn" binding:"required" example:"PlatBase"`       // 字段名称（英文）
	Choose  string `json:"choose" binding:"required,oneof='0' '1'" example:"0"` // 是否可被选择 0可选择 1不可选择
	Remark  string `json:"remark" example:"Business type dictionary"`           // 字典描述
	// Pid     uint64 `json:"pid"` // 父级字典ID 默认 1（根数据），暂时不开放变更
}

// DictAddReq 字典创建请求
type DictAddReq struct {
	GroupKey  string `json:"groupKey" binding:"required" example:"serviceCode"` // 字典分组KEY
	Val       string `json:"val" binding:"required" example:"1"`                // 字典值（字符型）
	DictDoReq        // 共通引用
}

// DictEditReq 字典编辑请求
type DictEditReq struct {
	Id        uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	DictDoReq        // 共通引用
}

// DictGetRes 字典查询结果
type DictGetRes struct {
	Id       uint64 `json:"id" example:"1"`                            // 数据ID
	GroupKey string `json:"groupKey" example:"serviceCode"`            // 字典分组KEY
	Val      string `json:"val" example:"1"`                           // 字典值（字符型）
	Label    string `json:"label" example:"PlatBase"`                  // 字段名称
	LabelEn  string `json:"labelEn" example:"PlatBase"`                // 字段名称（英文）
	Choose   string `json:"choose" example:"0"`                        // 是否可被选择 0可选择 1不可选择
	Remark   string `json:"remark" example:"Business type dictionary"` // 字典描述
}

// DictPageReq 字典分页请求
type DictPageReq struct {
	Local    string `json:"-"`                              // 字典语言
	GroupKey string `json:"groupKey" example:"serviceCode"` // 需要查询的字典分组
	Label    string `json:"label" example:"Success"`        // 字典展示文言，跟随请求语言模糊检索
	baseModel.PageReq
}

// DictPageRes 字典分页响应
type DictPageRes struct {
	Id       uint64 `json:"id" example:"1"`                           // 数据ID
	GroupKey string `json:"groupKey" example:"serviceCode"`           // 字典分组KEY
	Label    string `json:"label" example:"Plat"`                     // 字段名称
	LabelEn  string `json:"labelEn" example:"Plat"`                   // 字段名称（英文）
	Choose   string `json:"choose" example:"1"`                       // 是否可被选择 0可选择 1不可选择
	Val      string `json:"val" example:"1"`                          // 字典值（字符型）
	Sort     uint8  `json:"sort" example:"0"`                         // 字典排序
	Remark   string `json:"remark" example:"Business code: Platform"` // 字典描述
	Mark     string `json:"mark" example:"1"`                         // 变更标识 0可变更 1禁止变更
}

// SysConfigGetRes 系统配置查询结果
type SysConfigGetRes struct {
	LoginSwitch       string `json:"loginSwitch" example:"0"`       // 并发限制开关，0限制 1不限制
	LoginNum          uint16 `json:"loginNum" example:"1"`          // 最大登陆并发量，最小为1
	LoginFailSwitch   string `json:"loginFailSwitch" example:"0"`   // 登陆失败限制开关，0限制 1不限制
	LoginFailUnit     string `json:"loginFailUnit" example:"1"`     // 登陆失败限制 1秒 2分 3时 4天
	LoginFailNum      uint16 `json:"loginFailNum" example:"1"`      // 登陆失败最大尝试次数，最小为1
	LoginFailLockUnit string `json:"loginFailLockUnit" example:"1"` // 登陆失败锁定 1秒 2分 3时 4天
	LoginFailLockNum  uint16 `json:"loginFailLockNum" example:"1"`  // 登陆失败锁定数量，最小为1
	LoginFailTryNum   uint16 `json:"loginFailTryNum"`               // 登陆失败尝试次数
	LogoutSwitch      string `json:"logoutSwitch" example:"0"`      // 登陆过期开关，0限制 1不限制
	LogoutUnit        string `json:"logoutUnit" example:"1"`        // 登陆过期单位，1秒 2分 3时 4天
	LogoutNum         uint16 `json:"logoutNum" example:"1"`         // 登陆过期长度数量，最小为1
}

// SysConfigEditReq 系统配置编辑请求
type SysConfigEditReq struct {
	LoginSwitch       string `json:"loginSwitch" binding:"required,oneof='0' '1'" example:"0"`     // 并发限制开关，0限制 1不限制
	LoginNum          uint16 `json:"loginNum" binding:"numeric" example:"1"`                       // 最大登陆并发量，最小为1
	LoginFailSwitch   string `json:"loginFailSwitch" binding:"required,oneof='0' '1'" example:"0"` // 登陆失败限制开关，0限制 1不限制
	LoginFailUnit     string `json:"loginFailUnit" binding:"numeric" example:"1"`                  // 登陆失败限制 1秒 2分 3时 4天
	LoginFailNum      uint16 `json:"loginFailNum" binding:"numeric" example:"1"`                   // 登陆失败最大尝试次数，最小为1
	LoginFailLockUnit string `json:"loginFailLockUnit" binding:"numeric" example:"1"`              // 登陆失败锁定 1秒 2分 3时 4天
	LoginFailLockNum  uint16 `json:"loginFailLockNum" binding:"numeric" example:"1"`               // 登陆失败锁定数量，最小为1
	LoginFailTryNum   uint16 `json:"loginFailTryNum" binding:"numeric" example:"1"`                // 登陆失败尝试次数
	LogoutSwitch      string `json:"logoutSwitch" binding:"required,oneof='0' '1'" example:"0"`    // 登陆过期开关，0限制 1不限制
	LogoutUnit        string `json:"logoutUnit" binding:"numeric" example:"1"`                     // 登陆过期单位，1秒 2分 3时 4天
	LogoutNum         uint16 `json:"logoutNum" binding:"numeric" example:"1"`                      // 登陆过期长度数量，最小为1
}
