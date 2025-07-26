package products

import (
	"main/internal/models"
	"main/internal/storage"
)

func Populate() error {
	// Check if products already exist
	var count int64
	storage.DB.Model(&models.Product{}).Count(&count)
	if count > 0 {
		return nil // Already seeded
	}

	// Get categories for association
	var categories []models.Category
	if err := storage.DB.Find(&categories).Error; err != nil {
		return err
	}

	// Create a map for easier category lookup
	categoryMap := make(map[string]uint)
	for _, cat := range categories {
		categoryMap[cat.Slug] = cat.ID
	}

	// Define products with their category associations
	type ProductSeed struct {
		Product       models.Product
		CategorySlugs []string
	}

	productSeeds := []ProductSeed{
		// Electronics -> Smartphones -> iPhone
		{
			Product: models.Product{
				Name:        "iPhone 15 Pro",
				Price:       999.99,
				Stock:       50,
				Status:      "active",
				Description: "Latest iPhone with advanced camera system and titanium design.",
			},
			CategorySlugs: []string{"iphone", "smartphones"},
		},
		{
			Product: models.Product{
				Name:        "iPhone 14",
				Price:       799.99,
				Stock:       75,
				Status:      "active",
				Description: "Previous generation iPhone with great performance and features.",
			},
			CategorySlugs: []string{"iphone", "smartphones"},
		},

		// Electronics -> Smartphones -> Android
		{
			Product: models.Product{
				Name:        "Samsung Galaxy S24",
				Price:       899.99,
				Stock:       60,
				Status:      "active",
				Description: "Premium Android smartphone with AI-powered features.",
			},
			CategorySlugs: []string{"android", "smartphones"},
		},
		{
			Product: models.Product{
				Name:        "Google Pixel 8",
				Price:       699.99,
				Stock:       40,
				Status:      "active",
				Description: "Pure Android experience with exceptional camera capabilities.",
			},
			CategorySlugs: []string{"android", "smartphones"},
		},

		// Electronics -> Laptops
		{
			Product: models.Product{
				Name:        "MacBook Pro 16\"",
				Price:       2499.99,
				Stock:       25,
				Status:      "active",
				Description: "Professional laptop with M3 Pro chip for demanding workflows.",
			},
			CategorySlugs: []string{"laptops"},
		},
		{
			Product: models.Product{
				Name:        "Dell XPS 13",
				Price:       1299.99,
				Stock:       30,
				Status:      "active",
				Description: "Ultrabook with premium build quality and excellent display.",
			},
			CategorySlugs: []string{"laptops"},
		},
		{
			Product: models.Product{
				Name:        "Gaming Laptop RTX 4070",
				Price:       1799.99,
				Stock:       20,
				Status:      "active",
				Description: "High-performance gaming laptop with latest graphics card.",
			},
			CategorySlugs: []string{"laptops"},
		},

		// Electronics -> Audio
		{
			Product: models.Product{
				Name:        "AirPods Pro",
				Price:       249.99,
				Stock:       100,
				Status:      "active",
				Description: "Premium wireless earbuds with active noise cancellation.",
			},
			CategorySlugs: []string{"audio"},
		},
		{
			Product: models.Product{
				Name:        "Sony WH-1000XM5",
				Price:       399.99,
				Stock:       45,
				Status:      "active",
				Description: "Industry-leading noise canceling headphones.",
			},
			CategorySlugs: []string{"audio"},
		},
		{
			Product: models.Product{
				Name:        "Bluetooth Speaker",
				Price:       129.99,
				Stock:       80,
				Status:      "active",
				Description: "Portable wireless speaker with rich, clear sound.",
			},
			CategorySlugs: []string{"audio"},
		},

		// Clothing -> Men's Clothing
		{
			Product: models.Product{
				Name:        "Men's Cotton T-Shirt",
				Price:       24.99,
				Stock:       200,
				Status:      "active",
				Description: "Comfortable 100% cotton t-shirt in various colors.",
			},
			CategorySlugs: []string{"mens-clothing"},
		},
		{
			Product: models.Product{
				Name:        "Men's Jeans",
				Price:       79.99,
				Stock:       150,
				Status:      "active",
				Description: "Classic fit denim jeans with modern styling.",
			},
			CategorySlugs: []string{"mens-clothing"},
		},

		// Clothing -> Women's Clothing
		{
			Product: models.Product{
				Name:        "Women's Summer Dress",
				Price:       59.99,
				Stock:       120,
				Status:      "active",
				Description: "Elegant summer dress perfect for casual and formal occasions.",
			},
			CategorySlugs: []string{"womens-clothing"},
		},
		{
			Product: models.Product{
				Name:        "Women's Blouse",
				Price:       39.99,
				Stock:       180,
				Status:      "active",
				Description: "Professional blouse suitable for office and business settings.",
			},
			CategorySlugs: []string{"womens-clothing"},
		},

		// Clothing -> Shoes
		{
			Product: models.Product{
				Name:        "Running Shoes",
				Price:       119.99,
				Stock:       90,
				Status:      "active",
				Description: "Lightweight running shoes with excellent cushioning and support.",
			},
			CategorySlugs: []string{"shoes"},
		},
		{
			Product: models.Product{
				Name:        "Leather Boots",
				Price:       199.99,
				Stock:       60,
				Status:      "active",
				Description: "Premium leather boots with durable construction and timeless style.",
			},
			CategorySlugs: []string{"shoes"},
		},

		// Home & Garden -> Furniture
		{
			Product: models.Product{
				Name:        "Dining Table Set",
				Price:       599.99,
				Stock:       15,
				Status:      "active",
				Description: "Solid wood dining table with 4 matching chairs.",
			},
			CategorySlugs: []string{"furniture"},
		},
		{
			Product: models.Product{
				Name:        "Office Chair",
				Price:       299.99,
				Stock:       35,
				Status:      "active",
				Description: "Ergonomic office chair with lumbar support and adjustable height.",
			},
			CategorySlugs: []string{"furniture"},
		},

		// Home & Garden -> Garden Tools
		{
			Product: models.Product{
				Name:        "Garden Tool Set",
				Price:       89.99,
				Stock:       70,
				Status:      "active",
				Description: "Complete garden tool set with hand tools and storage bag.",
			},
			CategorySlugs: []string{"garden-tools"},
		},
		{
			Product: models.Product{
				Name:        "Electric Lawn Mower",
				Price:       449.99,
				Stock:       25,
				Status:      "active",
				Description: "Cordless electric lawn mower with self-propelled technology.",
			},
			CategorySlugs: []string{"garden-tools"},
		},

		// Cross-category products (products that belong to multiple categories)
		{
			Product: models.Product{
				Name:        "Smart Watch",
				Price:       349.99,
				Stock:       80,
				Status:      "active",
				Description: "Multi-functional smart watch with health tracking and phone integration.",
			},
			CategorySlugs: []string{"electronics", "audio"}, // Could be in electronics and audio
		},
		{
			Product: models.Product{
				Name:        "Wireless Charging Pad",
				Price:       39.99,
				Stock:       150,
				Status:      "active",
				Description: "Universal wireless charging pad compatible with most smartphones.",
			},
			CategorySlugs: []string{"electronics", "smartphones"}, // Accessories for smartphones
		},

		// Test products with different statuses
		{
			Product: models.Product{
				Name:        "Test Product - Inactive",
				Price:       99.99,
				Stock:       0,
				Status:      "inactive",
				Description: "This is an inactive product for testing purposes.",
			},
			CategorySlugs: []string{"electronics"},
		},
		{
			Product: models.Product{
				Name:        "Draft Product - Coming Soon",
				Price:       45.99,
				Stock:       0,
				Status:      "draft",
				Description: "This product is in draft status and will be available soon.",
			},
			CategorySlugs: []string{"clothing"},
		},
	}

	// Create products and associate with categories
	for _, productSeed := range productSeeds {
		// Start transaction
		tx := storage.DB.Begin()
		
		// Create the product
		if err := tx.Create(&productSeed.Product).Error; err != nil {
			tx.Rollback()
			return err
		}

		// Associate with categories
		for _, slug := range productSeed.CategorySlugs {
			if categoryID, exists := categoryMap[slug]; exists {
				productCategory := models.ProductCategory{
					ProductID:  productSeed.Product.ID,
					CategoryID: categoryID,
				}
				if err := tx.Create(&productCategory).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
		}

		tx.Commit()
	}

	return nil
}
