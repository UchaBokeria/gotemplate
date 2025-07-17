package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderNumber   string    `gorm:"uniqueIndex;not null"`
	CustomerName  string    `gorm:"not null"`
	CustomerEmail string    `gorm:"not null"`
	Total         float64   `gorm:"type:decimal(10,2);not null"`
	Status        string    `gorm:"not null;default:'pending'"`
	Date          time.Time `gorm:"not null"`
	Items         string    `gorm:"type:text"`
}
