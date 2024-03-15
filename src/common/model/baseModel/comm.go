package baseModel

// IdReq ID查询对象
type IdReq struct {
	Id uint64 `json:"id" example:"1"` // 数据ID
}

// SelectRes 通用下拉响应
type SelectRes struct {
	Label string `json:"label" example:"展示名"` // 展示名
	Value string `json:"value" example:"展示值"` // 展示值
}

// SortReq 排序对象
type SortReq struct {
	ID   uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	Sort uint16 `json:"sort" example:"1"`                  // 序号
}

// SortRes 通用排序前置响应列表
type SortRes struct {
	Id   uint64 `json:"id" example:"1"`      // 默认数据ID
	Name string `json:"name" example:"账号管理"` // 权限名称，界面展示，建议与界面导航一致
	Sort uint8  `json:"sort" example:"1"`    // 权限排序
}
