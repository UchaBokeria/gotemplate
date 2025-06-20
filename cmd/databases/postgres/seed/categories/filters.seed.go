package categories

import (
	"main/internal/enums"
	"main/internal/models"
)

var CategoryFilters = []models.Category_filters{
	{
		Name: "Filter 1",
		Slug: "filter_1",
		Type: enums.Text,
		Options: []models.Category_filters_option{
			{
				Name:  "Option 1",
				Key:   "option_1",
				Value: "value_1",
			},
			{
				Name:  "Option 2",
				Key:   "option_2",
				Value: "value_2",
			},
			{
				Name:  "Option 3",
				Key:   "option_3",
				Value: "value_3",
			},
		},
		Default_value: "",
	},
	{
		Name: "Filter 2",
		Slug: "filter_2",
		Type: enums.Select,
		Options: []models.Category_filters_option{
			{
				Name:  "Option 4",
				Key:   "option_4",
				Value: "value_4",
			},
			{
				Name:  "Option 5",
				Key:   "option_5",
				Value: "value_5",
			},
			{
				Name:  "Option 6",
				Key:   "option_6",
				Value: "value_6",
			},
		},
		Default_value: "",
	},
	{
		Name:          "Filter 3",
		Slug:          "filter_3",
		Type:          enums.Text,
		Options:       nil,
		Default_value: "",
	},
}
