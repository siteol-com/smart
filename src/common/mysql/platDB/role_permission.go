package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// RolePermission 角色权限
type RolePermission struct {
	Id           uint64 `json:"id"`           // 默认数据ID
	RoleId       uint64 `json:"roleId"`       // 角色ID
	PermissionId uint64 `json:"permissionId"` // 权限ID
	CheckType    string `json:"checkType"`    // 选择类型 0 check 1 halfCheck
}

// RolePermissionTable 角色权限泛型造器
var RolePermissionTable actuator.Table[RolePermission]

// DataBase 实现指定数据库
func (t RolePermission) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t RolePermission) TableName() string {
	return "role_permission"
}
