package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	if isDevelopment() {
		log.Println("Development mode detected, loading .env file")
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using system environment variables")
		}
	}

}

// isDevelopment returns true if the environment is development
func isDevelopment() bool {
	return os.Getenv("ENV") == "" || os.Getenv("ENV") == "development"
}
