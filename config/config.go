package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DBConnectString is the statement to connect the application with the database
	DBConnectString = ""

	// DB_PORT represents the database port
	DB_PORT = 0

	// API_PORT
	API_PORT = 0
)

// LoadEnvs will initialize the environment variables for the application.
// Use "testPackage" or "testMain" for loading test environment for packages or main
// 	LoadEnvs("testPackage")
// Use "dev" for loading development environment
// 	LoadEnvs("dev")
// And "prod for loading prod ready environment
// 	LoadEnvs("prod")
func LoadEnvs(arg string) {

	var env string

	switch arg {
	case "testPackage":
		env = "../.env.test"
	case "testMain":
		env = ".env.test"
	case "dev":
		env = ".env.dev"
	case "prod":
		env = ".env"
	default:
		log.Fatal("unknown environment")
	}

	// Load .env file
	var err error
	if err = godotenv.Load(env); err != nil {
		log.Fatal(err)
	}

	// Get API_PORT and DB_PORT and check ports availability
	API_PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		API_PORT = 8000
	}

	DB_PORT, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		DB_PORT = 5432
	}

	// Get DB related variables
	DB_USER := os.Getenv("POSTGRES_USER")
	DB_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("POSTGRES_DB")

	DBConnectString = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

}
