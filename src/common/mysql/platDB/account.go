package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
	"time"
)

// Account 账号
type Account struct {
	Id             uint64     `json:"id"`             // 默认数据ID
	Account        string     `json:"account"`        // 登陆账号
	Encryption     string     `json:"encryption"`     // 密码密文
	SaltKey        string     `json:"saltKey"`        // 密码盐值（AES加密KEY）
	DeptId         uint64     `json:"deptId"`         // 部门ID
	PermissionType string     `json:"permissionType"` // 权限类型 0全局数据 1跟随部门 2仅本部门 3本部门及子部门
	PwdExpTime     *time.Time `json:"pwdExpTime"`     // 密码过期时间
	LastLoginTime  *time.Time `json:"lastLoginTime"`  // 最后登陆时间
	Common
}

// AccountTable 账号泛型造器
var AccountTable actuator.Table[Account]

// DataBase 实现指定数据库
func (t Account) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Account) TableName() string {
	return "account"
}
