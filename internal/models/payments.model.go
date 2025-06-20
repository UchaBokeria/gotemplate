package models

import "gorm.io/gorm"

type Payments struct {
	gorm.Model
	Amount          float64
	OrderID         int
	Order           Orders `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Paid            bool   `gorm:"default:false"`
	Ref             string
	Comment         string
	StatusID        int
	Status          PaymentStatus      `gorm:"foreignKey:StatusID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaymentBranches []*PaymentBranches `gorm:"many2many:PaymentBranches_joins;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type PaymentStatus struct {
	gorm.Model
	Name string
}

type PaymentBranches struct {
	gorm.Model
	Name string
	Slug string
}
