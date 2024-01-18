package dbPlat

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
	"time"
)

// LoginRecord 登陆记录
type LoginRecord struct {
	Id        uint64     `json:"id"`        // 默认数据ID
	AccountId uint64     `json:"accountId"` // 账号ID
	LoginType uint8      `json:"loginType"` // 登陆类型 1平台账号登录
	LoginTime *time.Time `json:"loginTime"` // 登陆时间
	Token     string     `json:"token"`     // 登陆Token
	Common
}

// LoginRecordTable 登陆记录泛型造器
var LoginRecordTable actuator.Table[LoginRecord]

// DataBase 实现指定数据库
func (t LoginRecord) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t LoginRecord) TableName() string {
	return "login_record"
}
