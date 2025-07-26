package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:255;" json:"name"`
	Label       string         `gorm:"size:255;" json:"label"` // Additional label field as requested
	Slug        string         `gorm:"size:255;" json:"slug"`
	Description string         `gorm:"type:text" json:"description"`
	RID         string         `gorm:"size:255;index" json:"rid"` // RID field for JSON data mapping
	Type        string         `gorm:"size:50;index" json:"type"` // Type: vehicle, maker, model, type, parts
	ParentID    *uint          `gorm:"index" json:"parentId"`
	Parent      *Category      `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"parent,omitempty"`
	Children    []Category     `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Level       int            `gorm:"not null;default:0;index" json:"level"`
	Path        string         `gorm:"size:500;index" json:"path"` // For efficient tree queries
	SortOrder   int            `gorm:"not null;default:0" json:"sortOrder"`
	IsActive    bool           `gorm:"not null;default:true;index" json:"isActive"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Category) TableName() string {
	return "categories"
}

// ProductCategory represents the many-to-many relationship between products and categories
type ProductCategory struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	ProductID  uint      `gorm:"not null;index" json:"productId"`
	CategoryID uint      `gorm:"not null;index" json:"categoryId"`
	Product    Product   `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
	Category   Category  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}

// BeforeCreate hook to set the path and level for new categories
func (c *Category) BeforeCreate(tx *gorm.DB) error {
	if c.ParentID != nil {
		var parent Category
		if err := tx.First(&parent, *c.ParentID).Error; err != nil {
			return err
		}
		c.Level = parent.Level + 1
		c.Path = parent.Path + "/" + parent.Slug
	} else {
		c.Level = 0
		c.Path = ""
	}
	return nil
}

// BeforeUpdate hook to update path and level when parent changes
func (c *Category) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed("ParentID") {
		if c.ParentID != nil {
			var parent Category
			if err := tx.First(&parent, *c.ParentID).Error; err != nil {
				return err
			}
			c.Level = parent.Level + 1
			c.Path = parent.Path + "/" + parent.Slug
		} else {
			c.Level = 0
			c.Path = ""
		}

		// Update all children paths recursively
		var children []Category
		if err := tx.Where("parent_id = ?", c.ID).Find(&children).Error; err != nil {
			return err
		}

		for _, child := range children {
			child.Path = c.Path + "/" + c.Slug
			child.Level = c.Level + 1
			if err := tx.Save(&child).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
