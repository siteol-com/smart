package model

// DocsReq 演示查询对象
type DocsReq struct {
	HttpCode int `json:"httpCode" example:"200"` // 响应HTTPCode，不传默认响应200，支持200/400/401/403/500
}

// DocsOk 200
type DocsOk struct {
	Code  string `json:"code" example:"S0X0X/F0X0X"` // 响应码
	Msg   string `json:"msg" example:"操作成功/失败"`      // 响应消息
	Data  string `json:"data" example:"响应数据"`        // 响应数据
	UnPop bool   `json:"unPop" example:"true"`       // 免弹窗提示
}

// DocsErr 500
type DocsErr struct {
	Code string `json:"code" example:"E0000"` // 响应码
	Msg  string `json:"msg" example:"系统异常"`   // 响应消息
}

// DocsVail 400
type DocsVail struct {
	Code string `json:"code" example:"E0001"`    // 响应码
	Msg  string `json:"msg" example:"xx字段应该为必填"` // 响应消息
}

// DocsAuthLg 401
type DocsAuthLg struct {
	Code string `json:"code" example:"E0002"` // 响应码
	Msg  string `json:"msg" example:"当前尚未登陆"` // 响应消息
}

// DocsAuthNg 403
type DocsAuthNg struct {
	Code string `json:"code" example:"E0003"`   // 响应码
	Msg  string `json:"msg" example:"当前接口禁止访问"` // 响应消息
}
