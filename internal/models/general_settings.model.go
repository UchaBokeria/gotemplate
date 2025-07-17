package models

import (
	"time"

	"gorm.io/gorm"
)

type GeneralSettings struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CompanyName string         `gorm:"size:255;not null" json:"companyName"`
	Address     string         `gorm:"type:text" json:"address"`
	Phone       string         `gorm:"size:50" json:"phone"`
	Email       string         `gorm:"size:255" json:"email"`
	Facebook    string         `gorm:"size:500" json:"facebook"`
	Instagram   string         `gorm:"size:500" json:"instagram"`
	TikTok      string         `gorm:"size:500" json:"tiktok"`
	AboutUs     string         `gorm:"type:text" json:"aboutUs"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (GeneralSettings) TableName() string {
	return "general_settings"
}
