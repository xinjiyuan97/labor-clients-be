package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int64          `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
