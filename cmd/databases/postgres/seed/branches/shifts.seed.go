package branches

import (
	"main/internal/models"
	"main/internal/storage"
	"time"
)

var Shift = []models.Branch_shifts{
	{
		BranchesID: 1,
		Name:       "ორშ-შაბ",
		Slug:       "sunday",
		OpensAt:    time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:   time.Date(0, 1, 1, 16, 00, 0, 0, time.UTC),
	},
	{
		BranchesID: 1,
		Name:       "კვირა",
		Slug:       "sunday",
		OpensAt:    time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:   time.Date(0, 1, 1, 16, 00, 0, 0, time.UTC),
	},
	{
		BranchesID: 2,
		Name:       "ყოველდღე",
		Slug:       "sunday",
		OpensAt:    time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:   time.Date(0, 1, 1, 16, 00, 0, 0, time.UTC),
	},
	{
		BranchesID: 3,
		Name:       "ორშ-შაბ",
		Slug:       "sunday",
		OpensAt:    time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:   time.Date(0, 1, 1, 16, 00, 0, 0, time.UTC),
	},
	{
		BranchesID: 3,
		Name:       "ორშ-პარ",
		Slug:       "sunday",
		OpensAt:    time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:   time.Date(0, 1, 1, 16, 00, 0, 0, time.UTC),
	},
	{
		BranchesID: 4,
		Name:       "კვირა",
		Slug:       "sunday",
		OpensAt:    time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:   time.Date(0, 1, 1, 16, 00, 0, 0, time.UTC),
	},
	{
		BranchesID: 4,
		Name:       "კვირა",
		Slug:       "sunday",
		OpensAt:    time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:   time.Date(0, 1, 1, 16, 00, 0, 0, time.UTC),
	},
	{
		BranchesID: 5,
		Name:       "ყოველდღე",
		Slug:       "sunday",
		OpensAt:    time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:   time.Date(0, 1, 1, 16, 00, 0, 0, time.UTC),
	},
	{
		BranchesID: 6,
		Name:       "სამშაბ",
		Slug:       "sunday",
		OpensAt:    time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:   time.Date(0, 1, 1, 16, 00, 0, 0, time.UTC),
	},
}

func Shifts() {
	for _, row := range Shift {
		storage.DB.Create(&row)
	}
}
