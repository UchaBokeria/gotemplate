package categories

import (
	"main/internal/models"
	"main/internal/storage"
	"strings"
)

func SeedCategories() error {
	// Check if categories already exist
	var count int64
	storage.DB.Model(&models.Category{}).Count(&count)
	if count > 0 {
		return nil // Already seeded
	}

	// Create root categories
	electronics := models.Category{
		Name:        "Electronics",
		Slug:        "electronics",
		Description: "Electronic devices and accessories",
		SortOrder:   1,
		IsActive:    true,
	}

	clothing := models.Category{
		Name:        "Clothing",
		Slug:        "clothing",
		Description: "Clothing and fashion items",
		SortOrder:   2,
		IsActive:    true,
	}

	homeGarden := models.Category{
		Name:        "Home & Garden",
		Slug:        "home-garden",
		Description: "Home improvement and garden supplies",
		SortOrder:   3,
		IsActive:    true,
	}

	// Create the root categories
	if err := storage.DB.Create(&electronics).Error; err != nil {
		return err
	}
	if err := storage.DB.Create(&clothing).Error; err != nil {
		return err
	}
	if err := storage.DB.Create(&homeGarden).Error; err != nil {
		return err
	}

	// Create electronics subcategories
	electronicsSubcats := []models.Category{
		{
			Name:        "Smartphones",
			Slug:        "smartphones",
			Description: "Mobile phones and accessories",
			ParentID:    &electronics.ID,
			SortOrder:   1,
			IsActive:    true,
		},
		{
			Name:        "Laptops",
			Slug:        "laptops",
			Description: "Portable computers",
			ParentID:    &electronics.ID,
			SortOrder:   2,
			IsActive:    true,
		},
		{
			Name:        "Audio",
			Slug:        "audio",
			Description: "Headphones, speakers, and audio equipment",
			ParentID:    &electronics.ID,
			SortOrder:   3,
			IsActive:    true,
		},
	}

	for _, subcat := range electronicsSubcats {
		if err := storage.DB.Create(&subcat).Error; err != nil {
			return err
		}

		// Create third level categories for smartphones
		if subcat.Slug == "smartphones" {
			smartphoneSubcats := []models.Category{
				{
					Name:        "iPhone",
					Slug:        "iphone",
					Description: "Apple iPhone models",
					ParentID:    &subcat.ID,
					SortOrder:   1,
					IsActive:    true,
				},
				{
					Name:        "Android",
					Slug:        "android",
					Description: "Android smartphones",
					ParentID:    &subcat.ID,
					SortOrder:   2,
					IsActive:    true,
				},
			}

			for _, smartSubcat := range smartphoneSubcats {
				if err := storage.DB.Create(&smartSubcat).Error; err != nil {
					return err
				}
			}
		}
	}

	// Create clothing subcategories
	clothingSubcats := []models.Category{
		{
			Name:        "Men's Clothing",
			Slug:        "mens-clothing",
			Description: "Clothing for men",
			ParentID:    &clothing.ID,
			SortOrder:   1,
			IsActive:    true,
		},
		{
			Name:        "Women's Clothing",
			Slug:        "womens-clothing",
			Description: "Clothing for women",
			ParentID:    &clothing.ID,
			SortOrder:   2,
			IsActive:    true,
		},
		{
			Name:        "Shoes",
			Slug:        "shoes",
			Description: "Footwear for all",
			ParentID:    &clothing.ID,
			SortOrder:   3,
			IsActive:    true,
		},
	}

	for _, subcat := range clothingSubcats {
		if err := storage.DB.Create(&subcat).Error; err != nil {
			return err
		}
	}

	// Create home & garden subcategories
	homeSubcats := []models.Category{
		{
			Name:        "Furniture",
			Slug:        "furniture",
			Description: "Home furniture and decor",
			ParentID:    &homeGarden.ID,
			SortOrder:   1,
			IsActive:    true,
		},
		{
			Name:        "Garden Tools",
			Slug:        "garden-tools",
			Description: "Tools for gardening",
			ParentID:    &homeGarden.ID,
			SortOrder:   2,
			IsActive:    true,
		},
	}

	for _, subcat := range homeSubcats {
		if err := storage.DB.Create(&subcat).Error; err != nil {
			return err
		}
	}

	return nil
}

// Helper function to generate slug from name
func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "&", "and")
	return slug
}
