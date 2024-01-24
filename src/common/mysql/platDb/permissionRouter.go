package platDb

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// PermissionRouter 权限路由
type PermissionRouter struct {
	Id           uint64 `json:"id"`           // 默认数据ID
	PermissionId uint64 `json:"permissionId"` // 权限ID
	RouterId     uint64 `json:"routerId"`     // 路由ID
}

// PermissionRouterTable 权限路由泛型造器
var PermissionRouterTable actuator.Table[PermissionRouter]

// DataBase 实现指定数据库
func (t PermissionRouter) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t PermissionRouter) TableName() string {
	return "permission_router"
}
