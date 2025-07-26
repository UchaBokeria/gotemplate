# Category Seeding from JSON Data

This package provides functionality to seed your categories table from the `parsed.json` file containing automotive data.

## Structure

The seeding creates a 5-level hierarchy:

1. **Level 0 (Vehicle)**: Root category (e.g., "cars")
2. **Level 1 (Maker/Brand)**: Vehicle manufacturers (e.g., "Abarth", "BMW")
3. **Level 2 (Model)**: Vehicle models (e.g., "124 Spider, 348 (2016-2021)")
4. **Level 3 (Type)**: Specific vehicle types (e.g., "124 Spider 1.4 Turbo Multi Air (2016-2017)")
5. **Level 4 (Parts)**: Auto parts for each type (default: "Air Filter", "Engine Oil")

## Database Schema Updates

The `Category` model has been extended with:
- `Label` field: Additional label from JSON data
- `RID` field: Reference ID from JSON data (empty for Level 4 parts)
- `Type` field: Category type ("vehicle", "maker", "model", "type", "parts")

## Functions

### `SeedCategoriesFromJSON()`
- Reads `parsed.json` file
- Validates the JSON structure
- Creates categories in the database with proper hierarchy
- Automatically adds "Air Filter" and "Engine Oil" for each vehicle type

### `GenerateSeedFile()`
- Reads and parses the `parsed.json` file directly
- Generates a Go file (`generated_seed.go`) with static seed data
- Creates code that will recreate the exact same hierarchy when executed
- Useful for creating reproducible seed data for deployment

### `ValidateJSONData(vehicle *JSONVehicle)`
- Validates the JSON structure for consistency
- Checks for missing labels, RIDs, and empty collections
- Returns detailed error messages for debugging

## Usage

```go
package main

import (
    "main/cmd/databases/postgres/seed/categories"
    "main/internal/storage"
)

func main() {
    // Initialize database connection
    storage.InitDB()
    
    // Seed categories from JSON
    if err := categories.SeedCategoriesFromJSON(); err != nil {
        log.Fatal(err)
    }
    
    // Generate static seed file
    if err := categories.GenerateSeedFile(); err != nil {
        log.Fatal(err)
    }
    
    // Show statistics
    categories.ShowCategoryStats()
}
```

## Files

- `categories.seed.go`: Main seeding logic
- `test_seed.go`: Test functions and statistics
- `generated_seed.go`: Auto-generated static seed data (created after running seed)
- `parsed.json`: Source JSON data with automotive categories

## Error Handling

The seeding process includes comprehensive validation:
- Checks for missing or empty labels
- Validates RID presence for levels 0-3
- Ensures proper hierarchy structure
- Provides detailed error messages with context

## Notes

- Only Level 4 (parts) categories have the `Path` field populated
- Level 4 categories don't have RID values (as specified)
- The seeding process is idempotent - it won't duplicate data if run multiple times
- All categories are created as active by default 