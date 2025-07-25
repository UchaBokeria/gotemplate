package invoices

import (
	"fmt"
	"net/http"
	"strconv"

	"main/internal/models"
	"main/internal/storage"
	"main/internal/types/dto"
	"main/web/admin/controllers/invoices/dtos"
	"main/web/admin/view/components"

	"github.com/UchaBokeria/goyard/controller"
	"gorm.io/gorm"
)

func list(ctx *controller.Context, listDto *dtos.InvoiceListDto) error {
	var invoices []models.Invoice
	var total int64

	query := storage.DB.Model(&models.Invoice{})

	if listDto.Search != "" {
		query = query.Where("customer_name ILIKE ? OR customer_email ILIKE ? OR invoice_number ILIKE ?",
			"%"+listDto.Search+"%", "%"+listDto.Search+"%", "%"+listDto.Search+"%")
	}

	if listDto.OrderID != "" {
		query = query.Where("order_id ILIKE ?", "%"+listDto.OrderID+"%")
	}

	if len(listDto.StatusMulti) > 0 {
		query = query.Where("status IN ?", listDto.StatusMulti)
	}

	query.Count(&total)

	sortBy := "created_at"
	sortOrder := "desc"

	if listDto.SortBy != "" {
		sortBy = listDto.SortBy
	}

	if listDto.SortOrder != "" {
		sortOrder = listDto.SortOrder
	}

	query = query.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))

	page := 1
	pageSize := 20

	if listDto.Page != "" {
		if p, err := strconv.Atoi(listDto.Page); err == nil && p > 0 {
			page = p
		}
	}

	if listDto.PageSize != "" {
		if ps, err := strconv.Atoi(listDto.PageSize); err == nil && ps > 0 && ps <= 100 {
			pageSize = ps
		}
	}

	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	if err := query.Find(&invoices).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to fetch invoices: "+err.Error())
	}

	return ctx.Html(components.InvoicesTable(invoices))
}

func show(ctx *controller.Context, showDto *dto.ByID) error {
	var invoice models.Invoice

	if err := storage.DB.First(&invoice, showDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Invoice not found")
		}
		return ctx.String(http.StatusInternalServerError, "Error fetching invoice: "+err.Error())
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Invoice #%s - Order: %s, Customer: %s, Amount: %.2f, Status: %s",
		invoice.InvoiceNumber, invoice.OrderID, invoice.CustomerName, invoice.Amount, invoice.Status))
}

func deleteInvoice(ctx *controller.Context, deleteDto *dto.ByID) error {
	var invoice models.Invoice

	if err := storage.DB.First(&invoice, deleteDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Invoice not found")
		}
		return ctx.String(http.StatusInternalServerError, "Error fetching invoice for deletion: "+err.Error())
	}

	if err := storage.DB.Delete(&invoice).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to delete invoice: "+err.Error())
	}

	var invoices []models.Invoice
	storage.DB.Order("created_at DESC").Find(&invoices)
	return ctx.Html(components.InvoicesTable(invoices))
}
