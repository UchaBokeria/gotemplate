package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvVarsType struct {
	Port        string
	Host        string
	GOENV       string
	Uploads     string
	PageMaxSize int
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASS     string
	DB_NAME     string
	DB_SSLMODE  string

	SENDGRID_API_KEY string
}

var Env EnvVarsType

func SetupEnvironmentVariables() {
	err := godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		log.Fatal(err)
	}

	/* Conversions */
	PageMaxSize, _ := strconv.Atoi(os.Getenv("PageMaxSize"))

	Env = EnvVarsType{
		Port:             os.Getenv("Port"),
		Host:             os.Getenv("Host"),
		GOENV:            os.Getenv("GOENV"),
		Uploads:          os.Getenv("Uploads"),
		PageMaxSize:      PageMaxSize,
		DB_HOST:          os.Getenv("DB_HOST"),
		DB_PORT:          os.Getenv("DB_PORT"),
		DB_USER:          os.Getenv("DB_USER"),
		DB_PASS:          os.Getenv("DB_PASS"),
		DB_NAME:          os.Getenv("DB_NAME"),
		DB_SSLMODE:       os.Getenv("DB_SSLMODE"),
		SENDGRID_API_KEY: os.Getenv("SENDGRID_API_KEY"),
	}
}
