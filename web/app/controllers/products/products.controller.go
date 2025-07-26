package products

import (
	"main/internal/models"
	"main/internal/storage"
	"main/web/app/controllers/products/dtos"
	"main/web/app/types"
	"main/web/app/view/components"
	"strconv"

	"github.com/UchaBokeria/goyard/controller"
)

func ProductsList(ctx *controller.Context, filters *dtos.ProductFilterDto) error {
	// Parse pagination parameters with defaults
	page := 1
	if filters.Page != 0 {
		page = filters.Page
	}

	limit := 12
	if filters.Limit != 0 {
		limit = filters.Limit
	}

	// Build query with category relationships preloaded
	query := storage.DB.Model(&models.Product{}).Preload("Categories").Where("status = ?", "active")

	// Search filter
	if filters.Search != "" {
		searchTerm := "%" + filters.Search + "%"
		query = query.Where("name ILIKE ? OR description ILIKE ?", searchTerm, searchTerm)
	}

	// Category filter - now using category IDs with many-to-many relationship
	if len(filters.CategoryIDs) > 0 {
		query = query.Joins("JOIN product_categories ON products.id = product_categories.product_id").
			Where("product_categories.category_id IN ?", filters.CategoryIDs).
			Distinct()
	}

	// Price range filters
	if filters.MinPrice != 0 {
		query = query.Where("price >= ?", filters.MinPrice)
	}

	if filters.MaxPrice != 0 {
		query = query.Where("price <= ?", filters.MaxPrice)
	}

	// Sorting
	sortBy := "created_at"
	sortOrder := "DESC"

	if filters.SortBy != "" {
		switch filters.SortBy {
		case "name", "price", "created_at":
			sortBy = filters.SortBy
		case "price-low":
			sortBy = "price"
			sortOrder = "ASC"
		case "price-high":
			sortBy = "price"
			sortOrder = "DESC"
		case "newest":
			sortBy = "created_at"
			sortOrder = "DESC"
		}
	}

	if filters.SortOrder != "" && (filters.SortOrder == "ASC" || filters.SortOrder == "DESC") {
		sortOrder = filters.SortOrder
	}

	// Get total count
	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
	}

	// Get products with pagination
	var products []models.Product
	offset := (page - 1) * limit
	if err := query.Order(sortBy + " " + sortOrder).Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
	}

	// Return HTML component for HTMX requests
	return ctx.Html(components.ProductsGridData(products, page, limit, int(totalCount), ctx.Get("webconfig").(types.WebConfig)))
}

func ProductDetailApp(ctx *controller.Context) error {
	id := ctx.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
	}

	var product models.Product
	if err := storage.DB.Preload("Categories").First(&product, productID).Error; err != nil {
		return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
	}

	// Get related products (same categories, excluding current product)
	var relatedProducts []models.Product
	if len(product.Categories) > 0 {
		// Get category IDs from the current product
		categoryIDs := make([]uint, len(product.Categories))
		for i, cat := range product.Categories {
			categoryIDs[i] = cat.ID
		}
		
		// Find products with same categories
		storage.DB.Preload("Categories").
			Joins("JOIN product_categories ON products.id = product_categories.product_id").
			Where("product_categories.category_id IN ? AND products.id != ? AND products.status = ?", 
				categoryIDs, product.ID, "active").
			Distinct().Order("created_at DESC").Limit(4).Find(&relatedProducts)
	}

	return ctx.Html(components.ProductDetailData(product, relatedProducts, ctx.Get("webconfig").(types.WebConfig)))
}

func ProductsCategories(ctx *controller.Context) error {
	// Get active categories that have products
	var categories []models.Category
	if err := storage.DB.
		Joins("JOIN product_categories ON categories.id = product_categories.category_id").
		Joins("JOIN products ON product_categories.product_id = products.id").
		Where("products.status = ? AND categories.is_active = ?", "active", true).
		Distinct().
		Order("categories.level ASC, categories.sort_order ASC, categories.name ASC").
		Find(&categories).Error; err != nil {
		return ctx.Html(components.ProductNotFound(ctx.Get("webconfig").(types.WebConfig)))
	}

	return ctx.Html(components.CategoriesList(categories, ctx.Get("webconfig").(types.WebConfig)))
}
