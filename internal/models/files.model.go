package models

import "gorm.io/gorm"

type Files struct {
	gorm.Model
	Name       string
	Original   string
	Location   string
	Path       string
	Size       int
	Base64     string
	Compressed bool
	TypeID     int
	Type       File_types `gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
}

type File_types struct {
	gorm.Model
	Name     string
	Ext      string
	Max_size int
}
