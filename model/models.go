package model

import "time"

// Model 通用字段
type Model struct {
	ID            int64     `json:"id" gorm:"column:id;primary_key;comment:记录ID"`             // 记录 ID
	RowVersion    int64     `json:"rowVersion" gorm:"column:row_version;comment:记录版本号"`       // 记录版本号
	RowCreateTime time.Time `json:"rowCreateTime" gorm:"column:row_create_time;default:null"` //
	RowUpdateTime time.Time `json:"rowUpdateTime" gorm:"column:row_update_time;default:null"`
	RowIsDeleted  bool      `json:"rowIsDeleted" gorm:"column:row_is_deleted"`
}
