package categories

import (
	"main/internal/models"
	"main/internal/storage"

	"github.com/UchaBokeria/goyard/pipes"

	"gorm.io/gorm"
)

var Seed = []models.Categories{
	{
		ParentID: nil,
		Level:    1,
		Name:     "Root",
		Slug:     "root",
		Public:   true,
		IconID:   pipes.Pint(15),
		Filters:  CategoryFilters[0:1],
		Children: []*models.Categories{
			{
				Level:   2,
				Name:    "middle",
				Slug:    "middle",
				Public:  true,
				IconID:  pipes.Pint(15),
				Filters: CategoryFilters[0:1],
				Children: []*models.Categories{
					{
						Level:           3,
						Name:            "baby",
						Slug:            "baby",
						Description:     "Baby",
						DescriptionHTML: "<p>Baby</p>",
						Public:          true,
						IconID:          pipes.Pint(15),
						Filters:         nil,
					},
				},
			},
		},
	},
	{
		ParentID: nil,
		Level:    1,
		Name:     "Root 2",
		Slug:     "root_2",
		Public:   true,
		IconID:   pipes.Pint(15),
		Filters:  CategoryFilters[0:1],
	},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&row)
	}
}
