package main

import (
	"main/internal/config"
	"main/web"
)

func main() {
	config.SetupEnvironmentVariables()
	// storage.Connect(storage.Default())
	web.New()
}
