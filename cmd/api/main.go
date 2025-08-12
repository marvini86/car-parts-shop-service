// @title       Car Parts Shop API
// @version     1.0
// @description Backend service for car parts store
// @termsOfService http://swagger.io/terms/

// @contact.name  Marvin
// @contact.email marvin@example.com

// @host      localhost:8080
// @BasePath  /api/v1
// @schemes   http

package main

import (
	"github.com/marvini86/car-parts-shop-service/internal/config"
	"github.com/marvini86/car-parts-shop-service/internal/db"
	"github.com/marvini86/car-parts-shop-service/internal/server"
	"log"
)

// main function
func main() {

	// Load environment variables from .env file

	if err := config.LoadEnv(); err != nil {
		log.Fatal("Failed to load environment variables: %v", err)
	}

	// Open database connection
	dbConnection, err := db.OpenConnection()

	if err != nil {
		log.Fatal("Failed to open database connection: %v", err)
	}
	// Create server instance
	s := server.NewServerConfig(dbConnection)
	if err = s.Init(); err != nil {
		log.Fatal("Failed to initialize server: %v", err)
	}
}
