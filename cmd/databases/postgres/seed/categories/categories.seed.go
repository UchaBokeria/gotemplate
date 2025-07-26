package categories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/internal/models"
	"main/internal/storage"
	"path/filepath"
	"strings"
)

// JSON data structures
type JSONType struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type JSONModel struct {
	Label string     `json:"label"`
	RID   string     `json:"rid"`
	Types []JSONType `json:"types"`
}

type JSONMaker struct {
	Label  string      `json:"label"`
	RID    string      `json:"rid"`
	Models []JSONModel `json:"models"`
}

type JSONVehicle struct {
	Label  string      `json:"label"`
	RID    int         `json:"rid"`
	Makers []JSONMaker `json:"makers"`
}

// ValidateJSONData validates the structure and content of the parsed JSON
func ValidateJSONData(vehicle *JSONVehicle) error {
	// Validate root vehicle
	if vehicle.Label == "" {
		return fmt.Errorf("vehicle label is empty")
	}
	if vehicle.RID == 0 {
		return fmt.Errorf("vehicle RID is missing or zero")
	}
	if len(vehicle.Makers) == 0 {
		return fmt.Errorf("vehicle '%s' has no makers", vehicle.Label)
	}

	// Validate makers
	for i, maker := range vehicle.Makers {
		if maker.Label == "" {
			return fmt.Errorf("maker #%d in vehicle '%s' has empty label", i+1, vehicle.Label)
		}
		if maker.RID == "" {
			return fmt.Errorf("maker '%s' has empty RID", maker.Label)
		}
		if len(maker.Models) == 0 {
			return fmt.Errorf("maker '%s' has no models", maker.Label)
		}

		// Validate models
		for j, model := range maker.Models {
			if model.Label == "" {
				return fmt.Errorf("model #%d in maker '%s' has empty label", j+1, maker.Label)
			}
			if model.RID == "" {
				return fmt.Errorf("model '%s' has empty RID", model.Label)
			}
			if len(model.Types) == 0 {
				return fmt.Errorf("model '%s' has no types", model.Label)
			}

			// Validate types
			for k, typeItem := range model.Types {
				if typeItem.Label == "" {
					return fmt.Errorf("type #%d in model '%s' has empty label", k+1, model.Label)
				}
				if typeItem.Value == "" {
					return fmt.Errorf("type '%s' has empty value", typeItem.Label)
				}
			}
		}
	}

	return nil
}

// GenerateSlug creates a URL-friendly slug from a string
func GenerateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "&", "and")
	slug = strings.ReplaceAll(slug, "(", "")
	slug = strings.ReplaceAll(slug, ")", "")
	slug = strings.ReplaceAll(slug, ",", "")
	slug = strings.ReplaceAll(slug, ".", "")
	slug = strings.ReplaceAll(slug, "/", "-")
	// Remove multiple consecutive dashes
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}
	slug = strings.Trim(slug, "-")
	return slug
}

