package constant

// TransLangSupport 支持更多语言请添加
var TransLangSupport = []string{"zh-CN", "en-US"}

const (
	ProjectName = "smart" // 项目名

	HeaderToken = "Token" // 固定请求头（登陆Token）

	ContextLang    = "Lang"       // 语言
	ContextTraceID = "TraceID"    // 日志链路跟踪ID
	ContextRouterC = "RouterConf" // 路由配置对象
	ContextRouterI = "RouterInDb" // 路由入库对象

	StatusOpen  = "0" // 正常 启动 可变更 选择
	StatusLock  = "1" // 禁用 锁定 登出 不可变更 半选
	StatusClose = "2" // 移除 弃用 踢出

	DBDuplicateErr = "Error 1062 (23000): Duplicate entry" // 唯一索引错误

	CacheSysConf  = ProjectName + "::SysConf"  // 系统配置缓存
	CacheResTrans = ProjectName + "::ResTrans" // 响应码缓存
	CacheRouters  = ProjectName + "::Routers"  // 路由缓存
	CacheRoles    = ProjectName + "::Roles"    // 角色缓存
)
