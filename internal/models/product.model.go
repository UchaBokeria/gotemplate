package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:255;not null" json:"name"`
	Category    string         `gorm:"size:100;not null;index" json:"category"`
	Price       float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock       int            `gorm:"not null;default:0" json:"stock"`
	Status      string         `gorm:"size:20;not null;default:'active';index" json:"status"`
	Description string         `gorm:"type:text" json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Product) TableName() string {
	return "products"
}
