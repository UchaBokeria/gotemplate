package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	UserID uint
	User   Users `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	StatusID uint
	Status   Order_status `gorm:"foreignKey:StatusID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Products []Products `gorm:"many2many:Order_products_joins;"`
	Total    float64    `gorm:"default:0"`
	Comment  string

	Deadline time.Time
}

type Order_status struct {
	gorm.Model
	Name string
}
