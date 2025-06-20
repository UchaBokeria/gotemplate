package models

import "gorm.io/gorm"

type Languages struct {
	gorm.Model
	Name               string
	Slug               string
	IconID             int
	Icon               Files
	InterfaceID        uint
	LanguageDictionary []LanguageDictionary `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:LanguageID"`
}

type LanguageDictionary struct {
	gorm.Model
	LanguageID uint
	Key        string
	Value      string
}
