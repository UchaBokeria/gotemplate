package products

import (
	"fmt"
	"net/http"

	"main/internal/models"
	"main/internal/storage"
	"main/internal/types/dto"
	"main/web/admin/controllers/products/dtos"
	"main/web/admin/view/components"

	"github.com/UchaBokeria/goyard/controller"
	"gorm.io/gorm"
)

func list(ctx *controller.Context[any], listDto *dtos.ProductListDto) error {
	var products []models.Product
	var total int64

	query := storage.DB.Model(&models.Product{})

	if listDto.Search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?",
			"%"+listDto.Search+"%", "%"+listDto.Search+"%")
	}

	if len(listDto.CategoryMulti) > 0 {
		query = query.Where("category IN ?", listDto.CategoryMulti)
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

func show(ctx *controller.Context[any], idDto *dto.ByID) error {
	var product models.Product

	if err := storage.DB.First(&product, idDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Product not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to fetch product: "+err.Error())
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Product: %s (%s) - $%.2f - Stock: %d - Status: %s",
		product.Name, product.Category, product.Price, product.Stock, product.Status))
}

func create(ctx *controller.Context[any], createDto *dtos.CreateProductDto) error {
	product := models.Product{
		Name:        createDto.Name,
		Category:    createDto.Category,
		Price:       createDto.Price,
		Stock:       createDto.Stock,
		Status:      createDto.Status,
		Description: createDto.Description,
	}

	if err := storage.DB.Create(&product).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to create product: "+err.Error())
	}

	var products []models.Product
	storage.DB.Order("id DESC").Find(&products)
	return ctx.Html(components.ProductsTable(products))
}

func update(ctx *controller.Context[any], updateDto *dtos.UpdateProductDto) error {
	var product models.Product

	if err := storage.DB.First(&product, updateDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Product not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to find product: "+err.Error())
	}

	product.Name = updateDto.Name
	product.Category = updateDto.Category
	product.Price = updateDto.Price
	product.Stock = updateDto.Stock
	product.Status = updateDto.Status
	product.Description = updateDto.Description

	if err := storage.DB.Save(&product).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to update product: "+err.Error())
	}

	var products []models.Product
	storage.DB.Order("id DESC").Find(&products)
	return ctx.Html(components.ProductsTable(products))
}

func deleteProduct(ctx *controller.Context[any], deleteDto *dto.ByID) error {
	var product models.Product

	if err := storage.DB.First(&product, deleteDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Product not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to find product: "+err.Error())
	}

	if err := storage.DB.Delete(&product).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to delete product: "+err.Error())
	}

	var products []models.Product
	storage.DB.Order("id DESC").Find(&products)
	return ctx.Html(components.ProductsTable(products))
}
