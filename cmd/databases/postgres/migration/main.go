package main

import (
	"main/cmd/databases/postgres/schema"
	"main/internal/config"
	"main/internal/storage"
)

func main() {
	config.SetupEnvironmentVariables()
	storage.Connect(storage.Default())
	storage.DB.Migrator().AutoMigrate(schema.Models...)
}
