package categories

import (
	"fmt"
	"log"
	"main/internal/models"
	"main/internal/storage"
)

// TestSeedProcess demonstrates the complete seeding process
func TestSeedProcess() {
	fmt.Println("=== Testing Category Seeding Process ===")

	// Initialize database connection (assuming it's already set up)
	if storage.DB == nil {
		log.Fatal("Database connection not initialized")
	}

	// Step 1: Generate seed file from JSON (this now reads parsed.json directly)
	fmt.Println("Step 1: Generating seed file from parsed.json...")
	if err := GenerateSeedFile(); err != nil {
		log.Fatalf("Failed to generate seed file: %v", err)
	}

	// Step 2: Seed categories from JSON to database
	fmt.Println("Step 2: Seeding categories to database...")
	if err := SeedCategoriesFromJSON(); err != nil {
		log.Fatalf("Failed to seed categories: %v", err)
	}

	// Step 3: Show statistics
	fmt.Println("Step 3: Showing statistics...")
	ShowCategoryStats()

	fmt.Println("=== Seeding Process Completed Successfully ===")
}

// ShowCategoryStats displays statistics about the seeded categories
func ShowCategoryStats() {
	var stats struct {
		Total   int64
		ByLevel map[int]int64
		ByType  map[string]int64
	}

	stats.ByLevel = make(map[int]int64)
	stats.ByType = make(map[string]int64)

	// Get total count
	storage.DB.Model(&models.Category{}).Count(&stats.Total)

	// Get count by level
	for level := 0; level <= 4; level++ {
		var count int64
		storage.DB.Model(&models.Category{}).Where("level = ?", level).Count(&count)
		stats.ByLevel[level] = count
	}

	// Get count by type
	types := []string{"vehicle", "maker", "model", "type", "parts"}
	for _, catType := range types {
		var count int64
		storage.DB.Model(&models.Category{}).Where("type = ?", catType).Count(&count)
		stats.ByType[catType] = count
	}

	// Display statistics
	fmt.Printf("Total Categories: %d\n", stats.Total)
	fmt.Println("Categories by Level:")
	for level := 0; level <= 4; level++ {
		levelName := getLevelName(level)
		fmt.Printf("  Level %d (%s): %d\n", level, levelName, stats.ByLevel[level])
	}

	fmt.Println("Categories by Type:")
	for _, catType := range types {
		fmt.Printf("  %s: %d\n", catType, stats.ByType[catType])
	}
}

// getLevelName returns a human-readable name for a category level
func getLevelName(level int) string {
	switch level {
	case 0:
		return "Vehicle"
	case 1:
		return "Maker/Brand"
	case 2:
		return "Model"
	case 3:
		return "Type"
	case 4:
		return "Parts"
	default:
		return "Unknown"
	}
}
