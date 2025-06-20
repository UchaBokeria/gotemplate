package main

import (
	"main/cmd/databases/postgres/seed/branches"
	"main/cmd/databases/postgres/seed/categories"
	"main/cmd/databases/postgres/seed/chat"
	"main/cmd/databases/postgres/seed/company"
	"main/cmd/databases/postgres/seed/faq"
	"main/cmd/databases/postgres/seed/files"
	"main/cmd/databases/postgres/seed/interfaces"
	"main/cmd/databases/postgres/seed/orders"
	"main/cmd/databases/postgres/seed/payments"
	"main/cmd/databases/postgres/seed/posts"
	"main/cmd/databases/postgres/seed/products"
	"main/cmd/databases/postgres/seed/users"
	"main/internal/config"
	"main/internal/storage"
)

func main() {
	config.SetupEnvironmentVariables()
	storage.Connect(storage.Default())

	company.Populate()
	branches.Cities()
	users.Populate()

	files.Types()
	files.Populate()

	posts.Types()
	posts.Populate()

	branches.Populate()
	branches.Shifts()

	faq.Populate()

	chat.Populate()

	categories.Populate()

	products.Populate()

	interfaces.Populate()
	interfaces.SocialMedia()
	interfaces.Slideshow()
	interfaces.Inovation()
	interfaces.Contact()
	interfaces.About()
	interfaces.Language()

	orders.Populate()
	payments.Populate()
}
