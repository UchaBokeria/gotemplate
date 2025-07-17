package invoices

import (
	"main/internal/models"
	"main/internal/storage"
	"time"
)

var Seed = []models.Invoice{
	{
		InvoiceNumber: "INV-2024-001",
		OrderID:       "ORD-2024-001",
		CustomerName:  "John Smith",
		CustomerEmail: "john.smith@email.com",
		Amount:        124.98,
		Status:        "paid",
		Date:          time.Now().AddDate(0, 0, -15),
		DueDate:       time.Now().AddDate(0, 0, -5),
	},
	{
		InvoiceNumber: "INV-2024-002",
		OrderID:       "ORD-2024-002",
		CustomerName:  "Maria Rodriguez",
		CustomerEmail: "maria.r@email.com",
		Amount:        89.99,
		Status:        "pending",
		Date:          time.Now().AddDate(0, 0, -10),
		DueDate:       time.Now().AddDate(0, 0, 5),
	},
	{
		InvoiceNumber: "INV-2024-003",
		OrderID:       "ORD-2024-003",
		CustomerName:  "David Johnson",
		CustomerEmail: "d.johnson@email.com",
		Amount:        259.97,
		Status:        "paid",
		Date:          time.Now().AddDate(0, 0, -8),
		DueDate:       time.Now().AddDate(0, 0, 2),
	},
	{
		InvoiceNumber: "INV-2024-004",
		OrderID:       "ORD-2024-004",
		CustomerName:  "Sarah Wilson",
		CustomerEmail: "sarah.wilson@email.com",
		Amount:        45.99,
		Status:        "overdue",
		Date:          time.Now().AddDate(0, 0, -25),
		DueDate:       time.Now().AddDate(0, 0, -15),
	},
	{
		InvoiceNumber: "INV-2024-005",
		OrderID:       "ORD-2024-005",
		CustomerName:  "Michael Brown",
		CustomerEmail: "m.brown@email.com",
		Amount:        179.98,
		Status:        "pending",
		Date:          time.Now().AddDate(0, 0, -3),
		DueDate:       time.Now().AddDate(0, 0, 12),
	},
	{
		InvoiceNumber: "INV-2024-006",
		OrderID:       "ORD-2024-006",
		CustomerName:  "Lisa Davis",
		CustomerEmail: "lisa.davis@email.com",
		Amount:        67.99,
		Status:        "paid",
		Date:          time.Now().AddDate(0, 0, -12),
		DueDate:       time.Now().AddDate(0, 0, -2),
	},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Create(&row)
	}
}
