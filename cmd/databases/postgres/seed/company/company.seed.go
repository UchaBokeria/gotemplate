package company

import (
	"main/internal/models"
	"main/internal/storage"
	"os"
)

func getFile(path string) string {
	dat, err := os.ReadFile(path)
	if err != nil {
		// fmt.Printf("files error: %T", err)
		panic(err)
	}
	return string(dat)
}

var Seed = []models.Company{
	{
		Path: "about",
		Name: "ჩვენს შესახებ",
		Body: getFile("cmd/databases/postgres/seed/interfaces/about.html"),
	},
	{
		Path: "terms",
		Name: "წესები და პირობები",
		Body: getFile("cmd/databases/postgres/seed/interfaces/terms.txt"),
	},
	{
		Path: "terms",
		Name: "წესები და პირობები",
		Body: getFile("cmd/databases/postgres/seed/interfaces/terms.txt"),
	},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Create(&row)
	}
}
