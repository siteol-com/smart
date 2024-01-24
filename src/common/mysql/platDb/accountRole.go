package platDb

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// AccountRole 账号角色
type AccountRole struct {
	Id        uint64 `json:"id"`        // 默认数据ID
	AccountId uint64 `json:"accountId"` // 账号ID
	RoleId    uint64 `json:"roleId"`    // 角色ID
}

// AccountRoleTable 账号角色泛型造器
var AccountRoleTable actuator.Table[AccountRole]

// DataBase 实现指定数据库
func (t AccountRole) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t AccountRole) TableName() string {
	return "account_role"
}
