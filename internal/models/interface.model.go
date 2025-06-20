package models

import "gorm.io/gorm"

type Interface struct {
	gorm.Model
	Ver          int
	Name         string
	Slug         string
	Presentation string
	SlideShow    []Interface_slideShow
	Inovations   []Interface_inovations
	Products     []Products `gorm:"many2many:Interface_productsCarousel_joins;constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	News         []Posts    `gorm:"many2many:Interface_news_joins;constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	Events       []Posts    `gorm:"many2many:Interface_events_joins;constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	NewVideos    []Files    `gorm:"many2many:Interface_newVideos_joins;constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	SocialMedia  []Social_media
	Languages    []Languages
	Contact      Interface_contact
	About        Interface_about
}

type Interface_slideShow struct {
	gorm.Model
	InterfaceID uint
	Name        string
	Slug        string
	Slogan      string
	Desc        string
	Url         string
	Index       int
	TypeID      int
	Type        File_types
	PicID       *int
	Pic         Files `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:PicID"`
}

type Interface_inovations struct {
	gorm.Model
	InterfaceID uint
	Name        string
	Slug        string
	Title       string
	Url         string
	PicID       *int
	Pic         Files `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:PicID"`
}

type Interface_contact struct {
	gorm.Model
	InterfaceID    uint
	Ver            int
	Name           string
	Slug           string
	Phone          string
	Email          string
	Location       string
	ShortDesc      string
	LocationLink   string
	LocationIframe string
	Copyright      string
}

type Social_media struct {
	gorm.Model
	InterfaceID uint
	Name        string
	Slug        string
	Url         string
	IconID      int
	Icon        Files
}

type Interface_about struct {
	gorm.Model
	InterfaceID uint
	Ver         int
	Body        string
	Terms       string
}
