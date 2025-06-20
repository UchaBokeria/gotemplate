package models

import (
	"time"

	"gorm.io/gorm"
)

type Branches struct {
	gorm.Model
	Name        string
	Slug        string
	PhoneNumber string
	Map         string
	Shifts      []Branch_shifts
	DistrictID  int
	District    Districts `gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
}

type Branch_shifts struct {
	gorm.Model
	BranchesID uint
	Name       string
	Slug       string
	OpensAt    time.Time `gorm:"type:time"`
	ClosesAt   time.Time `gorm:"type:time"`
}

type Districts struct {
	gorm.Model
	CityID          int
	City            Cities
	Display_name    string
	Display_name_in string
	Lat             string
	Lng             string
	Streets_count   int
}

type Cities struct {
	gorm.Model
	Display_name    string
	Display_name_in string
	Lat             string
	Lng             string
	Streets_count   int
}
