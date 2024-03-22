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
	Success  = "S0000" // 默认成功文言（内置禁止修改）
	Fail     = "F0000" // 默认业务错误（内置禁止修改）
	Error    = "E0000" // 系统未知错误（内置禁止修改）
	ValidErr = "E0001" // 参数校验错误（内置禁止修改）（免翻译）
	LoginErr = "E0002" // 默认授权失败（内置禁止修改）
	AuthErr  = "E0003" // 默认授权刷新（内置禁止修改）

	DictAddSS      = "S0100" // 字典创建成功
	DictEditSS     = "S0101" // 字典编辑成功
	DictSortSS     = "S0102" // 字典排序成功
	DictDelSS      = "S0103" // 字典封存成功
	DictGroupGetNG = "F0100" // 字典分组查询失败
	DictGetNG      = "F0101" // 字典查询失败
	DictUniNG      = "F0102" // 字典分组下字典值唯一
	DictMarkNG     = "F0103" // 字典禁止编辑
	DictSortNG     = "F0104" // 字典排序失败

	SysConfigEditSS = "S0200" // 系统配置编辑成功
	SysConfigGetNG  = "F0200" // 系统配置查询失败

)
