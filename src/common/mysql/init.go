package mysql

import "siteol.com/smart/src/common/mysql/platDb"

// Init 初始化全部数据库
func Init(traceId string) {
	platDb.InitPlatFromDb(traceId)
}
