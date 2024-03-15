package constant

const (
	ProjectName = "smart" // 项目名

	HeaderToken = "Token" // 固定请求头（登陆Token）

	ContextLang    = "Lang"       // 语言
	ContextTraceID = "TraceID"    // 日志链路跟踪ID
	ContextRouterC = "RouterConf" // 路由配置对象

	StatusOpen  = "0" // 正常 启动 可变更
	StatusLock  = "1" // 禁用 锁定 登出 不可变更
	StatusClose = "2" // 移除 弃用 踢出

	DBDuplicateErr = "Error 1062 (23000): Duplicate entry" // 唯一索引错误
)
