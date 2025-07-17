package orders

import (
	"main/internal/models"
	"main/internal/storage"
	"time"
)

var Seed = []models.Order{
	{
		OrderNumber:   "ORD-2024-001",
		CustomerName:  "John Smith",
		CustomerEmail: "john.smith@email.com",
		Total:         124.98,
		Status:        "completed",
		Date:          time.Now().AddDate(0, 0, -15),
		Items:         "Premium Engine Air Filter x2, Cabin Air Filter x1",
	},
	{
		OrderNumber:   "ORD-2024-002",
		CustomerName:  "Maria Rodriguez",
		CustomerEmail: "maria.r@email.com",
		Total:         89.99,
		Status:        "processing",
		Date:          time.Now().AddDate(0, 0, -10),
		Items:         "Brake Pads Set - Front x1",
	},
	{
		OrderNumber:   "ORD-2024-003",
		CustomerName:  "David Johnson",
		CustomerEmail: "d.johnson@email.com",
		Total:         259.97,
		Status:        "shipped",
		Date:          time.Now().AddDate(0, 0, -8),
		Items:         "Brake Rotors - Front Pair x1, Brake Pads Set - Front x1, Brake Fluid DOT 4 x2",
	},
	{
		OrderNumber:   "ORD-2024-004",
		CustomerName:  "Sarah Wilson",
		CustomerEmail: "sarah.wilson@email.com",
		Total:         45.99,
		Status:        "pending",
		Date:          time.Now().AddDate(0, 0, -25),
		Items:         "Standard Oil Filter x2, Synthetic Oil Filter x1",
	},
	{
		OrderNumber:   "ORD-2024-005",
		CustomerName:  "Michael Brown",
		CustomerEmail: "m.brown@email.com",
		Total:         179.98,
		Status:        "processing",
		Date:          time.Now().AddDate(0, 0, -3),
		Items:         "Timing Belt Kit x1, Engine Coolant x3, Spark Plugs Set (4pc) x2",
	},
	{
		OrderNumber:   "ORD-2024-006",
		CustomerName:  "Lisa Davis",
		CustomerEmail: "lisa.davis@email.com",
		Total:         67.99,
		Status:        "completed",
		Date:          time.Now().AddDate(0, 0, -12),
		Items:         "High-Flow Oil Filter x1, Inline Fuel Filter x2, Engine Oil 5W-30 Synthetic x1",
	},
	{
		OrderNumber:   "ORD-2024-007",
		CustomerName:  "Robert Taylor",
		CustomerEmail: "robert.t@email.com",
		Total:         324.96,
		Status:        "shipped",
		Date:          time.Now(),
		Items:         "Brake Caliper Kit x2, Premium Engine Air Filter x1",
	},
	{
		OrderNumber:   "ORD-2024-008",
		CustomerName:  "Jennifer White",
		CustomerEmail: "j.white@email.com",
		Total:         142.97,
		Status:        "cancelled",
		Date:          time.Now().AddDate(0, 0, -20),
		Items:         "Performance Air Filter x1, High-Pressure Fuel Filter x3, Diesel Fuel Filter x1",
	},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Create(&row)
	}
}