// SeedCategoriesFromJSON reads the parsed.json file and creates categories
func SeedCategoriesFromJSON() error {
	// Check if categories already exist
	var count int64
	storage.DB.Model(&models.Category{}).Count(&count)
	if count > 0 {
		fmt.Println("Categories already exist, skipping seeding...")
		return nil
	}

	// Read the JSON file
	jsonPath := filepath.Join("cmd", "databases", "postgres", "seed", "categories", "parsed.json")
	jsonData, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("failed to read parsed.json: %v", err)
	}

	// Parse JSON
	var vehicle JSONVehicle
	if err := json.Unmarshal(jsonData, &vehicle); err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Validate data
	if err := ValidateJSONData(&vehicle); err != nil {
		return fmt.Errorf("JSON validation failed: %v", err)
	}

	fmt.Printf("Starting to seed categories from JSON. Vehicle: %s with %d makers\n", vehicle.Label, len(vehicle.Makers))

	// Create root vehicle category (Level 0)
	vehicleCategory := models.Category{
		Name:        vehicle.Label,
		Label:       vehicle.Label,
		Slug:        GenerateSlug(vehicle.Label),
		Description: fmt.Sprintf("Root category for %s", vehicle.Label),
		RID:         fmt.Sprintf("%d", vehicle.RID),
		Type:        "vehicle",
		Level:       0,
		Path:        "",
		SortOrder:   1,
		IsActive:    true,
	}

	if err := storage.DB.Create(&vehicleCategory).Error; err != nil {
		return fmt.Errorf("failed to create vehicle category: %v", err)
	}
	// Process makers (Level 1)
	for makerIndex, maker := range vehicle.Makers {
		fmt.Printf("Processing maker: %s\n", maker.Label)
		makerCategory := models.Category{
			Name:        maker.Label,
			Label:       maker.Label,
			Slug:        GenerateSlug(fmt.Sprintf("%s-%s", vehicle.Label, maker.Label)),
			Description: fmt.Sprintf("%s vehicles from %s", vehicle.Label, maker.Label),
			RID:         maker.RID,
			Type:        "maker",
			ParentID:    &vehicleCategory.ID,
			Level:       1,
			SortOrder:   makerIndex + 1,
			IsActive:    true,
		}

		if err := storage.DB.Create(&makerCategory).Error; err != nil {
			return fmt.Errorf("failed to create maker category '%s': %v", maker.Label, err)
		}

		// Process models (Level 2)
		for modelIndex, model := range maker.Models {
			fmt.Printf("Processing model: %s\n", model.Label)
			modelCategory := models.Category{
				Name:        model.Label,
				Label:       model.Label,
				Slug:        GenerateSlug(fmt.Sprintf("%s-%s-%s", vehicle.Label, maker.Label, model.Label)),
				Description: fmt.Sprintf("%s model from %s", model.Label, maker.Label),
				RID:         model.RID,
				Type:        "model",
				ParentID:    &makerCategory.ID,
				Level:       2,
				SortOrder:   modelIndex + 1,
				IsActive:    true,
			}

			if err := storage.DB.Create(&modelCategory).Error; err != nil {
				return fmt.Errorf("failed to create model category '%s': %v", model.Label, err)
			}

			// Process types (Level 3)
			for typeIndex, typeItem := range model.Types {
				fmt.Printf("Processing type: %s\n", typeItem.Label)
				typeCategory := models.Category{
					Name:        typeItem.Label,
					Label:       typeItem.Label,
					Slug:        GenerateSlug(fmt.Sprintf("%s-%s-%s-%s", vehicle.Label, maker.Label, model.Label, typeItem.Label)),
					Description: fmt.Sprintf("%s type", typeItem.Label),
					RID:         typeItem.Value,
					Type:        "type",
					ParentID:    &modelCategory.ID,
					Level:       3,
					SortOrder:   typeIndex + 1,
					IsActive:    true,
				}

				if err := storage.DB.Create(&typeCategory).Error; err != nil {
					return fmt.Errorf("failed to create type category '%s': %v", typeItem.Label, err)
				}

				// Create default parts categories (Level 4) - Air Filter and Engine Oil
				defaultParts := []string{"Air Filter", "Engine Oil"}
				for partIndex, partName := range defaultParts {
					fmt.Printf("Processing part: %s\n", partName)
					partCategory := models.Category{
						Name:        partName,
						Label:       partName,
						Slug:        GenerateSlug(fmt.Sprintf("%s-%s-%s-%s-%s", vehicle.Label, maker.Label, model.Label, typeItem.Label, partName)),
						Description: fmt.Sprintf("%s for %s", partName, typeItem.Label),
						RID:         "", // No RID for level 5 categories as specified
						Type:        "parts",
						ParentID:    &typeCategory.ID,
						Level:       4,
						Path:        fmt.Sprintf("/%s/%s/%s/%s", GenerateSlug(vehicle.Label), GenerateSlug(maker.Label), GenerateSlug(model.Label), GenerateSlug(typeItem.Label)),
						SortOrder:   partIndex + 1,
						IsActive:    true,
					}

					if err := storage.DB.Create(&partCategory).Error; err != nil {
						return fmt.Errorf("failed to create part category '%s' for type '%s': %v", partName, typeItem.Label, err)
					}
				}
			}
		}

		fmt.Printf("Processed maker: %s with %d models\n", maker.Label, len(maker.Models))
	}

	fmt.Printf("Successfully seeded categories from JSON. Total makers processed: %d\n", len(vehicle.Makers))
	return nil
}

