package platDb

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// DictGroup 字典分組
type DictGroup struct {
	Id     uint64 `json:"id"`     // 字典分组ID
	Key    string `json:"key"`    // 字典分组Key
	Name   string `json:"name"`   // 字典分组名称
	NameEn string `json:"nameEn"` // 字典分组名称英文
	Common
}

// DictGroupTable 字典分組泛型造器
var DictGroupTable actuator.Table[DictGroup]

// DataBase 实现指定数据库
func (t DictGroup) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t DictGroup) TableName() string {
	return "dict_group"
}
