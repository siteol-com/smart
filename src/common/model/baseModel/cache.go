package baseModel

var CacheRouterNormal = &CacheRouter{}

// CacheRouter 缓存路由对象
type CacheRouter struct {
	Id        uint64   `json:"id"`        // 数据ID，为0表示路由未配置
	NeedAuth  bool     `json:"needAuth"`  // 是否需要授权
	LogInDb   bool     `json:"logInDb"`   // 日志入库
	ReqPrint  bool     `json:"reqPrint"`  // 请求日志打印
	ReqSecure []string `json:"reqSecure"` // 请求日志脱敏数组字段
	ResPrint  bool     `json:"resPrint"`  // 响应日志打印
	ResSecure []string `json:"resSecure"` // 响应日志脱敏数组字段
}
