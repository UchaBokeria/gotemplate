package models

import (
	"time"

	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Views       int
	Title       string
	Body        string
	Public      bool
	PublishedAt time.Time
	Url         string
	ThumbnailID int
	Thumbnail   Files `gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	TypeID      int
	Type        Posts_types `gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
}

type Posts_types struct {
	gorm.Model
	Name string
	Slug string
}
