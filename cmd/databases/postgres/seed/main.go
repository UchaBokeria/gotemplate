package main

import (
	"log"
	"main/cmd/databases/postgres/seed/categories"
	"main/cmd/databases/postgres/seed/general_settings"
	"main/cmd/databases/postgres/seed/invoices"
	"main/cmd/databases/postgres/seed/orders"
	"main/cmd/databases/postgres/seed/products"
	"main/cmd/databases/postgres/seed/translations"
	"main/cmd/databases/postgres/seed/users"
	"main/internal/config"
	"main/internal/storage"
)

func main() {
	config.SetupEnvironmentVariables()
	storage.Connect(storage.Default())

	// Seed core data
	users.Populate()
	general_settings.Populate()
	translations.Populate()
	
	if err := categories.SeedCategories(); err != nil {
		log.Fatal("Failed to seed categories:", err)
	}

	// Seed business data
	if err := products.Populate(); err != nil {
		log.Fatal("Failed to seed products:", err)
	}
	orders.Populate()
	invoices.Populate()
}
