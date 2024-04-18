package constant

/**
 *
 * 响应码常量
 * 具体文言维护在数据库，但编码需要在此处维护，提高代码可读性
 *
 * @author 米虫丨www.mebugs.com
 * @since 2023-07-21
 */

const (
	CodeFmt = "%s%d%02d" // 成功码Format

	Success  = "S000" // 默认成功文言（内置禁止修改）
	Fail     = "F000" // 默认业务错误（内置禁止修改）
	Error    = "E000" // 系统未知错误（内置禁止修改）
	ValidErr = "E001" // 参数校验错误（内置禁止修改）（免翻译）
	LoginErr = "E002" // 默认登陆失败（内置禁止修改）
	AuthErr  = "E003" // 默认授权刷新（内置禁止修改）
	PathErr  = "E004" // 默认路由错误（内置禁止修改）

	DictAddSS      = "S100" // 字典创建成功
	DictEditSS     = "S101" // 字典编辑成功
	DictSortSS     = "S102" // 字典排序成功
	DictDelSS      = "S103" // 字典封存成功
	DictGroupGetNG = "F100" // 字典分组查询失败
	DictGetNG      = "F101" // 字典查询失败
	DictUniNG      = "F102" // 字典分组下字典值唯一
	DictMarkNG     = "F103" // 内置字典禁止刪除
	DictSortNG     = "F104" // 字典排序失败

	ResponseAddSS    = "S200" // 响应码创建成功
	ResponseAddSSWNC = "S201" // 响应码创建成功,实际响应码为{{code}}
	ResponseEditSS   = "S202" // 响应码编辑成功
	ResponseDelSS    = "S203" // 响应码封存成功
	ResponseGetNG    = "F200" // 响应码查询失败
	ResponseUniNG    = "F201" // 响应码全局唯一
	ResponseMarkNG   = "F202" // 内置响应码禁止删除

	RouterAddSS     = "S300" // 路由创建成功
	RouterEditSS    = "S301" // 路由编辑成功
	RouterDelSS     = "S302" // 路由删除成功
	RouterGetNG     = "F300" // 路由查询失败
	RouterUniUrlNG  = "F301" // 路由地址全局唯一
	RouterUniNameNG = "F302" // 路由名称全局唯一
	RouterMarkNG    = "F303" // 内置路由禁止删除

	SysConfigEditSS = "S800" // 系统配置编辑成功
	SysConfigGetNG  = "F800" // 系统配置查询失败

)
