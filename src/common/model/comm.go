package model

// SelectRes 通用下拉响应
type SelectRes struct {
	Label string `json:"label" example:"展示名"` // 展示名
	Value string `json:"value" example:"展示值"` // 展示值
}
