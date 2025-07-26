package main

import (
	"main/cmd/databases/postgres/seed/categories"
)

func main() {
	categories.GenerateSeedFile()
}
