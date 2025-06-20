package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name            string
	Slug            string
	Description     string
	DescriptionHtml string
	Price           float64      `gorm:"default:0"`
	Discount        float64      `gorm:"default:0"`
	Public          bool         `gorm:"default:true"`
	Category        []Categories `gorm:"many2many:Product_category_joins;constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	Pics            []Files      `gorm:"many2many:product_pics;"`
}
