package categories

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"main/internal/models"
	"main/internal/storage"
	"main/internal/types/dto"
	"main/web/admin/controllers/categories/dtos"
	"main/web/admin/view/components"

	"github.com/UchaBokeria/goyard/controller"
	"gorm.io/gorm"
)

func list(ctx *controller.Context, listDto *dtos.CategoryListDto) error {
	var categories []models.Category
	var total int64

	query := storage.DB.Model(&models.Category{}).Preload("Parent").Preload("Children")

	// Apply filters
	if listDto.Search != "" {
		query = query.Where("name ILIKE ? OR label ILIKE ? OR description ILIKE ?",
			"%"+listDto.Search+"%", "%"+listDto.Search+"%", "%"+listDto.Search+"%")
	}

	if listDto.Path != "" {
		query = query.Where("path ILIKE ?", "%"+listDto.Path+"%")
	}

	if listDto.Level != nil {
		query = query.Where("level = ?", *listDto.Level)
	}

	if listDto.ParentID != nil {
		query = query.Where("parent_id = ?", *listDto.ParentID)
	}

	if listDto.Type != "" {
		query = query.Where("type = ?", listDto.Type)
	}

	if len(listDto.VehicleIDs) > 0 {
		query = query.Where("type = ? AND id IN ?", "vehicle", listDto.VehicleIDs)
	}

	if len(listDto.MakerIDs) > 0 {
		query = query.Where("type = ? AND id IN ?", "maker", listDto.MakerIDs)
	}

	if len(listDto.ModelIDs) > 0 {
		query = query.Where("type = ? AND id IN ?", "model", listDto.ModelIDs)
	}

	if len(listDto.TypeIDs) > 0 {
		query = query.Where("type = ? AND id IN ?", "type", listDto.TypeIDs)
	}

	if len(listDto.CategoryIDs) > 0 {
		query = query.Where("type = ? AND id IN ?", "parts", listDto.CategoryIDs)
	}

	if listDto.IsActive != nil {
		query = query.Where("is_active = ?", *listDto.IsActive)
	}

	query.Count(&total)

	if listDto.TreeView {
		// For tree view, get all categories and organize them hierarchically
		query = query.Order("level ASC, sort_order ASC, name ASC")
		if err := query.Find(&categories).Error; err != nil {
			return ctx.String(http.StatusInternalServerError, "Failed to fetch categories: "+err.Error())
		}
		return ctx.Html(components.CategoriesTree(categories))
	} else {
		// Regular pagination for table view
		if listDto.Page < 1 {
			listDto.Page = 1
		}
		if listDto.Limit < 1 {
			listDto.Limit = 100
		}

		offset := (listDto.Page - 1) * listDto.Limit
		query = query.Offset(offset).Limit(listDto.Limit).Order("level ASC, sort_order ASC, name ASC")

		if err := query.Find(&categories).Error; err != nil {
			return ctx.String(http.StatusInternalServerError, "Failed to fetch categories: "+err.Error())
		}
		return ctx.Html(components.CategoriesTree(categories))
	}
}

func show(ctx *controller.Context, idDto *dto.ByID) error {
	var category models.Category

	if err := storage.DB.Preload("Parent").Preload("Children").First(&category, idDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Category not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to fetch category: "+err.Error())
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Category: %s (Level: %d, Type: %s)",
		category.Name, category.Level, category.Type))
}

