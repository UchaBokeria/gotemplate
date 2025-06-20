package models

import (
	"main/internal/enums"

	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	ParentID        *int
	Parent          *Categories   `gorm:"foreignKey:ParentID;references:ID;constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	Children        []*Categories `gorm:"foreignkey:ParentID"`
	Level           int           `gorm:"not null"`
	Name            string
	Slug            string
	Description     string
	DescriptionHTML string
	Public          bool
	IconID          *int
	Icon            *Files             `gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	Filters         []Category_filters `gorm:"many2many:Category_filters_joins;constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
}

type Category_filters struct {
	gorm.Model
	Name              string
	Slug              string
	Type              enums.FilterType          `gorm:"type:filter_types"`
	Options           []Category_filters_option `gorm:"many2many:Category_filters_options_join;constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	Default_value     string
	Default_option_id *int
	Default_option    Category_filters_option `gorm:"foreignKey:Default_option_id;"`
}

type Category_filters_option struct {
	gorm.Model
	Name  string
	Key   string
	Value string
}
