package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// Dept 部门
type Dept struct {
	Id             uint64 `json:"id"`             // 默认数据ID
	Name           string `json:"name"`           // 部门名称
	Pid            uint64 `json:"pid"`            // 父级部门ID，租户创建时默认创建根部门，父级ID=0
	Sort           uint16 `json:"sort"`           // 同级部门排序
	PermissionType string `json:"permissionType"` // 权限类型 0全局数据 2仅本部门 3本部门及子部门
	Common
}

// DeptTable 部门泛型造器
var DeptTable actuator.Table[Dept]

// DataBase 实现指定数据库
func (t Dept) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Dept) TableName() string {
	return "dept"
}