// GenerateSeedFile creates a Go file with all the seed data based on parsed.json
func GenerateSeedFile() error {
	// Read the JSON file
	jsonPath := filepath.Join("cmd", "databases", "postgres", "seed", "categories", "parsed.json")
	jsonData, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("failed to read parsed.json: %v", err)
	}

	// Parse JSON
	var vehicle JSONVehicle
	if err := json.Unmarshal(jsonData, &vehicle); err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Validate data
	if err := ValidateJSONData(&vehicle); err != nil {
		return fmt.Errorf("JSON validation failed: %v", err)
	}

	// Generate Go code
	var goCode strings.Builder
	goCode.WriteString(`package categories

import (
	"main/internal/models"
	"main/internal/storage"
)

// GeneratedSeedData contains all the category data generated from parsed.json
func GeneratedSeedData() error {
	// Check if categories already exist
	var count int64
	storage.DB.Model(&models.Category{}).Count(&count)
	if count > 0 {
		return nil // Already seeded
	}

	// Create categories with proper parent-child relationships
`)

	// Generate vehicle category (Level 0)
	goCode.WriteString(fmt.Sprintf(`
	// Create root vehicle category (Level 0)
	vehicleCategory := models.Category{
		Name:        %q,
		Label:       %q,
		Slug:        %q,
		Description: %q,
		RID:         %q,
		Type:        "vehicle",
		Level:       0,
		Path:        "",
		SortOrder:   1,
		IsActive:    true,
	}
	if err := storage.DB.Create(&vehicleCategory).Error; err != nil {
		return err
	}

`, vehicle.Label, vehicle.Label, GenerateSlug(vehicle.Label), fmt.Sprintf("Root category for %s", vehicle.Label), fmt.Sprintf("%d", vehicle.RID)))

	// Generate makers and their hierarchies
	for makerIndex, maker := range vehicle.Makers {
		goCode.WriteString(fmt.Sprintf(`
	// Create maker: %s
	maker%d := models.Category{
		Name:        %q,
		Label:       %q,
		Slug:        %q,
		Description: %q,
		RID:         %q,
		Type:        "maker",
		ParentID:    &vehicleCategory.ID,
		Level:       1,
		SortOrder:   %d,
		IsActive:    true,
	}
	if err := storage.DB.Create(&maker%d).Error; err != nil {
		return err
	}
`, maker.Label, makerIndex, maker.Label, maker.Label, GenerateSlug(fmt.Sprintf("%s-%s", vehicle.Label, maker.Label)), fmt.Sprintf("%s vehicles from %s", vehicle.Label, maker.Label), maker.RID, makerIndex+1, makerIndex))

		// Generate models for this maker
		for modelIndex, model := range maker.Models {
			goCode.WriteString(fmt.Sprintf(`
	// Create model: %s
	maker%dModel%d := models.Category{
		Name:        %q,
		Label:       %q,
		Slug:        %q,
		Description: %q,
		RID:         %q,
		Type:        "model",
		ParentID:    &maker%d.ID,
		Level:       2,
		SortOrder:   %d,
		IsActive:    true,
	}
	if err := storage.DB.Create(&maker%dModel%d).Error; err != nil {
		return err
	}
`, model.Label, makerIndex, modelIndex, model.Label, model.Label, GenerateSlug(fmt.Sprintf("%s-%s-%s", vehicle.Label, maker.Label, model.Label)), fmt.Sprintf("%s model from %s", model.Label, maker.Label), model.RID, makerIndex, modelIndex+1, makerIndex, modelIndex))

			// Generate types for this model
			for typeIndex, typeItem := range model.Types {
				goCode.WriteString(fmt.Sprintf(`
	// Create type: %s
	maker%dModel%dType%d := models.Category{
		Name:        %q,
		Label:       %q,
		Slug:        %q,
		Description: %q,
		RID:         %q,
		Type:        "type",
		ParentID:    &maker%dModel%d.ID,
		Level:       3,
		SortOrder:   %d,
		IsActive:    true,
	}
	if err := storage.DB.Create(&maker%dModel%dType%d).Error; err != nil {
		return err
	}
`, typeItem.Label, makerIndex, modelIndex, typeIndex, typeItem.Label, typeItem.Label, GenerateSlug(fmt.Sprintf("%s-%s-%s-%s", vehicle.Label, maker.Label, model.Label, typeItem.Label)), fmt.Sprintf("%s type", typeItem.Label), typeItem.Value, makerIndex, modelIndex, typeIndex+1, makerIndex, modelIndex, typeIndex))

				// Generate default parts for this type
				defaultParts := []string{"Air Filter", "Engine Oil"}
				for partIndex, partName := range defaultParts {
					goCode.WriteString(fmt.Sprintf(`
	// Create part: %s for %s
	_ = models.Category{
		Name:        %q,
		Label:       %q,
		Slug:        %q,
		Description: %q,
		RID:         "",
		Type:        "parts",
		ParentID:    &maker%dModel%dType%d.ID,
		Level:       4,
		Path:        %q,
		SortOrder:   %d,
		IsActive:    true,
	}
	if err := storage.DB.Create(&models.Category{
		Name:        %q,
		Label:       %q,
		Slug:        %q,
		Description: %q,
		RID:         "",
		Type:        "parts",
		ParentID:    &maker%dModel%dType%d.ID,
		Level:       4,
		Path:        %q,
		SortOrder:   %d,
		IsActive:    true,
	}).Error; err != nil {
		return err
	}
`, partName, typeItem.Label, partName, partName, GenerateSlug(fmt.Sprintf("%s-%s-%s-%s-%s", vehicle.Label, maker.Label, model.Label, typeItem.Label, partName)), fmt.Sprintf("%s for %s", partName, typeItem.Label), makerIndex, modelIndex, typeIndex, fmt.Sprintf("/%s/%s/%s/%s", GenerateSlug(vehicle.Label), GenerateSlug(maker.Label), GenerateSlug(model.Label), GenerateSlug(typeItem.Label)), partIndex+1, partName, partName, GenerateSlug(fmt.Sprintf("%s-%s-%s-%s-%s", vehicle.Label, maker.Label, model.Label, typeItem.Label, partName)), fmt.Sprintf("%s for %s", partName, typeItem.Label), makerIndex, modelIndex, typeIndex, fmt.Sprintf("/%s/%s/%s/%s", GenerateSlug(vehicle.Label), GenerateSlug(maker.Label), GenerateSlug(model.Label), GenerateSlug(typeItem.Label)), partIndex+1))
				}
			}
		}
	}

	goCode.WriteString(`
	return nil
}
`)

	// Write to file
	seedFilePath := filepath.Join("cmd", "databases", "postgres", "seed", "categories", "generated_seed.go")
	if err := ioutil.WriteFile(seedFilePath, []byte(goCode.String()), 0644); err != nil {
		return fmt.Errorf("failed to write generated seed file: %v", err)
	}

	fmt.Printf("Generated seed file: %s with data from parsed.json\n", seedFilePath)
	fmt.Printf("Vehicle: %s with %d makers\n", vehicle.Label, len(vehicle.Makers))

	// Count total categories that will be generated
	totalMakers := len(vehicle.Makers)
	totalModels := 0
	totalTypes := 0
	totalParts := 0

	for _, maker := range vehicle.Makers {
		totalModels += len(maker.Models)
		for _, model := range maker.Models {
			totalTypes += len(model.Types)
			totalParts += len(model.Types) * 2 // 2 parts per type
		}
	}

	fmt.Printf("Categories to be generated: %d total (%d makers, %d models, %d types, %d parts)\n",
		1+totalMakers+totalModels+totalTypes+totalParts, totalMakers, totalModels, totalTypes, totalParts)

	return nil
}

// SeedCategories is the original function (kept for backward compatibility)
func SeedCategories() error {
	// For now, call the new JSON-based seeding function
	return SeedCategoriesFromJSON()
}
