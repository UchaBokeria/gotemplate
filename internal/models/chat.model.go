package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Fullname     string
	Email        string
	TypeID       int
	Type         Chat_type `gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	ChatStatusID int
	ChatStatus   Chat_status `gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
}

type Chat_type struct {
	gorm.Model
	Name string
	Slug string
}

type Chat_status struct {
	gorm.Model
	Name string
	Slug string
}

type Chat_letters struct {
	gorm.Model
	ChatID         uint
	Chat           Chat `gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	Body           string
	From           string
	To             string
	LetterStatusID int
	LetterStatus   Chat_status `gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
}
