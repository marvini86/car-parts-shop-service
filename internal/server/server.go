package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/marvini86/car-parts-shop-service/docs"
	"github.com/marvini86/car-parts-shop-service/internal/handler"
	"github.com/marvini86/car-parts-shop-service/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"log"
	"os"
)

// ServerConfig represents the configuration for the server
type ServerConfig struct {
	db      *gorm.DB
	gin     *gin.Engine
	appPort string
}

// NewServerConfig creates a new ServerConfig instance
func NewServerConfig(db *gorm.DB) *ServerConfig {
	var apiPort string
	if port, ok := os.LookupEnv("API_PORT"); ok {
		log.Printf("Found APP_PORT environment variable, using value %s", port)
		apiPort = port
	} else {
		log.Printf("API port not set, using default value %s", "8080")
		apiPort = "8080"
	}

	return &ServerConfig{
		db:      db,
		appPort: apiPort,
		gin:     gin.Default(),
	}
}

// Init initializes the server
func (s *ServerConfig) Init() error {
	s.initRoutes()

	log.Printf("Starting server on port %s", s.appPort)

	if err := s.gin.Run(fmt.Sprintf(":%s", s.appPort)); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

// initRoutes initializes the routes for the server
func (s *ServerConfig) initRoutes() {
	r := s.gin

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	rg := r.Group("/api/v1")

	handler.NewItemHandler(service.NewItemService(s.db)).InitRoutes(rg)
	handler.NewOrderHandler(service.NewOrderService(s.db)).InitRoutes(rg)
}
