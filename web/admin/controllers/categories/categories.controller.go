package categories

import (
	"fmt"
	"net/http"
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

	query := storage.DB.Model(&models.Category{})

	if listDto.Search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?",
			"%"+listDto.Search+"%", "%"+listDto.Search+"%")
	}

	if listDto.ParentID != nil {
		query = query.Where("parent_id = ?", *listDto.ParentID)
	}

	if listDto.IsActive != nil {
		query = query.Where("is_active = ?", *listDto.IsActive)
	}

	if listDto.Level != nil {
		query = query.Where("level = ?", *listDto.Level)
	}

	query.Count(&total)

	if listDto.Page < 1 {
		listDto.Page = 1
	}
	if listDto.Limit < 1 {
		listDto.Limit = 100
	}

	offset := (listDto.Page - 1) * listDto.Limit
	query = query.Offset(offset).Limit(listDto.Limit)

	query = query.Order("level ASC, sort_order ASC, name ASC")

	if err := query.Find(&categories).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to fetch categories: "+err.Error())
	}

	// Check if this is an HTMX request for table view
	if ctx.Request().Header.Get("HX-Request") != "" {
		return ctx.Html(components.CategoriesTable(categories))
	}

	// Return JSON for API calls (like from product form)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"categories": categories,
		"total":      total,
	})
}

func tree(ctx *controller.Context, listDto *dtos.CategoryListDto) error {
	var categories []models.Category

	query := storage.DB.Model(&models.Category{})

	if listDto.Search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?",
			"%"+listDto.Search+"%", "%"+listDto.Search+"%")
	}

	if listDto.IsActive != nil {
		query = query.Where("is_active = ?", *listDto.IsActive)
	}

	query = query.Order("level ASC, sort_order ASC, name ASC")

	if err := query.Find(&categories).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to fetch categories: "+err.Error())
	}

	// Build tree structure
	treeCategories := buildCategoryTree(categories)

	return ctx.Html(components.CategoriesTree(treeCategories))
}

func show(ctx *controller.Context, idDto *dto.ByID) error {
	var category models.Category

	if err := storage.DB.Preload("Parent").Preload("Children").First(&category, idDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Category not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to fetch category: "+err.Error())
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Category: %s (Level: %d) - %s",
		category.Name, category.Level, category.Description))
}

func create(ctx *controller.Context, createDto *dtos.CreateCategoryDto) error {
	category := models.Category{
		Name:        createDto.Name,
		Slug:        createDto.Slug,
		Description: createDto.Description,
		ParentID:    createDto.ParentID,
		SortOrder:   createDto.SortOrder,
		IsActive:    createDto.IsActive,
	}

	if err := storage.DB.Create(&category).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to create category: "+err.Error())
	}

	var categories []models.Category
	storage.DB.Order("level ASC, sort_order ASC, name ASC").Find(&categories)
	treeCategories := buildCategoryTree(categories)
	return ctx.Html(components.CategoriesTree(treeCategories))
}

func update(ctx *controller.Context, updateDto *dtos.UpdateCategoryDto) error {
	var category models.Category

	if err := storage.DB.First(&category, updateDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Category not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to find category: "+err.Error())
	}

	category.Name = updateDto.Name
	category.Slug = updateDto.Slug
	category.Description = updateDto.Description
	category.ParentID = updateDto.ParentID
	category.SortOrder = updateDto.SortOrder
	category.IsActive = updateDto.IsActive

	if err := storage.DB.Save(&category).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to update category: "+err.Error())
	}

	var categories []models.Category
	storage.DB.Order("level ASC, sort_order ASC, name ASC").Find(&categories)
	treeCategories := buildCategoryTree(categories)
	return ctx.Html(components.CategoriesTree(treeCategories))
}

func deleteCategory(ctx *controller.Context, deleteDto *dto.ByID) error {
	var category models.Category

	if err := storage.DB.First(&category, deleteDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Category not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to find category: "+err.Error())
	}

	// Check if category has children
	var childCount int64
	storage.DB.Model(&models.Category{}).Where("parent_id = ?", category.ID).Count(&childCount)
	if childCount > 0 {
		return ctx.String(http.StatusConflict, "Cannot delete category with subcategories")
	}

	// Check if category is used by products
	var productCount int64
	storage.DB.Model(&models.ProductCategory{}).Where("category_id = ?", category.ID).Count(&productCount)
	if productCount > 0 {
		return ctx.String(http.StatusConflict, "Cannot delete category that is assigned to products")
	}

	if err := storage.DB.Delete(&category).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to delete category: "+err.Error())
	}

	var categories []models.Category
	storage.DB.Order("level ASC, sort_order ASC, name ASC").Find(&categories)
	treeCategories := buildCategoryTree(categories)
	return ctx.Html(components.CategoriesTree(treeCategories))
}

func getChildren(ctx *controller.Context, idDto *dto.ByID) error {
	var children []models.Category

	if err := storage.DB.Where("parent_id = ?", idDto.ID).Order("sort_order ASC, name ASC").Find(&children).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to fetch children: "+err.Error())
	}

	return ctx.Html(components.CategoriesTable(children))
}

func moveCategory(ctx *controller.Context, moveDto *dtos.UpdateCategoryDto) error {
	var category models.Category

	if err := storage.DB.First(&category, moveDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Category not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to find category: "+err.Error())
	}

	// Prevent moving category to itself or its descendants
	if moveDto.ParentID != nil {
		if *moveDto.ParentID == category.ID {
			return ctx.String(http.StatusBadRequest, "Cannot move category to itself")
		}

		// Check if target parent is a descendant
		var targetParent models.Category
		if err := storage.DB.First(&targetParent, *moveDto.ParentID).Error; err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid parent category")
		}

		if strings.Contains(targetParent.Path, category.Slug) {
			return ctx.String(http.StatusBadRequest, "Cannot move category to its descendant")
		}
	}

	category.ParentID = moveDto.ParentID
	category.SortOrder = moveDto.SortOrder

	if err := storage.DB.Save(&category).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to move category: "+err.Error())
	}

	var categories []models.Category
	storage.DB.Order("level ASC, sort_order ASC, name ASC").Find(&categories)
	treeCategories := buildCategoryTree(categories)
	return ctx.Html(components.CategoriesTree(treeCategories))
}

// Helper function to build category tree structure
func buildCategoryTree(categories []models.Category) []models.Category {
	categoryMap := make(map[uint]*models.Category)
	var rootCategories []models.Category

	// First pass: create map of all categories
	for i := range categories {
		categoryMap[categories[i].ID] = &categories[i]
		categories[i].Children = []models.Category{}
	}

	// Second pass: build tree structure
	for i := range categories {
		if categories[i].ParentID == nil {
			rootCategories = append(rootCategories, categories[i])
		} else {
			if parent, exists := categoryMap[*categories[i].ParentID]; exists {
				parent.Children = append(parent.Children, categories[i])
			}
		}
	}

	return rootCategories
}
