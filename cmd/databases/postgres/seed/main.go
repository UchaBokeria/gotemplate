package main

import (
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

	// Seed business data
	products.Populate()
	orders.Populate()
	invoices.Populate()
}
