package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// LoadEnv loads environment variables from .env file
func LoadEnv() error {
	if isDevelopment() {
		log.Println("Development mode detected, loading .env file")
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using system environment variables")
			return fmt.Errorf("failed to load .env file: %w", err)
		}
	}

	return nil

}

// isDevelopment returns true if the environment is development
func isDevelopment() bool {
	return os.Getenv("ENV") == "" || os.Getenv("ENV") == "development"
}
