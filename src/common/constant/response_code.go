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
	CodeFmt = "%s%s%02d" // 成功码Format

	Success  = "S000" // 处理成功（默认）
	Fail     = "F000" // 处理失败（默认）
	Error    = "E000" // 系统异常（默认）
	ValidErr = "E001" // 参数非法（默认）（免翻译）
	LoginErr = "E002" // 尚未登陆（默认）
	AuthErr  = "E003" // 无权访问（默认）
	PathErr  = "E004" // 路径不存在（默认）

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
	RouterDelNG     = "F304" // 路由删除

	PermissionAddSS      = "S400" // 权限创建成功
	PermissionEditSS     = "S401" // 权限编辑成功
	PermissionDelSS      = "S402" // 权限删除成功
	PermissionSortSS     = "S403" // 权限删除成功
	PermissionGetNG      = "F400" // 权限查询失败
	PermissionUniAliasNG = "F401" // 权限别名全局唯一
	PermissionUniNameNG  = "F402" // 权限名称全局唯一
	PermissionMarkNG     = "F403" // 内置权限禁止删除
	PermissionRouterNG   = "F404" // 权限配置路由同步失败
	PermissionDelNG      = "F405" // 权限配置删除失败

	RoleAddSS        = "S500" // 角色创建成功
	RoleEditSS       = "S501" // 角色编辑成功
	RoleDelSS        = "S502" // 角色删除失败
	RoleGetNG        = "F500" // 角色查询失败
	RoleMarkNG       = "F501" // 内置角色禁止编辑
	RoleNameUniNG    = "F502" // 角色名全局唯一
	RolePermissionNG = "F503" // 角色权限配置失败

	SysConfigEditSS = "S800" // 系统配置编辑成功
	SysConfigGetNG  = "F800" // 系统配置查询失败

)
