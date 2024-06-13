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

// FindPermissionIds 获取权限对应的路由ID
func (t RolePermission) FindPermissionIds(roleId uint64) (res []uint64, err error) {
	r := platDb.Table(t.TableName()).Distinct("permission_id").Where("permission_id", roleId).Find(&res)
	err = r.Error
	return
}

// DeleteByRoleId 根据权限ID移除路由
func (t RolePermission) DeleteByRoleId(roleId uint64) (err error) {
	r := platDb.Where("role_id = ?", roleId).Delete(&t)
	err = r.Error
	return
}
