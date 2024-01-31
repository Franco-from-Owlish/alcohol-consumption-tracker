package database

import (
	"gorm.io/gorm"
	"time"
)

type Omit bool

// Model Shadows gorm.Model, overriding the json tags.
type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
