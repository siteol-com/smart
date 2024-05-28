package baseModel

// DocsReq 演示查询对象
type DocsReq struct {
	HttpCode int `json:"httpCode" example:"200"` // 响应HTTPCode，不传默认响应200，支持200/400/401/403/500
}

// DocsOk 200
type DocsOk struct {
	Code  string `json:"code" example:"S0X0X/F0X0X"`   // 响应码
	Msg   string `json:"msg" example:"Success/Fail"`   // 响应消息
	Data  string `json:"data" example:"Response data"` // 响应数据
	UnPop bool   `json:"unPop" example:"true"`         // 免弹窗提示
}

// DocsErr 500
type DocsErr struct {
	Code string `json:"code" example:"E0000"`           // 响应码
	Msg  string `json:"msg" example:"System exception"` // 响应消息
}

// DocsVail 400
type DocsVail struct {
	Code string `json:"code" example:"E0001"`                      // 响应码
	Msg  string `json:"msg" example:"xx Field should be required"` // 响应消息
}

// DocsAuthLg 401
type DocsAuthLg struct {
	Code string `json:"code" example:"E0002"`                  // 响应码
	Msg  string `json:"msg" example:"Not currently logged in"` // 响应消息
}

// DocsAuthNg 403
type DocsAuthNg struct {
	Code string `json:"code" example:"E0003"`                                        // 响应码
	Msg  string `json:"msg" example:"Access to the current interface is prohibited"` // 响应消息
}
