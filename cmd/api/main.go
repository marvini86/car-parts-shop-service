package main

import (
	"github.com/marvini86/car-parts-shop-service/internal/config"
	"github.com/marvini86/car-parts-shop-service/internal/db"
	"github.com/marvini86/car-parts-shop-service/internal/server"
)

// main function
func main() {

	// Load environment variables from .env file
	config.LoadEnv()

	// Open database connection
	db := db.OpenConnection()
	// Create server instance
	s := server.NewServerConfig(db)
	s.Init()
}
