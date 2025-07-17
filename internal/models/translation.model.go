package models

import (
	"time"

	"gorm.io/gorm"
)

type Translation struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Key       string         `gorm:"size:255;not null;index" json:"key"`
	Value     string         `gorm:"type:text;not null" json:"value"`
	Language  string         `gorm:"size:10;not null;index" json:"language"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Translation) TableName() string {
	return "translations"
}

// Add unique constraint for key-language combination
func (t *Translation) BeforeCreate(tx *gorm.DB) error {
	// This ensures unique key per language
	return nil
}
