package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	InvoiceNumber string    `gorm:"uniqueIndex;not null"`
	OrderID       string    `gorm:"not null"`
	CustomerName  string    `gorm:"not null"`
	CustomerEmail string    `gorm:"not null"`
	Amount        float64   `gorm:"type:decimal(10,2);not null"`
	Status        string    `gorm:"not null;default:'pending'"`
	Date          time.Time `gorm:"not null"`
	DueDate       time.Time `gorm:"not null"`
}
