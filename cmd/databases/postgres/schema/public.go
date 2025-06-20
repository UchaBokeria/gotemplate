package schema

import (
	"main/internal/models"
	"main/internal/storage"
)

func FilterEnums() {
	storage.DB.Exec(`CREATE TYPE filter_types AS ENUM (
		'range',
		'checkbox',
		'select',
		'radio',
		'text',
		'date',
		'number',
		'color',
		'switch',
		'multiselect',
		'rating',
		'search'
	);`)
}

var Models = []interface{}{
	&models.Company{},
	&models.Users{},
	&models.Cities{},
	&models.Districts{},
	&models.Branches{},
	&models.Branch_shifts{},

	&models.Category_filters_option{},
	&models.Category_filters{},
	&models.Categories{},

	&models.Chat_status{},
	&models.Chat_type{},
	&models.Chat{},
	&models.Chat_letters{},

	&models.Faq{},

	&models.File_types{},
	&models.Files{},

	&models.Posts_types{},
	&models.Posts{},

	&models.Products{},

	&models.Interface{},
	&models.Interface_slideShow{},
	&models.Interface_inovations{},
	&models.Interface_contact{},
	&models.Interface_about{},
	&models.Social_media{},
	&models.Languages{},
	&models.LanguageDictionary{},

	&models.Orders{},
	&models.Order_status{},

	&models.Payments{},
	&models.PaymentStatus{},
	&models.PaymentBranches{},
}
