package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Fullname            string
	Password            string
	Token               string
	Email               string
	EmailVerified       bool `gorm:"default:false"`
	EmailVerifyedAt     time.Time
	Phone               string
	TypeID              int
	Type                UserTypes `gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	ResetToken          *string
	ResetTokenExpiresAt time.Time `gorm:"default:null"`
}

type UserTypes struct {
	gorm.Model
	Name string
}