func create(ctx *controller.Context, createDto *dtos.CreateCategoryDto) error {
	// Generate slug from name
	categorySlug := generateSlug(createDto.Name)

	// Determine type and level based on parent
	category := models.Category{
		Name:        createDto.Name,
		Label:       createDto.Label,
		Slug:        categorySlug,
		Description: createDto.Description,
		SortOrder:   createDto.SortOrder,
		IsActive:    createDto.IsActive,
	}

	if createDto.ParentID != "" {
		parentID, err := strconv.ParseUint(createDto.ParentID, 10, 64)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid parent ID")
		}
		var parent models.Category
		if err := storage.DB.First(&parent, parentID).Error; err != nil {
			return ctx.String(http.StatusBadRequest, "Parent category not found")
		}

		category.Level = parent.Level + 1
		category.Path = parent.Path + "/" + parent.Slug

		// Determine type based on level
		switch category.Level {
		case 0:
			category.Type = "vehicle"
		case 1:
			category.Type = "maker"
		case 2:
			category.Type = "model"
		case 3:
			category.Type = "type"
		case 4:
			category.Type = "parts"
		default:
			return ctx.String(http.StatusBadRequest, "Maximum category depth exceeded")
		}
	} else {
		category.ParentID = nil
		category.Level = 0
		category.Path = ""
		category.Type = "vehicle"
	}

	fmt.Println(category)
	if err := storage.DB.Create(&category).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to create category: "+err.Error())
	}

	// Return updated tree
	var categories []models.Category
	storage.DB.Preload("Parent").Preload("Children").Order("level ASC, sort_order ASC, name ASC").Find(&categories)
	return ctx.Html(components.CategoriesTree(categories))
}

func update(ctx *controller.Context, updateDto *dtos.UpdateCategoryDto) error {
	var category models.Category

	if err := storage.DB.Preload("Parent").Preload("Children").First(&category, updateDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Category not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to find category: "+err.Error())
	}

	// Update only allowed fields (not RID, level, type, parent, path, slug)
	category.Name = updateDto.Name
	category.Label = updateDto.Label
	category.Description = updateDto.Description
	category.SortOrder = updateDto.SortOrder
	category.IsActive = updateDto.IsActive

	// Update slug if name changed
	newSlug := generateSlug(updateDto.Name)
	if category.Slug != newSlug {
		category.Slug = newSlug
		// Note: Path will be updated automatically by the model's BeforeUpdate hook
	}

	if err := storage.DB.Save(&category).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to update category: "+err.Error())
	}

	// Return updated tree
	var categories []models.Category
	storage.DB.Preload("Parent").Preload("Children").Order("level ASC, sort_order ASC, name ASC").Find(&categories)
	return ctx.Html(components.CategoriesTree(categories))
}

func deleteCategory(ctx *controller.Context, deleteDto *dto.ByID) error {
	var category models.Category

	if err := storage.DB.Preload("Children").First(&category, deleteDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Category not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to find category: "+err.Error())
	}

	// Check if category has children
	if len(category.Children) > 0 {
		return ctx.String(http.StatusBadRequest, "Cannot delete category with children. Please delete or move child categories first.")
	}

	// Check if category is associated with any products
	var productCount int64
	storage.DB.Model(&models.ProductCategory{}).Where("category_id = ?", category.ID).Count(&productCount)
	if productCount > 0 {
		return ctx.String(http.StatusBadRequest, "Cannot delete category that is associated with products. Please remove product associations first.")
	}

	if err := storage.DB.Delete(&category).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to delete category: "+err.Error())
	}

	// Return updated tree
	var categories []models.Category
	storage.DB.Preload("Parent").Preload("Children").Order("level ASC, sort_order ASC, name ASC").Find(&categories)
	return ctx.Html(components.CategoriesTree(categories))
}

// Helper function to get categories for filters
func getCategoriesByType(ctx *controller.Context, categoryType string) error {
	var categories []models.Category

	query := storage.DB.Where("type = ? AND is_active = ?", categoryType, true).
		Order("level ASC, sort_order ASC, name ASC")

	if err := query.Find(&categories).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to fetch categories: "+err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"categories": categories,
	})
}

// generateSlug creates a URL-friendly slug from a string
func generateSlug(input string) string {
	// Convert to lowercase
	slug := strings.ToLower(input)

	// Replace spaces and special characters with hyphens
	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug = re.ReplaceAllString(slug, "-")

	// Remove leading and trailing hyphens
	slug = strings.Trim(slug, "-")

	return slug
}
