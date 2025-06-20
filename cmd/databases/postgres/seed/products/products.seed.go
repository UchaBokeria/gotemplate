package products

import (
	"main/internal/models"
	"main/internal/storage"
)

var Seed = []models.Products{
	{
		Name:            "Product 1",
		Slug:            "product-1",
		Description:     "Product 1 description",
		DescriptionHtml: `<p>Description</p>`,
		Public:          true,
		Price:           100,
		Pics: []models.Files{
			{
				Name:       `686-e90-384-thumb__600_0_0_0_crop.jpg`,
				Original:   `1 L | 1118114-001`,
				Location:   "local",
				Path:       "/uploads/products/images/686-e90-384-thumb__600_0_0_0_crop.jpg",
				Size:       664,
				Base64:     "",
				Compressed: false,
				// TypeID:     pipes.GetDbTypeIdByExtension(`jpg`),
			},
			{
				Name:       `012-88a-943-thumb__600_0_0_0_crop.jpg`,
				Original:   `5 L | 1118114-005`,
				Location:   "local",
				Path:       "/uploads/products/images/012-88a-943-thumb__600_0_0_0_crop.jpg",
				Size:       664,
				Base64:     "",
				Compressed: false,
				// TypeID:     pipes.GetDbTypeIdByExtension(`jpg`),
			},
		},
	},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Create(&row)

		storage.DB.
			Model(&row).
			Association("Category").
			Append(
			// repository.FindByID[models.Categories](1),
			// repository.FindByID[models.Categories](2),
			)

	}
}
