package main

import (
	"main/internal/config"
	"main/internal/storage"

	"main/cmd/databases/postgres/schema"
)

func main() {
	config.SetupEnvironmentVariables()
	storage.Connect(storage.Default())
	storage.DB.Migrator().DropTable(schema.Models...)
}
