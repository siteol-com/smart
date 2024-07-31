package baseModel

// AccountLoginReq 账号密码登陆请求
type AccountLoginReq struct {
	Acc string `json:"acc" binding:"required" example:"demo"` // 登陆账号
	Pwd string `json:"pwd" binding:"required" example:"demo"` // 登陆密码
}

// AccountLoginRes 账号密码登陆请求
type AccountLoginRes struct {
	Tk string `json:"tk" example:"demo"` // 登陆随机Token
	Re bool   `json:"re" example:"true"` // 登陆需要重置密码
}

// AccountResetReq 密码重置请求
type AccountResetReq struct {
	Pwd    string `json:"pwd" binding:"required" example:"demo"`                // 登陆密码
	PwdCfm string `json:"pwdCfm" binding:"required,eqfield=Pwd" example:"demo"` // 登陆密码确认
}
