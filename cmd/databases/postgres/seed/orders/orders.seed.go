package orders

import (
	"main/internal/models"
	"main/internal/storage"
	"time"
)

var Seed = []models.Orders{
	{
		UserID:   1,
		StatusID: 1,
		Total:    0,
		Deadline: time.Now(),
		Comment:  "",
	},
	{
		UserID:   2,
		StatusID: 2,
		Total:    0,
		Deadline: time.Now(),
		Comment:  "",
	},
}

var status = []models.Order_status{
	{
		Name: "new",
	},
	{
		Name: "pending",
	},
	{
		Name: "inprocess",
	},
	{
		Name: "issued",
	},
	{
		Name: "canceled",
	},
	{
		Name: "done",
	},
}

func Populate() {
	for _, row := range status {
		storage.DB.Create(&row)
	}

	for _, row := range Seed {
		storage.DB.Create(&row)

		storage.DB.
			Model(&row).
			Association("Products").
			Append(
			// repository.FindByID[models.Products](1),
			)
	}
}
