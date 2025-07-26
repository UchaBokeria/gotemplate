package products

import (
	"fmt"
	"net/http"
	"strings"

	"main/internal/models"
	"main/internal/storage"
	"main/internal/types/dto"
	"main/web/admin/controllers/products/dtos"
	"main/web/admin/view/components"

	"github.com/UchaBokeria/goyard/controller"
	"gorm.io/gorm"
)

func list(ctx *controller.Context, listDto *dtos.ProductListDto) error {
	var products []models.Product
	var total int64

	query := storage.DB.Model(&models.Product{}).Preload("Categories")

	if listDto.Search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?",
			"%"+listDto.Search+"%", "%"+listDto.Search+"%")
	}

	if listDto.CategoryIDs != "" {
		categoryIDs := strings.Split(listDto.CategoryIDs, ",")
		query = query.Joins("JOIN product_categories ON products.id = product_categories.product_id").
			Where("product_categories.category_id IN ?", categoryIDs).
			Distinct()
	}

	if len(listDto.StatusMulti) > 0 {
		query = query.Where("status IN ?", listDto.StatusMulti)
	}

	query.Count(&total)

	if listDto.Page < 1 {
		listDto.Page = 1
	}
	if listDto.Limit < 1 {
		listDto.Limit = 20
	}

	offset := (listDto.Page - 1) * listDto.Limit
	query = query.Offset(offset).Limit(listDto.Limit)

	query = query.Order("id DESC")

	if err := query.Find(&products).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to fetch products: "+err.Error())
	}

	return ctx.Html(components.ProductsTable(products))
}

func show(ctx *controller.Context, idDto *dto.ByID) error {
	var product models.Product

	if err := storage.DB.Preload("Categories").First(&product, idDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Product not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to fetch product: "+err.Error())
	}

	categoryNames := make([]string, len(product.Categories))
	for i, cat := range product.Categories {
		categoryNames[i] = cat.Name
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Product: %s (Categories: %v) - $%.2f - Stock: %d - Status: %s",
		product.Name, categoryNames, product.Price, product.Stock, product.Status))
}

func create(ctx *controller.Context, createDto *dtos.CreateProductDto) error {
	product := models.Product{
		Name:        createDto.Name,
		Price:       createDto.Price,
		Stock:       createDto.Stock,
		Status:      createDto.Status,
		Description: createDto.Description,
	}

	// Start transaction
	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return ctx.String(http.StatusInternalServerError, "Failed to create product: "+err.Error())
	}

	// Associate categories
	if createDto.CategoryIDs != "" {
		categoryIDs := strings.Split(createDto.CategoryIDs, ",")
		var categories []models.Category
		if err := tx.Find(&categories, categoryIDs).Error; err != nil {
			tx.Rollback()
			return ctx.String(http.StatusInternalServerError, "Failed to find categories: "+err.Error())
		}

		// Validate that all categories are level 4 (parts level)
		for _, category := range categories {
			if category.Level != 4 {
				tx.Rollback()
				return ctx.String(http.StatusBadRequest, fmt.Sprintf("Products can only be associated with level 4 categories (parts). Category '%s' is level %d", category.Name, category.Level))
			}
		}

		if err := tx.Model(&product).Association("Categories").Append(categories); err != nil {
			tx.Rollback()
			return ctx.String(http.StatusInternalServerError, "Failed to associate categories: "+err.Error())
		}
	}

	tx.Commit()

	var products []models.Product
	storage.DB.Preload("Categories").Order("id DESC").Find(&products)
	return ctx.Html(components.ProductsTable(products))
}

func update(ctx *controller.Context, updateDto *dtos.UpdateProductDto) error {
	var product models.Product

	if err := storage.DB.Preload("Categories").First(&product, updateDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Product not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to find product: "+err.Error())
	}

	// Start transaction
	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	product.Name = updateDto.Name
	product.Price = updateDto.Price
	product.Stock = updateDto.Stock
	product.Status = updateDto.Status
	product.Description = updateDto.Description

	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return ctx.String(http.StatusInternalServerError, "Failed to update product: "+err.Error())
	}

	// Update categories association
	if err := tx.Model(&product).Association("Categories").Clear(); err != nil {
		tx.Rollback()
		return ctx.String(http.StatusInternalServerError, "Failed to clear categories: "+err.Error())
	}

	if updateDto.CategoryIDs != "" {
		categoryIDs := strings.Split(updateDto.CategoryIDs, ",")
		var categories []models.Category
		if err := tx.Find(&categories, categoryIDs).Error; err != nil {
			tx.Rollback()
			return ctx.String(http.StatusInternalServerError, "Failed to find categories: "+err.Error())
		}

		// Validate that all categories are level 4 (parts level)
		for _, category := range categories {
			if category.Level != 4 {
				tx.Rollback()
				return ctx.String(http.StatusBadRequest, fmt.Sprintf("Products can only be associated with level 4 categories (parts). Category '%s' is level %d", category.Name, category.Level))
			}
		}

		if err := tx.Model(&product).Association("Categories").Append(categories); err != nil {
			tx.Rollback()
			return ctx.String(http.StatusInternalServerError, "Failed to associate categories: "+err.Error())
		}
	}

	tx.Commit()

	var products []models.Product
	storage.DB.Preload("Categories").Order("id DESC").Find(&products)
	return ctx.Html(components.ProductsTable(products))
}

func deleteProduct(ctx *controller.Context, deleteDto *dto.ByID) error {
	var product models.Product

	if err := storage.DB.First(&product, deleteDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Product not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to find product: "+err.Error())
	}

	// Start transaction
	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Clear categories association
	if err := tx.Model(&product).Association("Categories").Clear(); err != nil {
		tx.Rollback()
		return ctx.String(http.StatusInternalServerError, "Failed to clear categories: "+err.Error())
	}

	if err := tx.Delete(&product).Error; err != nil {
		tx.Rollback()
		return ctx.String(http.StatusInternalServerError, "Failed to delete product: "+err.Error())
	}

	tx.Commit()

	var products []models.Product
	storage.DB.Preload("Categories").Order("id DESC").Find(&products)
	return ctx.Html(components.ProductsTable(products))
}
