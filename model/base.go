package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        int32          `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	IsDeleted bool           `json:"is_deleted"`
}
