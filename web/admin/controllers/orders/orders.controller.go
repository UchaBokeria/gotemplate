package orders

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"main/internal/models"
	"main/internal/storage"
	"main/internal/types/dto"
	"main/web/admin/controllers/orders/dtos"
	"main/web/admin/view/components"

	"github.com/UchaBokeria/goyard/controller"
	"gorm.io/gorm"
)

func list(ctx *controller.Context[any], listDto *dtos.OrderListDto) error {
	var orders []models.Order
	var total int64

	query := storage.DB.Model(&models.Order{})

	if listDto.Search != "" {
		query = query.Where("customer_name ILIKE ? OR customer_email ILIKE ? OR order_number ILIKE ?",
			"%"+listDto.Search+"%", "%"+listDto.Search+"%", "%"+listDto.Search+"%")
	}

	if len(listDto.StatusMulti) > 0 {
		query = query.Where("status IN ?", listDto.StatusMulti)
	}

	if len(listDto.DateMulti) > 0 {
		today := time.Now()
		var dateConditions []string
		var dateArgs []interface{}

		for _, dateRange := range listDto.DateMulti {
			switch dateRange {
			case "today":
				startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
				endOfDay := startOfDay.Add(24 * time.Hour)
				dateConditions = append(dateConditions, "(date >= ? AND date < ?)")
				dateArgs = append(dateArgs, startOfDay, endOfDay)
			case "week":
				weekAgo := today.Add(-7 * 24 * time.Hour)
				dateConditions = append(dateConditions, "date >= ?")
				dateArgs = append(dateArgs, weekAgo)
			case "month":
				monthAgo := today.Add(-30 * 24 * time.Hour)
				dateConditions = append(dateConditions, "date >= ?")
				dateArgs = append(dateArgs, monthAgo)
			}
		}

		if len(dateConditions) > 0 {
			for i, condition := range dateConditions {
				if i == 0 {
					query = query.Where(condition, dateArgs[i*2:]...)
				} else {
					query = query.Or(condition, dateArgs[i*2:]...)
				}
			}
		}
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

	if err := query.Find(&orders).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to fetch orders: "+err.Error())
	}

	return ctx.Html(components.OrdersTable(orders))
}

func show(ctx *controller.Context[any], showDto *dto.ByID) error {
	var order models.Order

	if err := storage.DB.First(&order, showDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Order not found")
		}
		return ctx.String(http.StatusInternalServerError, "Error fetching order: "+err.Error())
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Order #%s - Customer: %s (%s), Total: %.2f, Status: %s, Items: %d, Date: %s",
		order.OrderNumber, order.CustomerName, order.CustomerEmail, order.Total, order.Status, order.Items, order.Date.Format("2006-01-02")))
}

func deleteOrder(ctx *controller.Context[any], deleteDto *dto.ByID) error {
	var order models.Order

	if err := storage.DB.First(&order, deleteDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Order not found")
		}
		return ctx.String(http.StatusInternalServerError, "Error fetching order for deletion: "+err.Error())
	}

	if err := storage.DB.Delete(&order).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to delete order: "+err.Error())
	}

	var orders []models.Order
	storage.DB.Order("created_at DESC").Find(&orders)
	return ctx.Html(components.OrdersTable(orders))
}
